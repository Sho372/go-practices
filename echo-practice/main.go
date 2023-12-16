package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name  string `json:"name,omitempty" query:"name"`
	Email string `json:"email" query:"email"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	// /show?team=x-men&member=wolverine
	e.GET("/show", show)
	// /user?user=shohei&email=foo@example.com
	e.GET("/user", user)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func user(c echo.Context) error {
	u := new(User)
	qp := c.QueryParams()
	fmt.Printf("QueryParamas: %s\n", qp)
	qs := c.QueryString()
	fmt.Printf("QueryString: %s\n", qs)
	pn := c.ParamNames()
	fmt.Printf("ParamNames: %s\n", pn)
	fmt.Printf("BEFORE BIND User.name: [%s], User.email: [%s]\n", u.Name, u.Email)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
