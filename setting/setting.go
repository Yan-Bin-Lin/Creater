package setting

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var (
	Config   *ConfigStruct
	DBs      map[string]*DBStruct
	Servers  map[string]*ServerStruct
	ErrorMap ErrorMapStruct
)

type ServerStruct struct {
	Host         string            `yaml:"host"`
	Port         int               `yaml:"port"`
	RunMode      string            `yaml:"RunMode,omitempty"`
	ReadTimeout  time.Duration     `yaml:"ReadTimeout,omitempty"`
	WriteTimeout time.Duration     `yaml:"WriteTimeout,omitempty"`
	FilePath     string            `yaml:"FilePath,omitempty"`
	PostKey      map[string]string `yaml:"PostKey,omitempty"`
}

type DBStruct struct {
	Driver   string         `yaml:"driver"`
	User     string         `yaml:"user"`
	Password string         `yaml:"password"`
	Name     string         `yaml:"name"`
	Param    string         `yaml:"param,omitempty"`
	Option   map[string]int `yaml:"option,omitempty"`
}

type ConfigStruct struct {
	Servers   map[string]*ServerStruct
	Databases map[string]*DBStruct
}

type ErrorMapStruct map[int]string

var (
	WorkPath = "../config/app/"
)

func init() {
	if strings.HasSuffix(os.Args[0], ".test") {
		// change path if in test
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "..")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}
	}

	config, err := ioutil.ReadFile(WorkPath + "app.yaml")
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
	DBs = Config.Databases
	Servers = Config.Servers

	// read error config file
	config, err = ioutil.ReadFile(WorkPath + "error.yaml")
	if err != nil {
		panic(err)
	}

	// binding
	ErrorMap = make(map[int]string)
	err = yaml.Unmarshal(config, &ErrorMap)
	if err != nil {
		panic(err)
	}
}
