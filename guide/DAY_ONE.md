# Echo Journey — From Setup to ..

## Start by creating a fresh folder

```bash
mkdir day_one
cd day_one
```

---

## Initialize your Go module

```bash
go mod init day_one
```

**Why:** gives your project a name so Go can manage dependencies and imports.

*Files created:*

* `go.mod` → defines the module and tracks dependencies
* `go.sum` → verifies dependency integrity

---

## Add Echo

```bash
go get github.com/labstack/echo/v5
```

**Why:** brings the Echo framework into your project.

---

## Clean dependencies

```bash
go mod tidy
```

**Why:** keeps only what is needed and syncs your module files.

---

## Create your first server

```bash
touch my_first_echo.go
```

Paste:

```go
package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Start(":1323")
}
```

---

## Run the server

```bash
go run my_first_echo.go
```

Open:

```
http://localhost:1323
```

You should see:

```
Hello, World!
```

---

## Routing

```go
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```

**What this means:** different HTTP methods mapped to different actions on the same resource.

---

## Path Parameters

```go
func getUser(c *echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
```

Open:

```
http://localhost:1323/users/joe
```

You should see:

```
joe
```

---

## Query Parameters

```go
func show(c *echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
```

Open:

```
http://localhost:1323/show?team=x-men&member=wolverine
```

You should see:

```
team:x-men, member:wolverine
```

---

## Form (application/x-www-form-urlencoded)

```go
func save(c *echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
```

Run:

```bash
curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save
```

Response:

```
name:Joe Smith, email:joe@labstack.com
```

---

## Form (multipart/form-data)

```go
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
```

Run:

```bash
curl -F "name=Joe Smith" -F "avatar=@/path/to/avatar.png" http://localhost:1323/save
```

---

## Handling Request Data

```go
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

e.POST("/users", func(c *echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
})
```

**What this does:** maps incoming request data into a Go struct automatically.

---

## Static Content

```go
e.Static("/static", "static")
```

**What this does:** serves files from a folder.

---

## Middleware

```go
e.Use(middleware.RequestLogger())
e.Use(middleware.Recover())
```

**What this does:** runs logic before/after requests globally.

---

### Group-level

```go
g := e.Group("/admin")
g.Use(middleware.BasicAuth(func(c *echo.Context, username, password string) (bool, error) {
	return username == "joe" && password == "secret", nil
}))
```

---

### Route-level

```go
track := func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		println("request to /users")
		return next(c)
	}
}

e.GET("/users", func(c *echo.Context) error {
	return c.String(http.StatusOK, "/users")
}, track)
```

---

## Where you are now

* Running a Go HTTP server
* Defining routes with different HTTP methods
* Reading path, query, and form data
* Handling file uploads
* Binding request data into structs
* Using middleware

*This is already enough to build a real API.*

---

## Next

Structure your project and separate handlers, services, and routes.
