package config_routing

import (
	upload_handler "github.com/Lyalyashechka/VDNX/internal/pkg/upload_data/handler/http"
	"github.com/labstack/echo/v4"
)

type ServerConfigRouting struct {
	UploadHandler *upload_handler.UploadHandler
}

func (cr *ServerConfigRouting) ConfigRouting(router *echo.Echo) {
	router.POST("api/upload/places", cr.UploadHandler.UploadPlaces)
}
