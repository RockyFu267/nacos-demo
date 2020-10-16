package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type FileUploadController struct {
	beego.Controller
}

// 上传文件的页面
func (c *FileUploadController) Get() {
	c.TplName = "fileupload.html"
}

// 上传文件
func (this *FileUploadController) PPP() {
	// uploadfilename，这是一个key值，对应的是html中input type-‘file’的name属性值
	f, h, err := this.GetFile("uploadfilename")
	if err != nil {
		log.Fatal("getfile err", err)
	}

	// 关闭上传的文件，不然的话会出现临时文件不能清除的情况
	defer f.Close()
	// 保存位置在 static/upload, 没有文件夹要先创建
	h.Filename = "fuck.html"
	this.SaveToFile("uploadfilename", "static/upload/"+h.Filename)
	// html页面
	this.TplName = "fileupload.html"
	return
}
