package config_routing

import (
	place_hanlder "github.com/Lyalyashechka/VDNX/internal/pkg/place/delivery/http"
	routes_handler "github.com/Lyalyashechka/VDNX/internal/pkg/routes/handler/http"
	upload_handler "github.com/Lyalyashechka/VDNX/internal/pkg/upload_data/delivery/http"
	"github.com/labstack/echo/v4"
)

type ServerConfigRouting struct {
	UploadHandler *upload_handler.Handler
	PlaceHandler  *place_hanlder.Handler
	RouteHandler  *routes_handler.Handler
}

func (cr *ServerConfigRouting) ConfigRouting(router *echo.Echo) {
	router.POST("api/upload/places", cr.UploadHandler.UploadPlaces)
	router.GET("api/places/get", cr.PlaceHandler.GetAllPlaces)
	router.GET("api/places/get/:id", cr.PlaceHandler.GetPlaceById)
	router.GET("api/routes/personal", cr.RouteHandler.GetPersonalRoutes)
}
