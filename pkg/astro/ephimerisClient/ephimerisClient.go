package ephimerisClient

import (
	"context"
	"log"
	"os"

	astroV1 "github.com/binuud/astro/gen/go/v1/astro"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var EphimerisClientConnection *EphimerisClient

type EphimerisClient struct {
	serverConn *grpc.ClientConn
}

func NewEphimerisClient() *EphimerisClient {

	// if singleton exists return same
	if EphimerisClientConnection != nil {
		logrus.Trace("Singleton exists, returned existing EphimerisClientConnection")
		return EphimerisClientConnection
	}

	ephimerisEndpoint := os.Getenv("ASTRO_EPHIMERIS_ENDPOINT")

	serverConn, err := grpc.NewClient(
		ephimerisEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// defer serverConn.Close()
	logrus.Info("Dialed in to GRPC server ", ephimerisEndpoint)

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
