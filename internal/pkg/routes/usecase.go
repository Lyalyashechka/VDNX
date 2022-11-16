package routes

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

type UseCase interface {
	GetPersonalRoutes(ctx context.Context, personalInfo models.PersonInfoRoute) (map[string]interface{}, error)
	GetSubPointsRoutes(ctx context.Context, route models.PointsRouteCoord) ([]models.Place, error)
}
