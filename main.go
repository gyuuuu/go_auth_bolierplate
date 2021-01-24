package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	h := &handler{}
	e.POST("/login", h.login)

	a := e.Group("/auth")

	a.Use(jwtconfig)
	a.GET("/private", h.private)
	a.GET("/admin", h.private, isAdmin)

	e.Logger.Fatal(e.Start(":1323"))
}
