package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// common ...
func common(e *echo.Echo) {
	common := e.Group("/api")

	common.POST("/register", authCtrl.PlayerRegister, playerVal.Register)
	common.POST("/login", authCtrl.PlayerLogin, playerVal.Login)

	// middleware
	common.Use(middleware.JWT([]byte(envVars.Jwt.SecretKey)))

	common.GET("/me", playerCtrl.MyProfile, playerVal.MyProfile)
	common.PUT("/me", playerCtrl.UpdateMyProfile, playerVal.MyProfile, playerVal.Update)
	common.PATCH("/me/password", playerCtrl.UpdateMyPassword)

	common.POST("/games/:id", gameCtrl.PlayByBotID, val.ValidateObjectID, gameVal.GameValue)
	common.POST("/games", gameCtrl.PlayRandom)
	common.GET("/games", gameCtrl.RecentGame)
}
