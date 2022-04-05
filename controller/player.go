package controller

import "github.com/labstack/echo/v4"

// Player ...
type Player struct{}

// Register ...
func (p Player) Register(c echo.Context) error {
	return nil
}

// Login ...
func (p Player) Login(c echo.Context) error {
	return nil
}

// MyProfile ...
func (p Player) MyProfile(c echo.Context) error {
	return nil
}

// UpdateMyProfile ...
func (p Player) UpdateMyProfile(c echo.Context) error {
	return nil
}

// UpdateMyPassword ...
func (p Player) UpdateMyPassword(c echo.Context) error {
	return nil
}
