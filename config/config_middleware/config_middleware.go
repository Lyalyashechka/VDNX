package config_middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var (
	allowOrigins  = []string{"*"}
	allowMethods  = []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions}
	allowHeaders  = []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}
	exposeHeaders = []string{"Authorization"}
)

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				context.NoContent(http.StatusInternalServerError)
			}
		}()

		return next(context)
	}
}

func GetCORSConfigStruct() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowMethods:     allowMethods,
		AllowHeaders:     allowHeaders,
		ExposeHeaders:    exposeHeaders,
		AllowCredentials: true,
	}
}

func ConfigMiddleware(router *echo.Echo) {
	router.Use(
		Recover,
		middleware.CORSWithConfig(GetCORSConfigStruct()),
	)
}
