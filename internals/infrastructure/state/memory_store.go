package state

import (
	"sync"

	"auptex.com/botnova/internals/domain/models"
)

type MemoryStateStore struct {
	mu sync.RWMutex
	// In-memory storage for robot states, keyed by robot ID
	robotStates map[string]*models.RobotState
	groupStates map[string]*models.RobotState
}

func NewMemoryStateStore() *MemoryStateStore {
	return &MemoryStateStore{
		robotStates: make(map[string]*models.RobotState),
		groupStates: make(map[string]*models.RobotState),
	}
}

func (s *MemoryStateStore) GetRobotState(robotId string) (*models.RobotState, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	state, exists := s.robotStates[robotId]
	return state, exists
}

func (s *MemoryStateStore) SetRobotState(robotId string, state *models.RobotState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.robotStates[robotId] = state
}

func (s *MemoryStateStore) GetGroupState(groupId string) (*models.RobotState, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	state, exists := s.groupStates[groupId]
	return state, exists
}

func (s *MemoryStateStore) SetGroupState(groupId string, state *models.RobotState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.groupStates[groupId] = state
}
