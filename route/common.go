package route

import (
	"card-game-golang/controller"
	"card-game-golang/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	auth      = controller.Auth{}
	player    = controller.Player{}
	game      = controller.Game{}
	playerVal = validations.Player{}
)

// common ...
func common(e *echo.Echo) {
	common := e.Group("/api")

	common.POST("/register", auth.PlayerRegister, playerVal.Register)
	common.POST("/login", auth.PlayerLogin, playerVal.Login)

	// middleware
	common.Use(middleware.JWT(envVars.Jwt.SecretKey))

	common.GET("/me", player.MyProfile)
	common.PUT("/me", player.UpdateMyProfile)
	common.PATCH("/me/password", player.UpdateMyPassword)

	common.POST("/api/games", game.Play)
	common.GET("/api/games", game.RecentGame)
}
