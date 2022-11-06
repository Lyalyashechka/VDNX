package usecase

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/data_science_manager/handler/http"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/routes"
	"github.com/sirupsen/logrus"
)

type GetPersonalRoutesParam struct {
	With      string
	Animals   int
	Kids      int
	Interests []string
	Transport string
	Position  []float64
}

type UseCase struct {
	DsManager *http.DataScienceServerHandler
	logger    *logrus.Logger
}

func New(dsManager *http.DataScienceServerHandler, logger *logrus.Logger) routes.UseCase {
	return &UseCase{
		DsManager: dsManager,
		logger:    logger,
	}
}

func (uc *UseCase) GetPersonalRoutes(ctx context.Context, personalInfo models.PersonInfoRoute) (map[string]interface{}, error) {
	personalRoutes, err := uc.DsManager.GetPersonalRoutes(personalInfo)
	if err != nil {
		uc.logger.WithError(err).Errorf("[GetPersonalRoutes] uc")
		return nil, nil
	}
	return personalRoutes, nil
}
