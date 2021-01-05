package module

import "github.com/skeyic/neuron/app/service"

type BarkAlert struct {
	ID string
}

func NewBarkAlert(id string) *BarkAlert {
	return &BarkAlert{ID: id}
}

func (b *BarkAlert) Register() {
	service.TheBarkAlertMaster.Register(service.NewBarkAlertService(b.ID))
}

func (b *BarkAlert) Send(title, content string) {
	service.TheBarkAlertMaster.Send(b.ID, title, content)
}
