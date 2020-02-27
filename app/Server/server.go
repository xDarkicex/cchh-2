package server

import (
	"fmt"
	"log"
	"time"

	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controllers "github.com/xDarkicex/CCHH-2.0/app/controllers/application"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	Echo    *echo.Echo
	Startup time.Time
	Port    string
}

var s *Server

//var CURRENT_SERVER = s
var SSL_CERT_LOCATION = "/var/www/.cache"

func init() {

}

func GetServer() *Server {
	return s
}

var app controllers.Application

//Routes loads all routes withing application
func (s *Server) SetRoutes() *Server {
	s.Echo.GET("/", app.Index)
	s.Echo.GET("/contact", app.Contact)
	s.Echo.GET("/about", app.About)
	s.Echo.GET("/gallery", app.Gallery)
	s.Echo.GET("/donate", app.Donate)
	s.Echo.GET("/get-involved", app.GetInvolved)
	s.Echo.GET("/services", app.Services)
	s.Echo.GET("/terms", app.Terms)
	s.Echo.GET("/mission", app.Mission)
	s.Echo.GET("/board", app.Board)
	s.Echo.Static("/static", "public")
	return s.SetEcho(s.Echo)
}

//NewServer returns server instance
func NewServer() *Server {
	return &Server{
		Echo:    NewEcho(),
		Startup: time.Now(),
	}
}

//GetPort returns port for  server
func (s *Server) GetPort() string {
	return s.Port
}

//SetPort sets port for server
func (s *Server) SetPort(port string) {
	s.Port = port
}

//SetStartup takes time.Now() for start time of server
func (s *Server) SetStartupTime() {
	s.Startup = time.Now()
}

//GetStartup Returns the initial startup time for server
func (s *Server) GetStartupTime() time.Time {
	return s.Startup
}

//ServerRunTime  measure server uptime
func (s *Server) ServerRunTime() (Duration float64, endTime string) {
	return time.Since(s.GetStartupTime()).Hours(), time.Now().Format(time.Stamp)
}

func (s *Server) Close(err error) {
	if err != nil {
		D, endTime := s.ServerRunTime()
		fmt.Println("Duration: ", D, "s shutoff at: ", endTime)
		log.Fatal(errors.Cause(err))
	}
}

//SetStartup takes time.Now() for start time of server
func (s *Server) SetEcho(e *echo.Echo) *Server {
	s.Echo = e
	return s
}

func (s *Server) GetEcho() *echo.Echo {
	return s.Echo
}

func (s *Server) UpdateEcho(e *echo.Echo) {
	s.Echo = e
}

//NewEcho Loads up the Echo instance
func NewEcho() *echo.Echo {
	return echo.New()
}

//MiddleWare Attach middle ware
func (s *Server) MiddleWare() *Server {
	s.Echo.Use(middleware.Recover())
	s.Echo.Use(middleware.Logger())
	return s
}

//TLS configures TLS Settings
func (s *Server) TLS(Dir string) *Server {
	s.Echo.AutoTLSManager.Cache = autocert.DirCache(Dir)
	return s
}
