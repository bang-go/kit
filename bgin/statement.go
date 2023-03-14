package bgin

import (
	"github.com/gin-gonic/gin"
)

var (
	DebugMode   = gin.DebugMode
	ReleaseMode = gin.ReleaseMode
	TestMode    = gin.TestMode

	MethodGet     = bhttp.MethodGet
	MethodHead    = bhttp.MethodHead
	MethodPost    = bhttp.MethodPost
	MethodPut     = bhttp.MethodPut
	MethodPatch   = bhttp.MethodPatch
	MethodDelete  = bhttp.MethodDelete
	MethodConnect = bhttp.MethodConnect
	MethodOptions = bhttp.MethodOptions
	MethodTrace   = bhttp.MethodTrace
)
