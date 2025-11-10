package lib

import (
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler(logger *Logger, env *Env) *RequestHandler {
	// gin.SetMode(gin.DebugMode)
	// gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	return &RequestHandler{Gin: engine}
}
