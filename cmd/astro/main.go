package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/dronasys-com/astro/cmd/db"
	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/dronasys-com/astro/pkg/servers/astroServer"
	"github.com/dronasys-com/astro/pkg/servers/panchangServer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	err = astroV1.RegisterAstroHandlerServer(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *server_http_port),
		Handler: gwMux,
	}

	log.Info("HTTP server started")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatalln(gwServer.ListenAndServe())

}

func runGRPCServer() {

	secureConn, err := net.Listen("tcp", fmt.Sprintf(":%d", *server_grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// secureServer := grpc.NewServer(
	// 	grpc.UnaryInterceptor(authServer.UnaryInterceptor))
	secureServer := grpc.NewServer()

	astroServer.NewGrpcAstroServer(secureServer)
	panchangServer.NewGrpcPanchangerver(secureServer)

	log.Infof("Astro Calculations Server *) (* listening on (%v)", secureConn.Addr())
	if err := secureServer.Serve(secureConn); err != nil {
		log.Fatal("failed to serve: ", err)
	}

}

func main() {

	flag.Parse()

	db.Init()

	log.Info("Astro Calc - Trance !")
	log.Info("--------->>>> Starting Astro Calc Microservice  <<<----------")

	go func() {
		runGRPCServer()
	}()

	runHtppServer()

}
