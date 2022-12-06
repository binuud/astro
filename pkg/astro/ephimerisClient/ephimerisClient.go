package ephimerisClient

import (
	"context"
	"os"

	"github.com/dronasys-com/astro/pkg/foundation/config"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var EphimerisClientConnection *EphimerisClient

type EphimerisClient struct {
	serverConn *grpc.ClientConn
}

func NewEphimerisClient() *EphimerisClient {

	// if singleton exists return same
	if EphimerisClientConnection != nil {
		logrus.Trace("Singleton exists, returned existing Mongo DB Connection")
		return EphimerisClientConnection
	}

	serverConn, err := grpc.Dial(config.Config.AstroEphimerisEndpoint, grpc.WithInsecure())
	if err != nil {
		logrus.Error("Cannot Dial GRPC server ", config.Config.AstroEphimerisEndpoint, err)
		os.Exit(1)
	}
	// defer serverConn.Close()
	logrus.Info("Dialed in to GRPC server ", config.Config.AstroEphimerisEndpoint, err)

	EphimerisClientConnection = &EphimerisClient{
		serverConn: serverConn,
	}

	return EphimerisClientConnection

}

func (e *EphimerisClient) GetPlanetsEphimeris(in *astroV1.BirthData) (string, error) {
	client := astroV1.NewEphimerisClient(e.serverConn)
	response, err := client.GetChart(context.TODO(), &astroV1.EphimerisRequest{BirthData: in})
	if err != nil {
		return "", err
	}
	return response.Output, nil
}

func (e *EphimerisClient) GetMoonEphimeris(in *astroV1.BirthData) (string, error) {
	client := astroV1.NewEphimerisClient(e.serverConn)
	response, err := client.GetMoonEphimeris(context.TODO(), &astroV1.EphimerisRequest{BirthData: in})
	if err != nil {
		return "", err
	}
	return response.Output, nil
}

func (e *EphimerisClient) GetPanchangMonth(in *astroV1.BirthData) (string, error) {
	client := astroV1.NewEphimerisClient(e.serverConn)
	response, err := client.GetPanchangMonth(context.TODO(), &astroV1.EphimerisRequest{BirthData: in})
	if err != nil {
		return "", err
	}
	return response.Output, nil
}
