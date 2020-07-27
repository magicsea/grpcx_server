package share

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"unsafe"
)

var routConfig *RouteConfig

type RouteConfig struct {
	Project RouteProject
	//Servers []DServerNode
	Rules []RuleCfg
}

type RouteProject struct {
	Version string
}

type DServerNode struct {
	Name    string
	Type    string
	OutAddr string
}

type RuleCfg struct {
	Type   string
	Protos []string
}

func initRoute() error {

	conf, err := loadRouteConfig("./route.toml")
	if err != nil {
		panic(err)
		return err
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&routConfig)), unsafe.Pointer(conf))
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("file changed and reload...", in.Name, in.Op)
		confNew, errNew := loadRouteConfig(*cfgFile)
		if errNew == nil {
			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&routConfig)), unsafe.Pointer(confNew))
		}
	})
	return nil
}

func loadRouteConfig(file string) (*RouteConfig, error) {
	// 获取配置信息
	viper.SetConfigType("toml")
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	appConfig := viper.GetViper()
	var cfg RouteConfig
	err := appConfig.Unmarshal(&cfg)

	if err != nil {
		println("load route config error:", err)

		return nil, err
	}
	println("load route config ok!")
	return &cfg, err
}

func GetRouteConfig() *RouteConfig {
	conf := (*RouteConfig)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&routConfig))))
	return conf
}

//获取一类服务器
//func GetServersByType(t string) []DServerNode {
//	var l []DServerNode
//	for _, v := range GetRouteConfig().Servers {
//		if v.Type == t {
//			l = append(l, v)
//		}
//	}
//	return l
//}
//
////获取一个服务器
//func GetServerByName(name string) (*DServerNode, bool) {
//	for _, v := range GetRouteConfig().Servers {
//		if v.Name == name {
//			return &v, true
//		}
//	}
//	return nil, false
//}

//匹配服务
func GetRuleMatchProto(fullMethodName string) (*RuleCfg, bool) {
	for _, v := range GetRouteConfig().Rules {
		for _, p := range v.Protos {
			if strings.HasPrefix(fullMethodName, p) {
				return &v, true
			}
		}
	}
	return nil, false
}
