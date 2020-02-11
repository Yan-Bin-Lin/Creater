package setting

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

var (
	Config *ConfigStruct
	Databases map[string]*DatabaseStruct
	Servers map[string]*ServerStruct
)


type ServerStruct struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	RunMode string `yaml:"RunMode,omitempty"`
	ReadTimeout time.Duration `yaml:"ReadTimeout,omitempty"`
	WriteTimeout time.Duration `yaml:"WriteTimeout,omitempty"`
}

type DatabaseStruct struct {
	Driver string `yaml:"driver"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
	Option map[string]int `yaml:"option,omitempty"`
}

type ConfigStruct struct {
	Servers map[string]*ServerStruct
	Databases map[string]*DatabaseStruct
}


func init() {
	// read file
	config, err := ioutil.ReadFile("../config/app/app.yaml")
	if err != nil {
		panic(err)
	}

	// binding
	var c ConfigStruct
	err = yaml.Unmarshal(config, &c)
	if err != nil {
		panic(err)
	}

	Config = &c
	Databases = Config.Databases
	Servers = Config.Servers
}
