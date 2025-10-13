package citiesServer

import (
	"context"

	citiesV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/sirupsen/logrus"
)

type grpcCitiesServer struct {
	db *CityDB
	citiesV1.UnimplementedCitiesServer
}

func NewGrpcCitiesServer() *grpcCitiesServer {

	db := NewCityDB()

	server := &grpcCitiesServer{
		db: db,
		// citiesV1.UnimplementedCitiesServer{},
	}

	return server
}

func (s *grpcCitiesServer) Get(ctx context.Context, in *citiesV1.CitiesGetRequest) (*citiesV1.CitiesGetResponse, error) {

	logrus.Info("Received: Cities Get request ", in)
	city, err := s.db.GetCityInfo(int(in.Index))
	if err != nil {
		return nil, err
	}

	return &citiesV1.CitiesGetResponse{
		Details: city,
	}, nil

}

func (s *grpcCitiesServer) Search(ctx context.Context, in *citiesV1.CitiesSearchRequest) (*citiesV1.CitiesSearchResponse, error) {

	logrus.Info("Received: Cities Search request ", in)
	cities, err := s.db.SearchCity(in.Name, in.CountryCode)
	if err != nil {
		return nil, err
	}

	return &citiesV1.CitiesSearchResponse{
		Cities: cities,
	}, nil

}
