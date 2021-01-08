package model

import "github.com/skeyic/neuron/app/service"

type BarkAlert struct {
	ID   string
	Name string
}

func NewBarkAlert(id, name string) *BarkAlert {
	alert := &BarkAlert{ID: id, Name: name}
	alert.Register()
	return alert
}

func (b *BarkAlert) Register() {
	service.TheBarkAlertMaster.Register(service.NewBarkAlertService(b.ID))
}

func (b *BarkAlert) Send(title, content string) error {
	return service.TheBarkAlertMaster.Send(b.ID, title, content)
}
