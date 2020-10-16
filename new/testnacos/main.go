/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var ConfigNacos string

type URLX struct {
	URLA string `json:"url_a"`
	URLB string `json:"url_b"`
	URLC string `json:"url_c"`
}

func main() {

	BurningUserTMP, err := strconv.ParseInt("123123", 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(BurningUserTMP)
	listener, err := net.Listen("tcp", "0.0.0.0:18000")
	if err != nil {
		log.Fatal(err)
	}
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
		// LogDir:              "/tmp/nacos/log",
		// CacheDir:            "/tmp/nacos/cache",
		LogDir:     "D:/go-mod/testnacos/tmp/nacos/log",
		CacheDir:   "D:/go-mod/testnacos/tmp/nacos/cache",
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

	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}

	// for i := 0; i < 10000; i++ {
	// 	time.Sleep(1 * time.Second)

	// 	fmt.Println(ConfigNacos)
	// 	fmt.Println(i)
	// }
	// //publish config
	// //config key=dataId+group+namespaceId
	// _, err = client.PublishConfig(vo.ConfigParam{
	// 	DataId:  "test-data",
	// 	Group:   "test-group",
	// 	Content: "hello world!",
	// })
	// _, err = client.PublishConfig(vo.ConfigParam{
	// 	DataId:  "test-data-2",
	// 	Group:   "test-group",
	// 	Content: "hello world!",
	// })
	// if err != nil {
	// 	fmt.Printf("PublishConfig err:%+v \n", err)
	// }

	// //get config
	// content, err := client.GetConfig(vo.ConfigParam{
	// 	DataId: "test-data",
	// 	Group:  "test-group",
	// })
	// fmt.Println("GetConfig,config :" + content)

	// //Listen config change,key=dataId+group+namespaceId.
	// err = client.ListenConfig(vo.ConfigParam{
	// 	DataId: "test-data",
	// 	Group:  "test-group",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
	// 	},
	// })

	// err = client.ListenConfig(vo.ConfigParam{
	// 	DataId: "test-data-2",
	// 	Group:  "test-group",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
	// 	},
	// })

	// _, err = client.PublishConfig(vo.ConfigParam{
	// 	DataId:  "test-data",
	// 	Group:   "test-group",
	// 	Content: "test-listen",
	// })

	// time.Sleep(2 * time.Second)

	// _, err = client.PublishConfig(vo.ConfigParam{
	// 	DataId:  "test-data-2",
	// 	Group:   "test-group",
	// 	Content: "test-listen",
	// })

	// time.Sleep(2 * time.Second)

	// //cancel config change
	// err = client.CancelListenConfig(vo.ConfigParam{
	// 	DataId: "test-data",
	// 	Group:  "test-group",
	// })

	// time.Sleep(2 * time.Second)
	// _, err = client.DeleteConfig(vo.ConfigParam{
	// 	DataId: "test-data",
	// 	Group:  "test-group",
	// })
	// time.Sleep(5 * time.Second)

	// searchPage, _ := client.SearchConfig(vo.SearchConfigParm{
	// 	Search:   "blur",
	// 	DataId:   "",
	// 	Group:    "",
	// 	PageNo:   1,
	// 	PageSize: 10,
	// })
	// fmt.Printf("Search config:%+v \n", searchPage)
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, ConfigNacos)
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
