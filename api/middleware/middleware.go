package middleware

import (
	// "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/labstack/echo/v4"
)

func Handler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					c.Set("error", fmt.Sprintf("%+v\n\n%s", r, string(debug.Stack())))
					c.Error(echo.NewHTTPError(http.StatusInternalServerError, errors.New("Internal Server Error")))
				}
			}()
			c.Set("elapsed", time.Now().UnixMicro())
			err := next(c)
			if err != nil {
				c.Set("error", fmt.Sprintf("%+v", err))
			}
			return err
		}
	}
}
