package render

import (
	"io"

	"github.com/alecthomas/template"

	"github.com/labstack/echo/v4"
	server "github.com/xDarkicex/CCHH-2.0/app/Server"
)

type TemplateRenderer struct {
	Templates *template.Template
}

func init() {
}

//Register Register renderer to server struct
func Register(s *server.Server) *server.Server {
	t := template.Must(template.ParseGlob("app/views/*.html")).Funcs(GetFuncMap())
	t = template.Must(t.ParseGlob("app/views/layouts/*.html"))
	renderer := &TemplateRenderer{
		Templates: t,
	}
	e := s.GetEcho()
	e.Renderer = renderer
	return s.SetEcho(e)
}
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
