package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/nowex35/event_management_app/openapi"
)

type Route struct {
	Version string
}

func (h Route) Echo(ctx echo.Context, request openapi.EchoRequestObject) (openapi.EchoResponseObject, error) {

	return openapi.Echo200JSONResponse{
		Code: 0,
	}, nil
}
