package middlewares

import "go.uber.org/fx"

type Middleware interface{ Setup() }

type Middlewares []Middleware

func NewMiddlewares(
	cors *CorsMiddleware,
	jwtAuth *JwtAuthMiddleware,
) *Middlewares {
	return &Middlewares{
		cors,
		jwtAuth,
	}
}

func (m Middlewares) Setup() {
	for _, md := range m {
		md.Setup()
	}
}

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewMiddlewares),
	fx.Provide(NewJwtAuthMiddleware),
)
