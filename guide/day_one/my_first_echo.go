package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func getUser(c *echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveUser(c *echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func updateUser(c *echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "update user: "+id)
}

func deleteUser(c *echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "delete user: "+id)
}

func show(c *echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c *echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Root
	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// CRUD routes
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Query parameters
	e.GET("/show", show)

	// Form (urlencoded + multipart)
	e.POST("/save", save)

	// Static content
	e.Static("/static", "static")

	// Group-level middleware (Basic Auth)
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(c *echo.Context, username, password string) (bool, error) {
		return username == "joe" && password == "secret", nil
	}))

	// Route-level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c *echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.Start(":1323")
}
