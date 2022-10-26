package http

import (
	"encoding/json"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	upload_data "github.com/Lyalyashechka/VDNX/internal/pkg/upload_data"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type UploadHandler struct {
	useCase upload_data.UseCase
	logger  *logrus.Logger
	db      *gorm.DB
}

func New(logger *logrus.Logger, db *gorm.DB, useCase upload_data.UseCase) *UploadHandler {
	return &UploadHandler{
		useCase: useCase,
		logger:  logger,
		db:      db,
	}
}

func (uh *UploadHandler) UploadPlaces(ctx echo.Context) error {
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

	var places []models.Place
	jsonMap := jsonResp.(map[string]interface{})
	for _, v := range jsonMap {
		var place models.Place
		mapstructure.Decode(v, &place)
		places = append(places, place)
	}

	err = uh.useCase.CreatePlaces(ctx.Request().Context(), places)
	if err != nil {
		uh.logger.WithError(err).Errorf("[UploadPlaces] error create place")
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}
