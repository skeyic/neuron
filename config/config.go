package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Config = struct {
	DebugMode bool `default:"true"`
	Port      uint `default:"8000"`
}{}

func init() {
	if err := configor.Load(&Config); err != nil {
		panic(err)
	}
	fmt.Printf("config: %#v", Config)
}
