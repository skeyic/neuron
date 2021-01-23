package model

import "github.com/skeyic/neuron/app/service"

type BarkAlert struct {
	ID   string
	Name string
}

func NewBarkAlert(id, name string) *BarkAlert {
	alert := &BarkAlert{ID: id, Name: name}
	return alert
}

func (b *BarkAlert) GetID() string {
	return b.ID
}

func (b *BarkAlert) Send(body *AlertBody) error {
	return service.TheBarkAlertMaster.SendAlert(b.ID, body.Title, body.Content)
}
