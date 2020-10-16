package main

import (
	"fmt"
	_ "jaegerdemo/routers"

	"github.com/astaxie/beego"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var ConfigNacos string

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "10.88.10.85",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		//NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //namespace id
		NamespaceId:         "a6af0981-723d-4924-9211-e881b86e20d9",
		TimeoutMs:           5000,
		ListenInterval:      10000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		//LogDir:     "D:/go-mod/testnacos/tmp/nacos/log",
		//CacheDir:   "D:/go-mod/testnacos/tmp/nacos/cache",
		RotateTime: "1h",
		MaxAge:     3,
		LogLevel:   "debug",
		AccessKey:  "nacos",
		SecretKey:  "nacos",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		panic(err)
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "test1015",
		Group:  "test",
	})
	fmt.Println("GetConfig,config :" + content)
	ConfigNacos = content
	//Listen config change,key=dataId+group+namespaceId.
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test1015",
		Group:  "test",
		OnChange: func(namespace, group, dataId, data string) {
			//OnChange: func("a6af0981-723d-4924-9211-e881b86e20d9", "test", "test1015", content) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
			ConfigNacos = data
		},
	})
	beego.Run()
}
