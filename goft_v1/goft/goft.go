package goft

import (
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

func NewGoft() *Goft {
	g := Goft{Engine: gin.New()}

	/*
		func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {  使用中间件的函数
	*/
	g.Engine.Use()
	return &g
}

func (this *Goft) Use(middleware ...gin.HandlerFunc) *Goft {
	this.Engine.Use(middleware...)

	return this
}

/*
 *  type HandlerFunc func(*Context)
 */
func (this *Goft) HandleService(httpMethod, relativePath string, handler ...gin.HandlerFunc) *Goft {
	this.Engine.Handle(httpMethod, relativePath, handler...)
	return this
}

func (this *Goft) Launch() {
	this.Run()
}
