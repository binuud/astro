package main

import (
	"github.com/dronasys-com/astro/pkg/astro"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/dronasys-com/trance-microservices/pkg/foundation/bootstrap"
	"github.com/sirupsen/logrus"
)

func main() {

	bootstrap.Init()

	test_dasa()
	// test_calendar()
	// test_chart()
	// test_calendar()
}

func test_calendar() {
	p, err := astro.NewPanchangMoon(2022, 2)
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	p.CalculatePanchang()
	cal := astro.NewCalendarPrinter(p.Panchang)
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

func test_dasa() {
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

	dasa := astro.NewVimshattotriDasha(chart.BirthData.Tob.Tobtime.AsTime())
	curStar := chart.Planets.PlanetsMap[astro.PLANET_MOON].Star - 1
	dasa.AdjustStartTime(chart.Planets.PlanetsMap[astro.PLANET_MOON].Longitude - float64(curStar)*float64(astro.STAR_DEGREES))
	dasa.Calculate()
	dasa.Print()
}
