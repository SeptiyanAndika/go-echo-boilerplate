package utils

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
)

var Config config

var onceConfig sync.Once

func init() {
	onceConfig.Do(func() {
		if _, err := toml.DecodeFile("./config.toml", &Config); err != nil {
			fmt.Println(err)
		}
	})
}

type config struct {
	App      app
	Database database
}

type database struct {
	Server   string
	Port     string
	Database string
	User     string
	Password string
}

type app struct {
	Name string
	Port string
}
