package main

import (
	"log"

	"github.com/juju/errors"
	"github.com/xDarkicex/CCHH-2.0/app/helpers/render"
	"github.com/xDarkicex/CCHH-2.0/app/server"
	"github.com/xDarkicex/CCHH-2.0/app/server/action"
)

var s *server.Server

func init() {

	s = server.NewServer()
	s.SetPort(":3002")
	s.MiddleWare()
	s.TLS(server.SSL_CERT_LOCATION)
	s = render.Register(s)
	s = action.NewAction().SetRoutes(s)
}

func main() {
	log.Fatal(errors.Cause(s.Echo.Start(s.GetPort())))
}
