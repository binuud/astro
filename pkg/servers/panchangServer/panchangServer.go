package panchangServer

import (
	"context"

	astroV1 "github.com/binuud/astro/gen/go/v1/astro"
	"github.com/binuud/astro/pkg/astro"
	"github.com/sirupsen/logrus"
)

const (
	ASTRO_BUCKET_NAME = "astro"
)

type PanchangGRPCServer struct {
	astroV1.UnimplementedPanchangServer
}

func NewGrpcPanchangerver() *PanchangGRPCServer {

	server := &PanchangGRPCServer{}
	return server
}

func (s *PanchangGRPCServer) GetMonth(ctx context.Context, in *astroV1.PanchangMonthGetRequest) (*astroV1.PanchangMonthGetResponse, error) {

	logrus.Infof("Received GetMonth: (%v)", in)

	p, err := astro.NewPanchang(in.BirthData.Tob.Y, in.BirthData.Tob.M, in.BirthData.Tob.D)
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	// p.PrintMonth()

	return &astroV1.PanchangMonthGetResponse{
		Month: p.PanchangMonth,
	}, nil

}
