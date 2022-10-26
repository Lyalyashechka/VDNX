package usecase

import (
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
	"github.com/Lyalyashechka/VDNX/internal/pkg/upload_data"
	"github.com/sirupsen/logrus"
)

type UseCase struct {
	logger          *logrus.Logger
	placeRepository place.Repository
}

func New(logger *logrus.Logger, placeRepository place.Repository) upload_data.UseCase {
	return &UseCase{
		placeRepository: placeRepository,
		logger:          logger,
	}
}
