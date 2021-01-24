package model

import (
	"encoding/json"
	"errors"
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

var (
	ErrAlertAlreadyExist = errors.New("alert already exist")
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
		u, err := NewUserFromBytes(userByte.Content)
		if err != nil {
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

type User struct {
	ID string `json:"id"`
	UserInput
	lock       *sync.RWMutex
	BarkAlerts []*BarkAlert `json:"bark_alerts"`
	postMen    map[string]PostMan
}

func NewUserFromBytes(userByte []byte) (*User, error) {
	var (
		u = &User{}
	)
	err := json.Unmarshal(userByte, &u)
	if err != nil {
		glog.Errorf("failed to unmarshal user, byte: %s, err: %v", userByte, err)
		return nil, err
	}

	u.lock = &sync.RWMutex{}
	u.postMen = make(map[string]PostMan)
	for _, barkAlert := range u.BarkAlerts {
		u.postMen[barkAlert.ID] = barkAlert
	}

	return u, nil
}

type UserInput struct {
	Name string `json:"name"`
}

func (u *UserInput) Validate() error {
	return nil
}

func (u *UserInput) ToUser() *User {
	return &User{
		ID:        utils.GenerateUUID(),
		UserInput: *u,
		postMen:   make(map[string]PostMan),
		lock:      &sync.RWMutex{},
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

func (u *User) NewAlertService(barkAlert *BarkAlert) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	_, hit := u.postMen[barkAlert.GetID()]
	if !hit {
		u.postMen[barkAlert.GetID()] = barkAlert
		u.BarkAlerts = append(u.BarkAlerts, barkAlert)
	} else {
		return ErrAlertAlreadyExist
	}
	return u.Save()
}

func (u *User) Send(body *AlertBody) {
	u.lock.RLock()
	defer u.lock.RUnlock()

	for _, thePostMan := range u.postMen {
		go func(pm PostMan) {
			err := pm.Send(body)
			if err != nil {
				glog.Errorf("POSTMAN send failed, PostMan: %v, err: %v", pm, err)
			}
			glog.V(8).Infof("POSTMAN send successfully, PostMan: %v, title: %s", pm, body.Title)
		}(thePostMan)
	}
}

func (u *User) SendByID(id string, body *AlertBody) {
	u.lock.RLock()
	thePostMan := u.postMen[id]
	u.lock.RUnlock()

	if thePostMan == nil {
		glog.Errorf("POSTMAN not found, PostMan id: %s", id)
		return
	}

	go func(pm PostMan) {
		err := pm.Send(body)
		if err != nil {
			glog.Errorf("POSTMAN send failed, PostMan: %v, err: %v", pm, err)
		}
		glog.V(8).Infof("POSTMAN send successfully, PostMan: %v, title: %s", pm, body.Title)
	}(thePostMan)
}
