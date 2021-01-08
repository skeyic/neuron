package model

import (
	"github.com/skeyic/neuron/utils"
	"sync"
)

type PostMan interface {
	Send(title, content string) error
}

type User struct {
	ID string
	UserInput
	lock       *sync.RWMutex
	BarkAlerts map[string]*BarkAlert `json:"alerts"`
}

type UserInput struct {
	Name string `json:"name"`
}

func (u *UserInput) Validate() error {
	return nil
}

func (u *UserInput) ToUser() *User {
	return &User{
		ID:         utils.GenerateUUID(),
		UserInput:  *u,
		BarkAlerts: nil,
		lock:       &sync.RWMutex{},
	}
}

func (u *User) Save() error {
	return nil
}

func (u *User) NewAlertService(alertID, alertName string) {
	u.lock.Lock()
	u.BarkAlerts[alertID] = NewBarkAlert(alertID, alertName)
	u.lock.RUnlock()
}

func (u *User) Send(title, content string) {
	u.lock.RLock()
	defer u.lock.RUnlock()
	for _, barkAlert := range u.BarkAlerts {
		go barkAlert.Send(title, content)
	}
}
