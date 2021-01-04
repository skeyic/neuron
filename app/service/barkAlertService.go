package service

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/skeyic/neuron/utils"
	"net/http"
	"strings"
)

const (
	barkURL = "https://api.day.app/%s/%s/%s"
)

var theEscapeMap = map[string]string{
	"%": "百分比",
}

func EscapeString(content string) string {
	var (
		result = content
	)
	for key, value := range theEscapeMap {
		result = strings.ReplaceAll(result, key, value)
	}

	return result
}

type BarkAlertService struct {
	ID string
}

func NewBarkAlertService(id string) *BarkAlertService {
	return &BarkAlertService{ID: id}
}

func (b *BarkAlertService) SendAlert(title, content string) error {
	rCode, rBody, rError := utils.SendRequest(http.MethodPost, fmt.Sprintf(barkURL, b.ID, EscapeString(title), EscapeString(content)), nil)
	if rError != nil {
		glog.Errorf("Failed to send alert, rCode: %d, rBody: %s, rError: %v", rCode, rBody, rError)
		return rError
	}

	return nil
}
