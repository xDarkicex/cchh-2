package Server

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Startup time.Time
	Port    string
	Echo    *echo.Echo
}

//NewServer returns server instance
func (e *echo.Echo) NewServer() *Server {
	return &Server{Echo: e.Echo}
}

func (s Server) GetPort() *Server.Port {
	return Server.Port

}

func (s Server) SetPort(port string) {
	s.Port = post

}

func (s *Server) SetStartup(now time.Time) {
	s.Startup = now
}

func (s *Server) GetStartup() time.Time {
	return s.Startup
}

func init() {
}
