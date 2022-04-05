package route

import (
	"card-game-golang/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// common ...
func common(e *echo.Echo) {
	auth := controller.Auth{}
	player := controller.Player{}
	game := controller.Game{}

	common := e.Group("/api")

	common.POST("/register", auth.PlayerRegister)
	common.POST("/login", auth.PlayerLogin)

	// middleware
	common.Use(middleware.JWT(envVars.Jwt.SecretKey))

	common.GET("/me", player.MyProfile)
	common.PUT("/me", player.UpdateMyProfile)
	common.PATCH("/me/password", player.UpdateMyPassword)

	common.POST("/api/games", game.Play)
	common.GET("/api/games", game.RecentGame)
}
