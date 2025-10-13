package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dronasys-com/astro/cmd/db"
	"github.com/dronasys-com/astro/pkg/foundation/bootstrap"
	"github.com/dronasys-com/astro/pkg/servers/astroServer"
	"github.com/dronasys-com/astro/pkg/servers/panchangServer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	grpcPort = flag.Int("grpcPort", 50061, "The Secure Server port, with Astro Calculations")
)

func main() {

	bootstrap.Init()
	db.Init()

	logrus.Info("Astro Calc - Trance !")
	logrus.Info("--------->>>> Starting Astro Calc Microservice  <<<----------")

	flag.Parse()
	secureConn, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// secureServer := grpc.NewServer(
	// 	grpc.UnaryInterceptor(authServer.UnaryInterceptor))
	secureServer := grpc.NewServer()

	astroServer.NewGrpcAstroServer(secureServer)
	panchangServer.NewGrpcPanchangerver(secureServer)

	logrus.Infof("Astro Calculations Server *) (* listening on (%v)", secureConn.Addr())
	if err := secureServer.Serve(secureConn); err != nil {
		logrus.Fatal("failed to serve: ", err)
	}

}
