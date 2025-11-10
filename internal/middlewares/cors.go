package middlewares

import (
	"kingcom_api/internal/lib"

	"github.com/gin-contrib/cors"
)

type CorsMiddleware struct {
	handler *lib.RequestHandler
	logger  *lib.Logger
	env     *lib.Env
}

func NewCorsMiddleware(
	handler *lib.RequestHandler,
	logger *lib.Logger,
	env *lib.Env,
) *CorsMiddleware {
	return &CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

func (m *CorsMiddleware) Setup() {
	m.handler.Gin.Use(cors.New(cors.Config{
		AllowOrigins:     m.env.Cors.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
}
