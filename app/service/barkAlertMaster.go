package service

import (
	"errors"
	"sync"
)

var (
	TheBarkAlertMaster = NewBarkAlertMaster()
)

// Errors
var (
	ErrSvcNotFound = errors.New("service not found")
)

type BarkAlertMaster struct {
	svcLock *sync.RWMutex
	svcMap  map[string]*BarkAlertService
}

func NewBarkAlertMaster() *BarkAlertMaster {
	return &BarkAlertMaster{
		svcLock: &sync.RWMutex{},
		svcMap:  make(map[string]*BarkAlertService),
	}
}

func (m *BarkAlertMaster) Register(s *BarkAlertService) {
	m.svcLock.Lock()
	m.svcMap[s.ID] = s
	m.svcLock.Unlock()
}

func (m *BarkAlertMaster) Get(id string) *BarkAlertService {
	m.svcLock.RLock()
	defer m.svcLock.RUnlock()
	return m.svcMap[id]
}

func (m *BarkAlertMaster) Send(id, title, content string) error {
	svc := m.Get(id)
	if svc == nil {
		return ErrSvcNotFound
	}
	return svc.SendAlert(title, content)
}
