package http

import (
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase place.UseCase
	logger  *logrus.Logger
}

func New(logger *logrus.Logger, useCase place.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (ph *Handler) GetAllPlaces(ctx echo.Context) error {
	var getAllPlacesParam place.GetAllPlacesParam
	isEventString := ctx.QueryParam("is_event")
	if isEventString != "" {
		isEvent, err := strconv.ParseBool(isEventString)
		if err != nil {
			ph.logger.WithError(errors.Wrap(err, "failed to parse event param"))
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		getAllPlacesParam.IsEvent = models.DefinedBool(isEvent)
	}
	
	places, err := ph.useCase.GetAllPlaces(ctx.Request().Context(), getAllPlacesParam)
	if err != nil {
		ph.logger.WithError(err).Errorf("[GetAllPlaces] error usecase")
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, places)
}
