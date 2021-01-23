package service

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/skeyic/neuron/utils"
	"net/http"
)

var (
	TheBarkAlertMaster = NewBarkAlertMaster()
)

//// Errors
//var (
//	ErrSvcNotFound = errors.New("service not found")
//)

type BarkAlertMaster struct {
	//svcLock *sync.RWMutex
	//svcMap  map[string]*BarkAlertService
}

func NewBarkAlertMaster() *BarkAlertMaster {
	return &BarkAlertMaster{
		//		svcLock: &sync.RWMutex{},
		//		svcMap:  make(map[string]*BarkAlertService),
	}
}

//
//func (m *BarkAlertMaster) Register(s *BarkAlertService) {
//	m.svcLock.Lock()
//	m.svcMap[s.ID] = s
//	m.svcLock.Unlock()
//}
//
//func (m *BarkAlertMaster) Get(id string) *BarkAlertService {
//	m.svcLock.RLock()
//	defer m.svcLock.RUnlock()
//	return m.svcMap[id]
//}
//
//func (m *BarkAlertMaster) Send(id, title, content string) error {
//	svc := m.Get(id)
//	if svc == nil {
//		return ErrSvcNotFound
//	}
//	return svc.SendAlert(title, content)
//}

func (m *BarkAlertMaster) SendAlert(id, title, content string) error {
	rCode, rBody, rError := utils.SendRequest(http.MethodPost, fmt.Sprintf(barkURL, id, EscapeString(title), EscapeString(content)), nil)
	if rError != nil {
		glog.Errorf("Failed to send alert, rCode: %d, rBody: %s, rError: %v", rCode, rBody, rError)
		return rError
	}

	return nil
}
