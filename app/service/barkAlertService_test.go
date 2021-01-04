package service

import (
	"fmt"
	"testing"
)

const (
	testBarkID = "kMHL4X8KSWDWzhZyZY3hgk"
)

func TestSendAlert(t *testing.T) {
	svc := NewBarkAlertService(testBarkID)
	fmt.Println(svc.SendAlert("TITLE", "It's 4:54 PM"))
}
