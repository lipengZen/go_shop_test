package main

import (
	"encoding/json"
	"fmt"
	"oldpackage_go_shop/nacos_test/config"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "ba63c7bc-827d-49fb-9d56-0cbf184c420c", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		fmt.Println("nacos err : ", err)
		return
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user_web.json",
		Group:  "dev"})

	fmt.Println(content, err)

	// 监听动态变化
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user_web.yaml",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	fmt.Println("动态监听 err: ", err)

	// time.Sleep(30 * time.Second)

	config := &config.ServerConfig{}
	json.Unmarshal([]byte(content), config)
	fmt.Println("转化成结构体: ", config)
}
