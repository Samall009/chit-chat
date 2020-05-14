package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// app类
type App struct {
	Address string
	Static  string
	Log     string
}

// db类
type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
	Charset  string
}

// 配置容器
type Configuration struct {
	App App
	Db  Database
}

// 配置容器
var config *Configuration

// 单例
var once sync.Once

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return config
}
