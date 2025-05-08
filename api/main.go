package main

import (
	"net/http"

	"github.com/nowex35/event_management_app/datastore"
	"github.com/nowex35/event_management_app/handler"
	"github.com/nowex35/event_management_app/openapi"

	m "github.com/nowex35/event_management_app/middleware"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var err error

	err = godotenv.Load(".env")
	if err != nil {
		e.Logger.Fatal(err)
	}

	err = datastore.Init()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))
	// e.Use(middleware.BodyDump(m.BodyDump))
	e.Use(echo.MiddlewareFunc(m.Handler()))

	openapi.RegisterHandlersWithBaseURL(e, openapi.NewStrictHandler(handler.Route{
		Version: "1.0.0",
	}, nil), "")

	e.Logger.Fatal(e.Start(":6262"))
}
