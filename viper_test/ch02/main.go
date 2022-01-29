package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type User struct {
	UserName  string      `mapstructure:"name"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
}

func GetEnvInfo(env string) bool {

	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	v := viper.New()
	// 如何设置路径
	v.SetConfigFile("config-debug.yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v.Get("name"))

	user := User{}
	err := v.Unmarshal(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("user: ", user)

	fmt.Println("debug_env:", GetEnvInfo("GO_SHOP_DEBUG"))

	debug := GetEnvInfo("GO_SHOP_DEBUG")
	configFileName := "config-pro.yaml"
	if debug {
		configFileName = "config-debug.yaml"
	}

	fmt.Println(configFileName)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config changed: ", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&user)
		fmt.Println("changed: ", user)
	})

	time.Sleep(time.Second * 30)
}
