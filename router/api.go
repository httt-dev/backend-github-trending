package router

import (
	"backend-github-trending/handler"
	"github.com/labstack/echo"
	myMiddleware "backend-github-trending/middleware"
)

type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn, myMiddleware.ISAdmin())
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
}