package main

import (
	"fmt"
	"goft_v1/goft"
	"goft_v1/services"
)

func main() {
	g := goft.NewGoft()
	g.HandleService("GET", "./services/print", services.Print).Run()
	fmt.Println("hello world")
}
