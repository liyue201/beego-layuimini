package main

import "github.com/astaxie/beego"

func main() {
	beego.SetStaticPath("/", "./layuimini")
	beego.Run()
}
