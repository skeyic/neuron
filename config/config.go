package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Config = struct {
	DebugMode bool `default:"true"`
	Port      uint `default:"8000"`

	DataFolder string `default:"C:/Users/15902/go/src/github.com/skeyic/neuron/data" env:"DATA_FOLDER"`
}{}

func init() {
	if err := configor.Load(&Config); err != nil {
		panic(err)
	}
	fmt.Printf("config: %#v", Config)
}
