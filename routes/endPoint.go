package routes

import (
	"be_fs/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func EndPoint() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{"GET", "POST", "PUT", "UPDATE", "DELETE", "PATCH"},
	}))

	//userEndPoint
	e.GET("user/:limit/:offset", service.ReadAllUser)
	e.GET("user/read", service.ReadAllUserNoLimit)
	e.GET("user/:userId", service.ReadIdUser)
	e.POST("user", service.CreateUser)
	e.PUT("user/:userId", service.UpdateUser)
	e.DELETE("user/:userId", service.DeleteUser)

	e.Logger.Fatal(e.Start(":2000"))
}
