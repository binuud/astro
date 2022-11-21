package panchangServer

import (
	"context"

	"github.com/dronasys-com/astro/pkg/astro"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	ASTRO_BUCKET_NAME = "astro"
)

type grpcAstroServer struct {
	astroV1.UnimplementedPanchangServer
}

func NewGrpcAstroServer(srv *grpc.Server) *grpcAstroServer {

	server := &grpcAstroServer{}
	astroV1.RegisterPanchangServer(srv, server)
	return server
}

func (s *grpcAstroServer) GetMonth(ctx context.Context, in *astroV1.PanchangMonthGetRequest) (*astroV1.PanchangMonthGetResponse, error) {

	logrus.Info("Received GetMonth: ", in)
	p, err := astro.NewPanchangMoon(in.BirthData.Tob.Y, in.BirthData.Tob.M)
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	p.CalculatePanchang()
	// cal := astro.NewCalendarPrinter(p)
	p.Transits.Ephimeris = nil
	return &astroV1.PanchangMonthGetResponse{}, nil

}
