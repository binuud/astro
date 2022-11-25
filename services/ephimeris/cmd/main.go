package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dronasys-com/astro/pkg/servers/ephimerisServer"
	"github.com/dronasys-com/trance-microservices/pkg/foundation/bootstrap"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	grpcPort = flag.Int("grpcPort", 50062, "The Secure Server port, with Astro Ephimeris")
)

func main() {

	bootstrap.Init()

	logrus.Info("Astro Ephimeris - Trance !")
	logrus.Info("--------->>>> Starting Astro Ephimeris Microservice  <<<----------")

	flag.Parse()
	secureConn, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// secureServer := grpc.NewServer(
	// 	grpc.UnaryInterceptor(authServer.UnaryInterceptor))
	secureServer := grpc.NewServer()

	ephimerisServer.NewGrpcEphimerisServer(secureServer)

	logrus.Infof("Astro Ephimeris Server *) (* listening on (%v)", secureConn.Addr())
	if err := secureServer.Serve(secureConn); err != nil {
		logrus.Fatal("failed to serve: ", err)
	}

}
