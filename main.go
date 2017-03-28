package main

import (
	"os"
	"strconv"

	_ "github.com/novelmanager/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		// beego.BConfig.WebConfig.DirectoryIndex = true
		// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetStaticPath("/favicon.ico", "static/favicon.ico")
	beego.SetStaticPath("/inline.bundle.js", "static/inline.bundle.js")
	beego.SetStaticPath("/inline.bundle.js.map", "static/inline.bundle.js.map")
	beego.SetStaticPath("/main.bundle.js", "static/main.bundle.js")
	beego.SetStaticPath("/main.bundle.js.map", "static/inline.bundle.js.map")
	beego.SetStaticPath("/polyfills.bundle.js", "static/polyfills.bundle.js")
	beego.SetStaticPath("/polyfills.bundle.js.map", "static/polyfills.bundle.js.map")
	beego.SetStaticPath("/styles.bundle.js", "static/styles.bundle.js")
	beego.SetStaticPath("/styles.bundle.js.map", "static/styles.bundle.js.map")
	beego.SetStaticPath("/vendor.bundle.js", "static/vendor.bundle.js")
	beego.SetStaticPath("/vendor.bundle.js.map", "static/vendor.bundle.js.map")

	beego.SetStaticPath("/", "static/index.html")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}

	beego.Run()
}
