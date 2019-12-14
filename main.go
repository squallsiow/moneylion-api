package main

import (
	"context"
	"html/template"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moneylion-api/handler"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// Create singleton handler
	hdl, err := handler.New()
	if err != nil {
		panic(err)
	}

	// set templates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	// create new echo web server
	api := echo.New()

	api.Renderer = t

	// Middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Secure())
	api.Use(middleware.Recover())

	//  route
	api.GET("/", hdl.Hello)

	// api.Logger.Fatal(api.Start(":8080"))

	// Start server
	go func() {
		if err := api.Start(":8080"); err != nil {
			api.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := api.Shutdown(ctx); err != nil {
		api.Logger.Fatal(err)
	}

}
