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

	common.GET("/profile", playerCtrl.MyProfile, playerVal.IDInToken)
	common.PUT("/players", playerCtrl.UpdatePlayer, playerVal.IDInToken, playerVal.Update)
	common.PATCH("/me/password", playerCtrl.UpdateMyPassword)

	common.POST("/games/:id", gameCtrl.PlayByBotID, val.ValidateObjectID, gameVal.GameValue)
	common.POST("/games", gameCtrl.PlayRandom, gameVal.GameValue)
	common.GET("/games", gameCtrl.RecentGame, playerVal.IDInToken)
}
