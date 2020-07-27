package share

import (
	"flag"
	"fmt"
	"sync/atomic"

	"sd"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"unsafe"
)

var globalVip *viper.Viper

var cfgFile = flag.String("conf", "./server.toml", "config file")
var appnameCfg = flag.String("appname", "", "app name")
var appName = ""

// Init
func Init(defaultAppName string) error {
	flag.Parse()
	if err := initRoute(); err != nil {
		return err
	}

	appName = defaultAppName
	if *appnameCfg != "" {
		appName = *appnameCfg
	}

	vip, err := loadServerConfig(*cfgFile)
	if err != nil {
		panic(err)
		return err
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&globalVip)), unsafe.Pointer(vip))
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("file changed and reload...", in.Name, in.Op)
		vipNew, errNew := loadServerConfig(*cfgFile)
		if errNew == nil {

			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&globalVip)), unsafe.Pointer(vipNew))
		}
	})

	//for ; ;  {
	//	fmt.Printf("%+v\n",GetConfig())
	//	time.Sleep(time.Second*3)
	//}

	werr := sd.RegisterSever(GetConfVip(), nil, appName)
	if werr != nil {
		return werr
	}
	return nil
}

func loadServerConfig(file string) (*viper.Viper, error) {
	// 获取配置信息
	viper.SetConfigType("toml")
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	appConfig := viper.GetViper()

	println("load server config ok!")
	return appConfig, nil
}

//GetConfVip 获取配置
func GetConfVip() *viper.Viper {
	conf := (*viper.Viper)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&globalVip))))
	return conf
}

//GetServerConf 当前服务配置
func GetServerConf(key string) interface{} {
	vip := GetConfVip()
	return vip.Get(appName + "." + key)
}

func GetServerConfByConf(conf *viper.Viper, key string) interface{} {
	return conf.Get(appName + "." + key)
}
