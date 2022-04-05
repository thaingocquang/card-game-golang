package route

import (
	"card-game-golang/config"
	"github.com/labstack/echo/v4"
)

var envVars = config.GetEnv()

// Route ...
func Route(e *echo.Echo) {
	common(e)
	admin(e)
}
