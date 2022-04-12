package route

import (
	"card-game-golang/controller"
	"card-game-golang/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	bot      = controller.Bot{}
	adminVal = validations.Admin{}
)

// admin ...
func admin(e *echo.Echo) {
	admin := e.Group("/admin")

	admin.POST("/login", auth.AdminLogin, adminVal.Login)

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
