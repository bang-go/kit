package bgin

import "github.com/gin-gonic/gin"

type IBRouterGroup interface {
	RouterHandle(httpMethod, relativePath string, handlers ...gin.HandlerFunc)
}

type BRouterGroup struct {
	Group *gin.RouterGroup
}

// RouterHandle Handle 路由
func (g *BRouterGroup) RouterHandle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	g.Group.Handle(httpMethod, relativePath, handlers...)
}
