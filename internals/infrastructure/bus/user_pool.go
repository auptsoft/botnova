package bus

import (
	"sync"
	"time"

	"auptex.com/botnova/internals/application/ports"
)

type pooledUser struct {
	pool     *WorkerPool
	lastUsed time.Time
}

type UserScopedPool struct {
	mu                   sync.RWMutex
	pools                map[string]*pooledUser
	subscriptionRegistry *HandlerRegistry
	cfg                  UserPoolConfig
	logger               ports.Logger
}

func NewUserScopedPool(logger ports.Logger, cfg UserPoolConfig, subscriptionRegistry *HandlerRegistry) *UserScopedPool {
	usp := &UserScopedPool{
		pools:                make(map[string]*pooledUser),
		subscriptionRegistry: subscriptionRegistry,
		cfg:                  cfg,
		logger:               logger,
	}

	go usp.cleanupLoop()
	return usp
}

func (usp *UserScopedPool) Subscribe(s *ports.Subscription) ports.SubscriptionID {
	return usp.subscriptionRegistry.Subscribe(s)
}

func (usp *UserScopedPool) Unsubscribe(id ports.SubscriptionID) {
	usp.subscriptionRegistry.Unsubscribe(id)
}

func (usp *UserScopedPool) getOrCreate(userID string) *WorkerPool {
	now := time.Now()

	usp.mu.RLock()
	pu, ok := usp.pools[userID]
	usp.mu.RUnlock()

	if ok {
		usp.mu.Lock()
		pu.lastUsed = now
		usp.mu.Unlock()
		return pu.pool
	}

	usp.mu.Lock()
	defer usp.mu.Unlock()

	// double-check
	if pu, ok := usp.pools[userID]; ok {
		pu.lastUsed = now
		return pu.pool
	}

	// enforce max users
	if usp.cfg.MaxUsers > 0 && len(usp.pools) >= usp.cfg.MaxUsers {
		usp.evictOne()
	}

	wp := NewWorkerPool(usp.logger, usp.cfg.PoolConfig, usp.subscriptionRegistry)

	usp.pools[userID] = &pooledUser{
		pool:     wp,
		lastUsed: now,
	}

	return wp
}

func (usp *UserScopedPool) Submit(e ports.Event) bool {
	pool := usp.getOrCreate(e.UserID)
	return pool.Submit(e)
}

func (usp *UserScopedPool) evictOne() {
	var oldestKey string
	var oldestTime time.Time

	for k, v := range usp.pools {
		if oldestKey == "" || v.lastUsed.Before(oldestTime) {
			oldestKey = k
			oldestTime = v.lastUsed
		}
	}

	if oldestKey != "" {
		usp.pools[oldestKey].pool.Stop()
		delete(usp.pools, oldestKey)
	}
}

func (usp *UserScopedPool) cleanupLoop() {
	ticker := time.NewTicker(usp.cfg.CleanupInterval)
	defer ticker.Stop()

	for range ticker.C {
		usp.cleanup()
	}
}

func (usp *UserScopedPool) cleanup() {
	usp.mu.Lock()
	defer usp.mu.Unlock()

	now := time.Now()

	for userID, pu := range usp.pools {
		if now.Sub(pu.lastUsed) > usp.cfg.IdleTTL {
			pu.pool.Stop()
			delete(usp.pools, userID)
		}
	}
}
