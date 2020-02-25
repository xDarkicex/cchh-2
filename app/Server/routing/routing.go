package routing

//Routes loads all routes withing application
func (e *echo.Echo) Routes() *echo.Echo {
	e.GET("/", app.Index)
	e.GET("/contact", app.Contact)
	e.GET("/about", app.About)
	e.GET("/gallery", app.Gallery)
	e.GET("/donate", app.Donate)
	e.GET("/get-involved", app.GetInvolved)
	e.GET("/services", app.Services)
	e.GET("/terms", app.Terms)
	e.GET("/mission", app.Mission)
	e.GET("/board", app.Board)
	e.Static("/static", "assets")
	return e
}
