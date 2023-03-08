package bgin

import (
	"github.com/bang-go/kit/http"
	"github.com/gin-gonic/gin"
)

var (
	DebugMode   = gin.DebugMode
	ReleaseMode = gin.ReleaseMode
	TestMode    = gin.TestMode

	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)
