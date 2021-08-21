package main

import (
	"context"
	"html/template"
	"io"
	"net/http"

	"github.com/moneylion-api/env"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moneylion-api/bootstrap"
	"github.com/moneylion-api/handler"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
func main() {

	// Initialize environment variable
	env.Init()
	// Create singleton handler
	ctx := context.Background()
	bs := bootstrap.New(ctx)
	hdl, err := handler.New(bs)
	if err != nil {
		panic(err)
	}

	// set templates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	// create new echo web server
	api := echo.New()
	api.Validator = &CustomValidator{validator: validator.New()}

	api.Renderer = t

	// Middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Secure())
	api.Use(middleware.Recover())

	//  route
	api.GET("/", hdl.Hello)
	api.POST("/user", hdl.CreateUser)
	api.POST("/feature-new", hdl.CreateFeature)

	api.GET("/feature", hdl.GetFeatureAccess)
	api.POST("/feature", hdl.AccessFeature)

	api.Logger.Fatal(api.Start(":8080"))

}
