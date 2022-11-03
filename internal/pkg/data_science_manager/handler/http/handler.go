package http

import (
	"bytes"
	json "encoding/json"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	PERSONAL_ROUTES = "/routes/personal"
)

type DataScienceServerHandler struct {
	Address string
	logger  *logrus.Logger
}

func New(address string) *DataScienceServerHandler {
	return &DataScienceServerHandler{
		Address: address,
	}
}

func (h *DataScienceServerHandler) GetPersonalRoutes(route models.PersonInfoRoute) (models.Routes, error) {
	jsonMarshal, err := json.Marshal(route)
	if err != nil {
		h.logger.WithError(err).Errorf("[GetPersonalRoutes] error marshal json")
		return models.Routes{}, err
	}

	resp, err := http.Post(h.Address+PERSONAL_ROUTES, "application/json", bytes.NewBuffer(jsonMarshal))
	if err != nil {
		h.logger.WithError(err).Errorf("[GetPersonalRoutes] error make request")
		return models.Routes{}, err
	}

	var personalRoutes models.Routes

	err = json.NewDecoder(resp.Body).Decode(&personalRoutes)
	if err != nil {
		h.logger.WithError(err).Errorf("[GetPersonalRoutes] error decode request")
		return models.Routes{}, err
	}

	return personalRoutes, nil
}
