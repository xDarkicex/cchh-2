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
type Config struct {
	Address string
}

var s *Server

//var CURRENT_SERVER = s
var SSL_CERT_LOCATION = ".cache/golang-autocert"

func init() {

}

func (s *Server) Initialize(config *Config) *Server {
	s.TLS(SSL_CERT_LOCATION).MiddleWare()
	if config != nil {
		s.SetPort(config.Address)
	}
	if config == nil {
		s.SetPort(":8080")
	}
	return s
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
func (s *Server) SetPort(port string) (server *Server) {
	s.Port = port
	return server
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

func (s *Server) GetRenderer() *echo.Renderer {
	return &s.Echo.Renderer
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
	s.Echo.Use(middleware.HTTPSRedirect())
	s.Echo.Use(middleware.Gzip())
	return s
}

//TLS configures TLS Settings
func (s *Server) TLS(Dir string) *Server {

	domains := []string{"compassionatecaremobileclinic.org", "www.compassionatecaremobileclinic.org"}
	s.Echo.AutoTLSManager.HostPolicy = autocert.HostWhitelist(domains...)
	s.Echo.AutoTLSManager.Cache = autocert.DirCache(Dir)
	return s
}
