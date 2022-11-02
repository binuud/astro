package main

import (
	"flag"

	"github.com/dronasys-com/astro/pkg/astro"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"github.com/sirupsen/logrus"
)

var (
	grpcPort = flag.Int("grpcPort", 50061, "The Secure Server port, with Astro Calculations")
)

func main() {

	// bootstrap.Init()
	// db.Init()

	// logrus.Info("Astro Calc - Trance !")
	// logrus.Info("--------->>>> Starting Astro Calc Microservice  <<<----------")

	// flag.Parse()
	// secureConn, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// secureServer := grpc.NewServer(
	// 	grpc.UnaryInterceptor(authServer.UnaryInterceptor))

	// astroServer.NewGrpcAstroServer(secureServer)

	// logrus.Infof("Astro Calculations Server *) (* listening on (%v)", secureConn.Addr())
	// if err := secureServer.Serve(secureConn); err != nil {
	// 	logrus.Fatal("failed to serve: ", err)
	// }

	test_chart()
	// test_calendar()
}

func test_calendar() {
	p, err := astro.NewPanchangMoon(2022)
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	cal := astro.NewCalendarPrinter(p)
	cal.PrintCalendar()
}

func test_chart() {
	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    04,
			M:    02,
			Y:    1978,
			Hour: 10,
			Min:  44,
			Sec:  0,
			Zone: "Asia/Kolkata",
		},
		Pob: &astroV1.PobData{
			City: "Chennai",
			Lat:  13.08,
			Long: 80.7,
		},
	})
	if err != nil {
		logrus.Error("Cannot prepare chart err(%v)", err)
	}
	chart.PrintChartTable()
}
