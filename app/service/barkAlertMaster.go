package service

import "sync"

var (
	TheBarkAlertMaster = NewBarkAlertMaster()
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
