package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type User struct {
	UserName string `mapstructure:"name"`
}

func main() {
	v := viper.New()
	// 如何设置路径
	v.SetConfigFile("config.yaml")

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

	for {
		fmt.Println("user: ", user.UserName)
		time.Sleep(time.Second)
	}

}
