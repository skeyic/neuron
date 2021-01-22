package model

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/skeyic/neuron/config"
	"github.com/skeyic/neuron/utils"
	"os"
	"sync"
)

var (
	theUsersFileStore = utils.NewMultiFileStoreSvc(config.Config.DataFolder+"/Users/", "")
	TheUsersMaster    = &UsersMaster{
		lock:    &sync.RWMutex{},
		dataMap: make(map[string]*User),
	}
)

type UsersMaster struct {
	lock    *sync.RWMutex
	dataMap map[string]*User
}

func (m *UsersMaster) Init() error {
	userBytes, err := theUsersFileStore.ReadAll()
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		glog.Errorf("failed to read all users, err: %v", err)
		return err
	}
	for _, userByte := range userBytes {
		var (
			u = &User{}
		)
		err := json.Unmarshal(userByte.Content, &u)
		if err != nil {
			glog.Errorf("failed to unmarshal user, byte: %s, err: %v", userByte.Content, err)
			continue
		}
		m.dataMap[u.ID] = u
		glog.V(4).Infof("Load user: %s", u.ID)
	}
	return nil
}

func (m *UsersMaster) AddUser(u *User) {
	m.lock.Lock()
	m.dataMap[u.ID] = u
	m.lock.Unlock()
}

func (m *UsersMaster) GetUsers() (users []*User) {
	m.lock.RLock()
	for _, u := range m.dataMap {
		users = append(users, u)
	}
	m.lock.RUnlock()
	return
}

func (m *UsersMaster) GetUser(ID string) (user *User) {
	m.lock.RLock()
	user = m.dataMap[ID]
	m.lock.RUnlock()
	return
}

type PostMan interface {
	Send(title, content string) error
}

type User struct {
	ID string `json:"id"`
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
	uByte, _ := json.Marshal(u)
	err := theUsersFileStore.Save(u.ID, uByte)
	if err != nil {
		return err
	}
	TheUsersMaster.AddUser(u)
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
