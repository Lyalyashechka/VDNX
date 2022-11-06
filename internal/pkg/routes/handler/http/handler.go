package http

import (
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
	"github.com/Lyalyashechka/VDNX/internal/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
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
	var param models.PersonInfoRoute

	with := ctx.QueryParam("with")
	param.With = with

	transport := ctx.QueryParam("transport")
	param.Transport = transport

	animals := ctx.QueryParam("animals")
	if animals != "" {
		isHaveAnimals, err := strconv.ParseBool(animals)
		if err != nil {
			h.logger.WithError(errors.Wrap(err, "failed to parse owner param")).
				Errorf("failed to get user events")

			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		if isHaveAnimals {
			param.Animals = 1
		}
	}

	kids := ctx.QueryParam("kids")
	if kids != "" {
		isHaveKids, err := strconv.ParseBool(kids)
		if err != nil {
			h.logger.WithError(errors.Wrap(err, "failed to parse owner param")).
				Errorf("failed to get user events")

			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		if isHaveKids {
			param.Kids = 1
		}
	}

	interestsString := ctx.QueryParam("interests")
	if interestsString != "" {
		param.Interests = strings.Split(interestsString, ",")
	}

	//personalRoutes1, err := h.useCaseRoutes.GetPersonalRoutes(ctx.Request().Context(), param)
	//if err != nil {
	//	h.logger.WithError(err).Errorf("[GetPersonalRoutes] handler")
	//	return ctx.JSON(http.StatusInternalServerError, err)
	//}

	return ctx.NoContent(http.StatusOK)
}
