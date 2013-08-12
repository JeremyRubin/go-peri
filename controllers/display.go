package controllers

import (
    "github.com/astaxie/beego"
    //"fmt"
)
type DemandController struct {
    beego.Controller
}
var layout [4][4]map[string]string =  [4][4]map[string]string{
        {
        map[string]string{"name":"/dev/dp0/s05_lcd6/display", "type":"sliders"},
        map[string]string{"name":"", "type":""},
        map[string]string{"name":"", "type":""},
        map[string]string{"name":"/dev/dp0/s02_roten/led", "type":"toggles"},
        },
        {
            map[string]string{"name":"/dev/dp0/s04_slider4/updates", "type":"physical_sliders"},
            map[string]string{"name":"/dev/dp0/s02_roten/", "type":"physical_rotators"},
            map[string]string{"name":"", "type":""},
            map[string]string{"name":"", "type":""},
        },
        {
            map[string]string{"name":"", "type":""},
            map[string]string{"name":"", "type":""},
            map[string]string{"name":"", "type":""},
            map[string]string{"name":"", "type":""},
        },
        {
            map[string]string{"name":"", "type":""},
            map[string]string{"name":"", "type":""},
            map[string]string{"name":"", "type":""},
            map[string]string{"name":"", "type":""},
        },
    }
func (this *DemandController) Get() {
    this.Data["Layout"] = &layout
    this.TplNames = "panel.tpl"
}
