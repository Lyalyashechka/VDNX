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

func (h *DataScienceServerHandler) GetPersonalRoutes(route models.PersonInfoRoute) (map[string]interface{}, error) {
	jsonMarshal, err := json.Marshal(route)
	if err != nil {
		h.logger.WithError(err).Errorf("[GetPersonalRoutes] error marshal json")
		return nil, err
	}

	resp, err := http.Post(h.Address+PERSONAL_ROUTES, "application/json", bytes.NewBuffer(jsonMarshal))
	if err != nil {
		h.logger.WithError(err).Errorf("[GetPersonalRoutes] error make request")
		return nil, err
	}

	var answer map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&answer)
	if err != nil {
		h.logger.WithError(err).Errorf("[GetPersonalRoutes] error decode request")
		return nil, err
	}

	return answer, nil
}
