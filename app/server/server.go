package server

import (
	"fmt"
	"log"
	"time"

	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	s.Echo.Use(middleware.Gzip())
	return s
}

//TLS configures TLS Settings
func (s *Server) TLS(Dir string) *Server {
	s.Echo.AutoTLSManager.Cache = autocert.DirCache(Dir)
	return s
}
