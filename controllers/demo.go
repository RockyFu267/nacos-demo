package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	PUB "jaegerdemo/pubilc"
	"net"
	"net/http"

	"github.com/astaxie/beego"
)

//TestDemoController 获取taskrun对应的podname
type TestDemoController struct {
	beego.Controller
}

//API-POST的结构体 DataProject
//returnDeleteProjectJSON 返回的数据结构
type returnDeleteProjectJSON struct {
	Code  int64    `json:"code"`
	Data  []string `json:"data"`
	Error string   `json:"error"`
}

//Demo **
func (p *TestDemoController) Demo() {
	var returnData returnDeleteProjectJSON

	returnData.Code = 200

	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ip.String())
	var IPTMP string
	IPTMP = ip.String()
	var A string
	A = "I am AA"
	A = A + ":" + IPTMP
	returnData.Data = append(returnData.Data, A)

	//testURL := "aa.testxiao.svc.cluster.local"
	//testURL := PUB.ConfigNacos
	jsonData, err := PUB.ReJsonNacosCon(PUB.ConfigNacos)
	if err != nil {
		beego.Error("Get ERROR", " ")
		p.Data["json"] = returnData
		p.Ctx.Output.JSON(returnData, true, true)
		return
	}
	//debug
	fmt.Println(jsonData.URLB)
	//
	url := "http://" + jsonData.URLB + "/demo/"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Connection", "close")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		beego.Error("Get ERROR", " ")
		p.Data["json"] = returnData
		p.Ctx.Output.JSON(returnData, true, true)
		return
	}
	resp.Header.Add("Connection", "close")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("Ioutil ERROR", " ")
		p.Data["json"] = returnData
		p.Ctx.Output.JSON(returnData, true, true)
		return
	}
	defer resp.Body.Close()
	var dat returnDeleteProjectJSON
	if resjson := json.Unmarshal([]byte(data), &dat); resjson != nil {
		beego.Error(resjson)
		p.Data["json"] = returnData
		p.Ctx.Output.JSON(returnData, true, true)
		return
	}
	for k := range dat.Data {
		returnData.Data = append(returnData.Data, dat.Data[k])
	}

	p.Data["json"] = returnData
	p.Ctx.Output.JSON(returnData, true, true)
	return
}

func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
