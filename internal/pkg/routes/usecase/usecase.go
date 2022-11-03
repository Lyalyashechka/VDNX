package usecase

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/data_science_manager/handler/http"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/routes"
	"github.com/sirupsen/logrus"
)

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

func (uc *UseCase) GetPersonalRoutes(ctx context.Context) (models.Routes, error) {
	personalRoutes, err := uc.DsManager.GetPersonalRoutes(models.PersonInfoRoute{
		With: "",
	})
	if err != nil {
		uc.logger.WithError(err).Errorf("[GetPersonalRoutes] uc")
		return models.Routes{}, nil
	}
	return personalRoutes, nil
}
