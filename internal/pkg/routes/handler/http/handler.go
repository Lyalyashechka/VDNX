package http

import (
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
	"github.com/Lyalyashechka/VDNX/internal/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger        *logrus.Logger
	useCaseRoutes routes.UseCase
	useCasePlace  place.UseCase
}

func New(logger *logrus.Logger, useCaseRoute routes.UseCase, useCasePlaces place.UseCase) *Handler {
	return &Handler{
		logger:        logger,
		useCaseRoutes: useCaseRoute,
		useCasePlace:  useCasePlaces,
	}
}

func (h *Handler) GetPersonalRoutes(ctx echo.Context) error {
	//personalRoutes1, err := h.useCaseRoutes.GetPersonalRoutes(ctx.Request().Context(), models.PersonInfoRoute{
	//	With:      "Компанией",
	//	Children:  true,
	//	Interests: []string{"Павильон"},
	//	Transport: "Пешком",
	//	Time:      "2h",
	//})
	//if err != nil {
	//	h.logger.WithError(err).Errorf("[GetPersonalRoutes] handler")
	//	return ctx.JSON(http.StatusInternalServerError, err)
	//}
	//h.logger.Info(personalRoutes1)

	place1, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 331)
	place2, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 3261)
	place3, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 14682)
	place4, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 329)

	var route1 models.Route
	route1.Places = append(route1.Places, place1)
	route1.Places = append(route1.Places, place2)
	route1.Places = append(route1.Places, place3)
	route1.Places = append(route1.Places, place4)

	place5, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 412)
	place6, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 2271)
	place7, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 6780)
	place8, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 324)

	var route2 models.Route
	route2.Places = append(route2.Places, place5)
	route2.Places = append(route2.Places, place6)
	route2.Places = append(route2.Places, place7)
	route2.Places = append(route2.Places, place8)

	place9, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 2312)
	place10, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 2265)
	place11, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 428)
	place12, _ := h.useCasePlace.GetPlaceById(ctx.Request().Context(), 3296)

	var route3 models.Route
	route3.Places = append(route3.Places, place9)
	route3.Places = append(route3.Places, place10)
	route3.Places = append(route3.Places, place11)
	route3.Places = append(route3.Places, place12)

	var personalRoutes models.Routes

	personalRoutes.Routes = append(personalRoutes.Routes, route3)
	personalRoutes.Routes = append(personalRoutes.Routes, route2)
	personalRoutes.Routes = append(personalRoutes.Routes, route1)

	return ctx.JSON(http.StatusOK, personalRoutes)
}
