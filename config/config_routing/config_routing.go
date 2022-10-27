package config_routing

import (
	place_hanlder "github.com/Lyalyashechka/VDNX/internal/pkg/place/delivery/http"
	upload_handler "github.com/Lyalyashechka/VDNX/internal/pkg/upload_data/delivery/http"
	"github.com/labstack/echo/v4"
)

type ServerConfigRouting struct {
	UploadHandler *upload_handler.Handler
	PlaceHandler  *place_hanlder.Handler
}

func (cr *ServerConfigRouting) ConfigRouting(router *echo.Echo) {
	router.POST("api/upload/places", cr.UploadHandler.UploadPlaces)
	router.GET("api/places/get", cr.PlaceHandler.GetAllPlaces)
}
