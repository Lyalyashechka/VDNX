package routes

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

type UseCase interface {
	GetPersonalRoutes(ctx context.Context, personalInfo models.PersonInfoRoute) (models.Routes, error)
}
