package bgin

import (
	"github.com/bang-go/kit/butil"
	"github.com/gin-gonic/gin"
)

var DefaultAddr = ":8080"

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
	Opt    *Options
	Engine *gin.Engine
}

func (c *Client) Use(middleware ...HandlerFunc) {
	c.Engine.Use(middleware...)
}

func (c *Client) Run() error {
	return c.Engine.Run(c.Opt.Addr)
}

func New(opt *Options) IGin {
	mode := butil.If(opt.Mode != "", opt.Mode, ReleaseMode)
	if opt.Addr == "" {
		opt.Addr = DefaultAddr
	}
	gin.SetMode(mode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(opt.Middlewares...)
	return &Client{
		Opt:    opt,
		Engine: engine,
	}
}

func (c *Client) RouterGroup(relativePath string, handlers ...gin.HandlerFunc) IBRouterGroup {
	rg := &BRouterGroup{}
	rg.Group = c.Engine.Group(relativePath, handlers...)
	return rg
}
