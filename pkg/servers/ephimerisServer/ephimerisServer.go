package ephimerisServer

import (
	"context"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcEphimerisServer struct {
	astroV1.UnimplementedEphimerisServer
}

func NewGrpcEphimerisServer(srv *grpc.Server) *grpcEphimerisServer {

	server := &grpcEphimerisServer{}
	astroV1.RegisterEphimerisServer(srv, server)
	return server
}

func (s *grpcEphimerisServer) GetChart(ctx context.Context, in *astroV1.EphimerisRequest) (*astroV1.EphimerisResponse, error) {

	logrus.Info("Received GetChart: ", in)

	e := NewEphimeris(in.BirthData)
	e.GetPlanetsEphimeris()
	out, err := e.ExecCommand()

	if err != nil {
		logrus.Errorf("Cannot get ephimeris data (%v)", err)
		return nil, err
	}
	// logrus.Info("Output %v", string(out))

	return &astroV1.EphimerisResponse{
		Output: string(out),
	}, nil

}

func (s *grpcEphimerisServer) GetPanchangMonth(ctx context.Context, in *astroV1.EphimerisRequest) (*astroV1.EphimerisResponse, error) {

	logrus.Info("Received GetPanchangMonth: ", in)

	e := NewEphimeris(in.BirthData)
	e.GetPanchangMonth()
	out, err := e.ExecCommand()

	if err != nil {
		logrus.Errorf("Cannot get ephimeris data (%v)", err)
		return nil, err
	}
	// logrus.Info("Output %v", string(out))

	return &astroV1.EphimerisResponse{
		Output: string(out),
	}, nil

}

func (s *grpcEphimerisServer) GetMoonEphimeris(ctx context.Context, in *astroV1.EphimerisRequest) (*astroV1.EphimerisResponse, error) {

	logrus.Info("Received GetMoonEphimeris: ", in)

	e := NewEphimeris(in.BirthData)
	e.GetMoonEphimeris()
	out, err := e.ExecCommand()

	if err != nil {
		logrus.Errorf("Cannot get ephimeris data (%v)", err)
		return nil, err
	}
	// logrus.Info("Output %v", string(out))

	return &astroV1.EphimerisResponse{
		Output: string(out),
	}, nil

}
