package bus

import (
	"context"
	"sync"
	"time"

	"auptex.com/botnova/internals/application/ports"
)

type WorkerPool struct {
	queue                chan ports.Event
	subscriptionRegistry *HandlerRegistry
	ctx                  context.Context
	cancel               context.CancelFunc
	wg                   sync.WaitGroup
	cfg                  PoolConfig
	logger               ports.Logger
}

func NewWorkerPool(logger ports.Logger, cfg PoolConfig, subscriptionRegistry *HandlerRegistry) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	wp := &WorkerPool{
		queue:                make(chan ports.Event, cfg.QueueSize),
		subscriptionRegistry: subscriptionRegistry,
		ctx:                  ctx,
		cancel:               cancel,
		cfg:                  cfg,
		logger:               logger,
	}

	wp.Start()
	return wp
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.cfg.MaxWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	// fmt.Printf("Worker %d started\n", id)
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			wp.logger.With(ports.Field{Key: "id", Value: id}).Debug("Worker shutting down")
			return
		case event := <-wp.queue:
			wp.logger.With(ports.Field{Key: "id", Value: id}).Debug("Worker processing event")
			wp.safeHandle(event)
		}
	}
}

func (wp *WorkerPool) safeHandle(event ports.Event) {
	defer func() {
		if r := recover(); r != nil {
			wp.logger.With(ports.Field{Key: "error", Value: r}).Error("Recovered from panic in worker")
		}
	}()
	wp.subscriptionRegistry.Dispatch(event, wp.logger)
}

func (wp *WorkerPool) Submit(event ports.Event) bool {
	// fmt.Printf("Submitting event to worker pool: %v", event)
	switch wp.cfg.BackPressureStrategy {
	case DropIfFull:
		select {
		case wp.queue <- event:
			return true
		default:
			wp.logger.Debug("Queue full, dropping event")
			return false
		}
	case BlockIfFull:
		wp.queue <- event
		return true
	case TimeoutIfFull:
		select {
		case wp.queue <- event:
			return true
		case <-time.After(wp.cfg.Timeout):
			wp.logger.Debug("Queue full, timeout reached, dropping event")
			return false
		}
	default:
		wp.logger.Debug("Unknown backpressure strategy, dropping event")
		return false
	}
}

func (wp *WorkerPool) Stop() {
	wp.cancel()
	wp.wg.Wait()
}
