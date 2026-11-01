package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/dronasys-com/astro/pkg/servers/astroServer"
	"github.com/dronasys-com/astro/pkg/servers/panchangServer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	server_grpc_port = flag.Int("grpc_port", 9090, "Astro Calculations GRPC Serverport, no token required unsecured")
	server_http_port = flag.Int("http_port", 9080, "Astro Calculations HTTP Server port, no token required unsecured")
)

// starting the http server from the GPRC generated code
func runHtppServer() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		fmt.Sprintf("0.0.0.0:%d", *server_grpc_port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwMux := runtime.NewServeMux()

	err = astroV1.RegisterAstroHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	err = astroV1.RegisterPanchangHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *server_http_port),
		Handler: gwMux,
	}

	log.Infof("HTTP server started on port %v", *server_http_port)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatalln(gwServer.ListenAndServe())

}

func runGRPCServer() {

	insecureConn, err := net.Listen("tcp", fmt.Sprintf(":%d", *server_grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	insecureServer := grpc.NewServer()

	astroV1.RegisterAstroServer(insecureServer, astroServer.NewGrpcAstroServer())

	astroV1.RegisterPanchangServer(insecureServer, panchangServer.NewGrpcPanchangerver())

	// Register reflection service on gRPC server.
	reflection.Register(insecureServer)

	log.Infof("Astro Calculations GRPC Server *) (* listening on (%v)", insecureConn.Addr())
	log.Fatalln(insecureServer.Serve(insecureConn))

}

func main() {

	flag.Parse()

	log.Info("Astro Calc - Trance !")
	log.Info("--------->>>> Starting Astro Calc Microservice  <<<----------")

	go func() {
		runGRPCServer()
	}()

	runHtppServer()

}
