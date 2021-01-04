package module

type BarkAlert struct {
	ID string
}

func NewBarkAlert(id string) *BarkAlert {
	return &BarkAlert{ID: id}
}

func (b *BarkAlert) Register() {

}
