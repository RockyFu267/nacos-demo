package controllers

import (
	"encoding/json"
	"os"

	"github.com/astaxie/beego"
)

type FileDownController struct {
	beego.Controller
}
type returnFileDownJSON struct {
	Code  int64  `json:"code"`
	Data  string `json:"data"`
	Error string `json:"error"`
}
type PostPubData struct {
	Path string `json:"path"`
}

// 上传文件
func (this *FileDownController) Down() {
	var returnData returnFileDownJSON
	var PostData PostPubData
	defer this.Ctx.Request.Body.Close()
	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&PostData)
	if err != nil {
		beego.Error(" Error ", err, PostData)
		returnData.Code = 500
		returnData.Error = err.Error()
		this.Data["json"] = returnData
		this.Ctx.Output.JSON(returnData, true, true)
		return
	}
	_, err = os.Stat(PostData.Path)
	if err != nil {
		beego.Error(" Error ", err, PostData)
		returnData.Code = 500
		returnData.Error = err.Error()
		this.Data["json"] = returnData
		this.Ctx.Output.JSON(returnData, true, true)
		return
	}
	this.Ctx.Output.Download(PostData.Path)
	
	return
}
