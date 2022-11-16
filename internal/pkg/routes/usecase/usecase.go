package usecase

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/data_science_manager/handler/http"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
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
	PlaceUC   place.UseCase
	logger    *logrus.Logger
}

func New(dsManager *http.DataScienceServerHandler, placeUC place.UseCase, logger *logrus.Logger) routes.UseCase {
	return &UseCase{
		DsManager: dsManager,
		logger:    logger,
		PlaceUC:   placeUC,
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

func (uc *UseCase) GetSubPointsRoutes(ctx context.Context, route models.PointsRouteCoord) ([]models.Place, error) {
	var ret []models.Place
	places, err := uc.PlaceUC.GetAllPlaces(ctx, place.GetAllPlacesParam{IsEvent: models.Bool{Value: false, Defined: true}})
	if err != nil {
		uc.logger.WithError(err).Errorf("error get all places")
		return nil, err
	}
	for i := 0; i < len(route.Coords)-1; i++ {
		placesBetweenTwoPoint := uc.TraceSubPointRoutes(ctx, route.Coords[i], route.Coords[i+1], places)
		ret = append(ret, placesBetweenTwoPoint...)
	}

	return ret, nil
}

func (uc *UseCase) TraceSubPointRoutes(ctx context.Context, firstPoint models.PointRouteCoord, secondPoint models.PointRouteCoord, allPlaces []models.Place) []models.Place {
	var ret []models.Place

	distance := firstPoint.CalcDistance(secondPoint)
	distanceBetweenPointTrace := distance / 4
	normalVector := firstPoint.CalcNormalVectorToPoint(secondPoint)
	for i := 0; i < 5; i++ {
		nextPoint := models.PointRouteCoord{
			Longitude: firstPoint.Longitude + float64(i)*distanceBetweenPointTrace*normalVector.Longitude,
			Latitude:  firstPoint.Latitude + float64(i)*distanceBetweenPointTrace*normalVector.Latitude,
		}

		placeNearPoint := uc.GetPlacesNearOnePoint(ctx, nextPoint, allPlaces)

		for _, v := range placeNearPoint {
			isExist := false
			for _, v1 := range placeNearPoint {
				if v.Id == v1.Id {
					isExist = true
					break
				}
			}
			if !isExist {
				ret = append(ret, v)
			}
		}
	}
	return ret
}

func (uc *UseCase) GetPlacesNearOnePoint(ctx context.Context, point models.PointRouteCoord, allPlaces []models.Place) []models.Place {
	var ret []models.Place
	for _, v := range allPlaces {
		if distance := point.CalcDistanceMetr(models.PointRouteCoord{
			Latitude:  v.Coordinates[0],
			Longitude: v.Coordinates[1],
		}); distance < 50 {
			ret = append(ret, v)
		}
	}

	return ret
}
