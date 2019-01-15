package routers

import (
	"github.com/astaxie/beego"
	"hotword/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tab", &controllers.MainController{}, "get:ShowTab")
	beego.Router("/find", &controllers.MainController{}, "get:FindWord")
	beego.Router("/findurl", &controllers.MainController{}, "get:FindUrl")
	beego.Router("/cloud", &controllers.MainController{}, "get:ShowCloud")
	beego.Router("/urls", &controllers.MainController{}, "get:ShowUrl")
	beego.Router("/findoneword", &controllers.MainController{}, "get,post:Search")
	beego.Router("/guanxi", &controllers.MainController{}, "get:ShowGuanxi")
	beego.Router("/add", &controllers.MainController{}, "get:ShowAdd;post:AddNewWord")
	beego.Router("/download", &controllers.MainController{}, "get:DownLoad")
}
