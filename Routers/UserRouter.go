package Routers

import (
	"Authentication/Configs"
	"Authentication/Configs/Middleware"
	"Authentication/Services"
	"github.com/labstack/echo/v4"
	"os"
)

type CustomEcho struct {
	*echo.Echo
}

func UserRouter() *echo.Echo {
	e := echo.New()
	Configs.InitEnv()

	//g := e.Group("/user", Middleware.jWTAuthen())
	g := e.Group("/user", Middleware.ValidateTokenMiddleware)
	g.GET(os.Getenv("ALL_API"), Services.UserAll)
	g.GET(os.Getenv("PROFILE_API"), Services.Profile)
	e.POST(os.Getenv("LOGIN_API"), Services.Login)
	e.POST(os.Getenv("REGISTER_API"), Services.Register)
	return e
}

func Execute(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8084"))
}
