package server

import "github.com/labstack/echo"

func Login(c echo.Context) error {
	return c.String(200, "Login")
}
