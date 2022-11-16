package http

import (
	"encoding/json"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	upload_data "github.com/Lyalyashechka/VDNX/internal/pkg/upload_data"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase upload_data.UseCase
	logger  *logrus.Logger
}

func New(logger *logrus.Logger, useCase upload_data.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (uh *Handler) UploadPlaces(ctx echo.Context) error {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		uh.logger.WithError(err).Errorf("[UploadPlaces] error readAll")
		return ctx.NoContent(http.StatusBadRequest)
	}

	var jsonResp interface{}
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		uh.logger.WithError(err).Errorf("[UploadPlaces] error Unmarshal")
		return ctx.NoContent(http.StatusBadRequest)
	}

	var isEvent bool
	isEventString := ctx.QueryParam("is_event")
	if isEventString != "" {
		isEvent, err = strconv.ParseBool(isEventString)
		if err != nil {
			uh.logger.WithError(errors.Wrap(err, "failed to parse event param"))
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
	}

	var places []models.Place
	jsonMap := jsonResp.(map[string]interface{})
	for _, v := range jsonMap {
		var place models.Place
		mapstructure.Decode(v, &place)
		place.IsEvent = isEvent
		places = append(places, place)
	}

	err = uh.useCase.CreatePlaces(ctx.Request().Context(), places)
	if err != nil {
		uh.logger.WithError(err).Errorf("[UploadPlaces] error create place")
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}
