package routers

import (
	"jaegerdemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/demo/", &controllers.TestDemoController{}, "get:Demo")
	beego.Router("/fileupload", &controllers.FileUploadController{}, "post:PPP")
	beego.Router("/down", &controllers.FileDownController{}, "post:Down")
}
