package controllers

import (
	"fmt"
	"net/http"

	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
)

type Application struct{}

func (app Application) Index(c echo.Context) error {
	err := c.Render(http.StatusOK, "index.html", &map[string]interface{}{
		"Title": "Compassionate Care Mobile Health",
		"View":  "index",
	})
	if err != nil {
		fmt.Println(errors.Cause(err))
	}
	return err
}

func (app Application) About(c echo.Context) error {

	return c.Render(http.StatusOK, "about.html", &map[string]interface{}{
		"Title": "About Us",
		"View":  "about",
	})
}
func (app Application) Board(c echo.Context) error {
	return c.Render(http.StatusOK, "board.html", &map[string]interface{}{
		"Title": "About Us",
		"View":  "about",
	})
}

func (app Application) Contact(c echo.Context) error {
	return c.Render(http.StatusOK, "Contact.html", &map[string]interface{}{
		"Title": "Contact",
		"View":  "contact",
	})
}

func (app Application) Gallery(c echo.Context) error {
	return c.Render(http.StatusOK, "gallery.html", &map[string]interface{}{
		"Title": "Gallery",
		"View":  "gallery",
	})
}

func (app Application) Donate(c echo.Context) error {
	return c.Render(http.StatusOK, "donate.html", &map[string]interface{}{
		"Title": "Donate",
		"View":  "donate",
	})
}

func (app Application) GetInvolved(c echo.Context) error {
	return c.Render(http.StatusOK, "get-involved.html", &map[string]interface{}{
		"Title": "Get Involved",
		"View":  "get-involved",
	})
}
func (app Application) Services(c echo.Context) error {
	return c.Render(http.StatusOK, "service.html", &map[string]interface{}{
		"Title": "service",
		"View":  "service",
	})
}

func (app Application) Mission(c echo.Context) error {
	return c.Render(http.StatusOK, "mission.html", &map[string]interface{}{
		"Title": "mission",
		"View":  "mission",
	})
}

func (app Application) Terms(c echo.Context) error {
	return c.Render(http.StatusOK, "terms.html", &map[string]interface{}{
		"Title": "Terms",
		"View":  "terms",
	})
}
