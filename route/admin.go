package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// admin ...
func admin(e *echo.Echo) {
	admin := e.Group("/admin")

	admin.POST("/login", authCtrl.AdminLogin, adminVal.Login)

	// middleware
	admin.Use(middleware.JWT([]byte(envVars.Jwt.SecretKey)))

	admin.GET("/players/:id", playerCtrl.GetByID, val.ValidateObjectID)
	admin.GET("/players", playerCtrl.GetList)
	admin.PUT("/players/:id", playerCtrl.UpdateProfileByID, val.ValidateObjectID, playerVal.Profile)
	admin.DELETE("/players/:id", playerCtrl.DeleteByID, val.ValidateObjectID)
	admin.DELETE("/players", playerCtrl.DeleteAll)

	admin.POST("/bots", botCtrl.Create, botVal.Create)
	admin.GET("/bots/:id", botCtrl.GetByID, val.ValidateObjectID)
	admin.GET("/bots", botCtrl.GetList)
	admin.PUT("/bots/:id", botCtrl.UpdateByID, val.ValidateObjectID, botVal.Update)
	admin.DELETE("/bots/:id", botCtrl.DeleteByID, val.ValidateObjectID)
	admin.DELETE("/bots", botCtrl.DeleteAll)

}
