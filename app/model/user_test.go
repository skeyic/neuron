package model

import (
	"fmt"
	"github.com/skeyic/neuron/config"
	"github.com/skeyic/neuron/utils"
	"testing"
)

func TestLoadUsers(t *testing.T) {
	tt := utils.NewMultiFileStoreSvc(config.Config.DataFolder, "Users")
	_, err := tt.ReadAll()
	fmt.Println(err)
}
