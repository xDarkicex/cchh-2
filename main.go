package main

import (
	"github.com/xDarkicex/CCHH-2.0/app/helpers/compile"
	"github.com/xDarkicex/CCHH-2.0/app/server"
	"github.com/xDarkicex/CCHH-2.0/app/server/action"
)

var s *server.Server

func init() {
	compile.Assets()
	s = server.NewServer()
	s.Initialize(&server.Config{Address: ":443"})
	action.NewAction("Templating").Register(s)
	action.NewAction("Routing").SetRoutes(s)
}

func main() {
	//go func() {
	//		s.Echo.Logger.Fatal(s.Echo.Start(":80"))
	//}()
	s.Echo.Logger.Fatal(s.Echo.StartAutoTLS(s.GetPort()))
}
