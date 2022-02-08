package main

import (
	"fmt"
	"goft_v1/goft"
	"goft_v1/services"
)

func main() {
	service_v1()
	//service_v2()
}
func service_v1() {
	g := goft.NewGoft()
	g.HandleServices("GET", "./services/print", services.Print).Run()
	fmt.Println("hello world")
}

// 添加路由组功能
func service_v2() {
	g := goft.NewGoft()
	g.
		GroupServices("v1").HandleServices("GET", "./services/print", services.Print).
		GroupServices("v2").HandleServices("GET", "./services/login", services.Login).
		Run()
	fmt.Println("hello world")
}
