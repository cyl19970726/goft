package goft

import (
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	currentRouterGroup *gin.RouterGroup
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
func (this *Goft) HandleServices(httpMethod, relativePath string, handler ...gin.HandlerFunc) *Goft {
	if this.currentRouterGroup != nil {
		this.currentRouterGroup.Handle(httpMethod, relativePath, handler...)
	} else {
		this.Engine.Handle(httpMethod, relativePath, handler...)
	}

	return this
}

/*
添加路由组功能
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		engine:   group.engine,
	}
}
*/
func (this *Goft) GroupServices(relativePath string, handlers ...gin.HandlerFunc) *Goft {
	this.currentRouterGroup = this.Group(relativePath, handlers...)
	return this
}

func (this *Goft) Launch() {
	this.Run()
}
