package models

import "math"

type Route struct {
	Places []Place `json:"places"`
}

type Routes struct {
	Routes []Route `json:"routes"`
}

type PointsRouteCoord struct {
	Coords []PointRouteCoord `json:"points_route"`
}

type PointRouteCoord struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (point *PointRouteCoord) CalcDistance(to PointRouteCoord) float64 {
	return math.Sqrt(math.Pow((to.Longitude-point.Longitude), 2) + math.Pow((to.Latitude-point.Longitude), 2))
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (point *PointRouteCoord) CalcDistanceMetr(to PointRouteCoord) float64 {
	earthRadiusKm := 6371.0 * 1000
	dLat := degreesToRadians(to.Latitude - point.Latitude)
	dLon := degreesToRadians(to.Longitude - point.Longitude)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(point.Latitude)*math.Cos(to.Latitude)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusKm * c
}

func (pint *PointRouteCoord) CalcNormalVectorToPoint(to PointRouteCoord) PointRouteCoord {
	vector := PointRouteCoord{
		Longitude: to.Longitude - pint.Longitude,
		Latitude:  to.Latitude - pint.Latitude,
	}

	length := math.Sqrt(math.Pow(vector.Latitude, 2) + math.Pow(vector.Longitude, 2))

	return PointRouteCoord{
		Longitude: vector.Longitude / length,
		Latitude:  vector.Latitude / length,
	}
}
