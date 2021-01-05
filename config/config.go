package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

//import (
//	"github.com/jinzhu/configor"
//)
//
//// Config application all configs
//var Config = struct {
//	Port      uint `default:"8000" env:"PORT"`
//	DebugMode bool `default:"true" env:"DEBUG_MODE"`
//}{}
//
//func init() {
//	if err := configor.Load(&Config); err != nil {
//		panic(err)
//	}
//}

var Config = struct {
	Port uint `default:"8000"`
}{}

func init() {
	if err := configor.Load(&Config); err != nil {
		panic(err)
	}
	fmt.Printf("config: %#v", Config)
}
