package route

import (
	"card-game-golang/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// admin ...
func admin(e *echo.Echo) {
	auth := controller.Auth{}
	bot := controller.Bot{}

	admin := e.Group("/admin")

	admin.POST("/login", auth.AdminLogin)

	// middleware
	admin.Use(middleware.JWT(envVars.Jwt.SecretKey))

	admin.DELETE("/players/:id", nil)
	admin.GET("/players/:id", nil)
	admin.GET("/players", nil)

	admin.POST("/bots", bot.CreateBot)
	admin.GET("/bots/:id", bot.GetByID)
	admin.GET("/bots/:id", bot.GetList)
	admin.PUT("/bots/:id", bot.UpdateByID)
	admin.DELETE("/bots/:id", bot.DeleteByID)
	admin.DELETE("/bots", bot.DeleteAll)

}
