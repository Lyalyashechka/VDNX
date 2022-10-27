package usecase

import (
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
	"github.com/sirupsen/logrus"
)

type UseCase struct {
	logger     *logrus.Logger
	repository place.Repository
}

func New(logger *logrus.Logger, repository place.Repository) *UseCase {
	return &UseCase{
		logger:     logger,
		repository: repository,
	}
}
