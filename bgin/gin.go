package bgin

import (
	"github.com/bang-go/kit/util"
	"github.com/gin-gonic/gin"
)

type IGin interface {
	Use(middleware ...HandlerFunc)
	Run() error
	RouterGroup(relativePath string, handlers ...gin.HandlerFunc) IBRouterGroup
}

type Options struct {
	Mode        string //模式
	Addr        string //:8080 监听地址
	Middlewares []HandlerFunc
}

type HandlerFunc = gin.HandlerFunc
type Context = gin.Context
type Client struct {
	opt    *Options
	Engine *gin.Engine
}

func (c *Client) Use(middleware ...HandlerFunc) {
	c.Engine.Use(middleware...)
}

func (c *Client) Run() error {
	return c.Engine.Run()
}

func New(opt *Options) IGin {
	mode := util.If(opt.Mode != "", opt.Mode, ReleaseMode)
	gin.SetMode(mode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(opt.Middlewares...)
	return &Client{
		opt:    opt,
		Engine: engine,
	}
}

func (c *Client) RouterGroup(relativePath string, handlers ...gin.HandlerFunc) IBRouterGroup {
	rg := &BRouterGroup{}
	rg.Group = c.Engine.Group(relativePath, handlers...)
	return rg
}
