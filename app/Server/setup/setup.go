package setup

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

//MiddleWare Attach middle ware
func (e *echo.Echo) MiddleWare() *echo.Echo {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	return e
}

func (e *echo.Echo) Port(post string) {
	re := echo.Router()
	re
}

//TLS configures TLS Settings
func (e *echo.Echo) TLS(Dir string) *echo.Echo {
	e.AutoTLSManager.Cache = autocert.DirCache(Dir)
	return e
}
