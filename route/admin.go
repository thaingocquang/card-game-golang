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

	admin.GET("/profiles/:id", playerCtrl.GetByID, val.ValidateObjectID)
	admin.GET("/profiles", playerCtrl.GetListProfile, val.Paging)
	admin.PUT("/profiles/:id", playerCtrl.UpdateProfileByID, val.ValidateObjectID, playerVal.Profile)
	admin.DELETE("/profiles/:id", playerCtrl.DeleteByID, val.ValidateObjectID)
	admin.DELETE("/profiles", playerCtrl.DeleteAll)

	admin.POST("/bots", botCtrl.Create, botVal.Create)
	admin.GET("/bots/:id", botCtrl.GetByID, val.ValidateObjectID)
	admin.GET("/bots", botCtrl.GetList, val.Paging)
	admin.PUT("/bots/:id", botCtrl.UpdateByID, val.ValidateObjectID, botVal.Update)
	admin.DELETE("/bots/:id", botCtrl.DeleteByID, val.ValidateObjectID)
	admin.DELETE("/bots", botCtrl.DeleteAll)

	admin.GET("/games", gameCtrl.GetList)
}
