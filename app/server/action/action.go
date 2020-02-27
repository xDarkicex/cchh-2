package action

import (
	"github.com/xDarkicex/CCHH-2.0/app/controllers/application"
	"github.com/xDarkicex/CCHH-2.0/app/server"
)

var app controllers.Application

type Action struct {
	Name string
}

func NewAction() *Action {
	return &Action{}
}

//Actions Actions list
//var Actions = make([]Action, 0)

//Routes loads all routes withing application
func (a *Action) SetRoutes(s *server.Server) *server.Server {
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
	s.Echo.POST("/contact", app.Contact)
	s.Echo.Static("/static", "public")
	return s.SetEcho(s.Echo)
}
