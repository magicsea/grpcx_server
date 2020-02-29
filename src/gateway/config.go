package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync/atomic"

	"unsafe"
)

var globalConfig *Config
type Config struct {
	Project Project
	Server Server
	Backends []BackendConfig
	Servers []ServerNode
}

type Project struct {
	Name string
	Version string
	HeartbeatTime  int64
}

type Server struct {
	ListenClient string
	ListenServer string
}

type BackendConfig struct {
	Tag int
	Backend []string
	Filter string
	Module []string
}
type ServerNode struct {
	Name string
	Type string
	ListenAddr string
	OutAddr string
}
func init() {

	conf,err := loadRouterConfig()
	if err!=nil {
		panic(err)
	}
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&globalConfig)),unsafe.Pointer(conf))
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("file changed and reload...",in.Name,in.Op)
		confNew,errNew := loadRouterConfig()
		if errNew==nil {
			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&globalConfig)),unsafe.Pointer(confNew))
		}
	})

	//for ; ;  {
	//	fmt.Printf("%+v\n",GetConfig())
	//	time.Sleep(time.Second*3)
	//}
}

func loadRouterConfig() (*Config,error) {
	// 获取配置信息
	viper.SetConfigType("toml")
	viper.SetConfigFile("./config.toml")
	if err := viper.ReadInConfig(); err != nil {
		return nil,err
	}
	appConfig := viper.GetViper()
	var cfg Config
	err := appConfig.Unmarshal(&cfg)
	if err == nil {
		println("load config ok!")
		return &cfg,nil
	}
	println("load config error:",err)
	return nil,err
}

//GetConfig
func GetConfig() *Config {
	conf := (*Config)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&globalConfig))))
	return conf
}