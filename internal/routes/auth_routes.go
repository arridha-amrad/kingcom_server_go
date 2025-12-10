package routes

import (
	authcontroller "kingcom_api/internal/controllers/auth_controller"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/middlewares"
)

type AuthRoutes struct {
	handler *lib.RequestHandler
	ctrl    *authcontroller.AuthController
	jwtAuth *middlewares.JwtAuthMiddleware
}

func NewAuthRoutes(
	handler *lib.RequestHandler,
	ctrl *authcontroller.AuthController,
	jwtAuth *middlewares.JwtAuthMiddleware,
) *AuthRoutes {
	return &AuthRoutes{
		handler: handler,
		ctrl:    ctrl,
		jwtAuth: jwtAuth,
	}
}

func (r *AuthRoutes) Setup() {
	rtr := r.handler.Gin.Group("/api/auth")
	{
		rtr.POST("/forgot-password", r.ctrl.ForgotPassword)
		rtr.POST("/login", r.ctrl.Login)
		rtr.POST("/logout", r.jwtAuth.Handler, r.ctrl.Logout)
		rtr.GET("/me", r.jwtAuth.Handler, r.ctrl.Me)
		rtr.POST("/refresh-token", r.ctrl.RefreshToken)
		rtr.POST("/resend-verification", r.ctrl.ResendVerification)
		rtr.POST("/reset-password", r.ctrl.ResetPassword)
		rtr.POST("/signup", r.ctrl.SignUp)
		rtr.POST("/verify-account", r.ctrl.VerifyNewAccount)
	}
}
