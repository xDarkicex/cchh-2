package main

import (
	"fmt"
	"log"

	"github.com/xDarkicex/CCHH-2.0/app/Server"
	"github.com/xDarkicex/CCHH-2.0/app/helpers/render"
)

var s *server.Server

func init() {
	s = server.NewServer()
	s.SetPort(":3002")
	s.MiddleWare()
	s.TLS(server.SSL_CERT_LOCATION)
	s = render.Register(s)
	s = s.SetRoutes()

}

func main() {
	fmt.Println(s.GetStartupTime())
	log.Fatal(s.Echo.Start(s.GetPort()))
}
