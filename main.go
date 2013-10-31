package main

import (
	"demand/controllers"
	"github.com/astaxie/beego"
	"github.com/fzzy/sockjs-go/sockjs"
	"runtime"

	//"log"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf := sockjs.NewConfig()
	pool := sockjs.NewSessionPool()
	dev := controllers.NewDevice(pool, 10)
	dev.AddReader("/dev/dp0/s02_roten/count")
	dev.AddReader("/dev/dp0/s04_slider4/updates")
	sockjshandler := sockjs.NewHandler("/socket", dev.SocketHandler, conf)
    
    dc := controllers.DemandController{}


	beego.Router("/", &dc)
	beego.Router("/test", &controllers.ChatController{})
	beego.RouterHandler("/socket/:info(.*)", sockjshandler)
	beego.Run()
}
