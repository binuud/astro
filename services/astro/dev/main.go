package main

import (
	"github.com/dronasys-com/astro/pkg/astro"
	"github.com/dronasys-com/astro/pkg/foundation/bootstrap"
	"github.com/dronasys-com/astro/pkg/foundation/config"
	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
)

func main() {

	bootstrap.Init()

	// test_dasa()
	test_newdasa()
	// test_calendar()
	// test_chart()
	// test_calendar()
}

func test_calendar() {
	p, err := astro.NewPanchang(2022, 2, 1)
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	p.PrintMonth()
	// p.CalculatePanchang()
	// cal := astro.NewCalendarPrinter(p.Panchang)
	// cal.PrintCalendar()
}

func test_chart() {
	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    04,
			M:    02,
			Y:    1978,
			Hour: 10,
			Min:  45,
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
			Min:  45,
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

	dasa := astro.NewVimshottariDasha(chart.BirthData.Tob.Tobtime.AsTime())
	curStar := chart.Planets.PlanetsMap[astro.PLANET_MOON].Star - 1
	dasa.AdjustStartTime(chart.Planets.PlanetsMap[astro.PLANET_MOON].Longitude - float64(curStar)*float64(astro.STAR_DEGREES))
	dasa.Calculate()
	dasa.Print()
}

func test_newdasa() {
	logrus.Infof("Config is (%v)", config.Config)
	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    04,
			M:    02,
			Y:    1978,
			Hour: 10,
			Min:  45,
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

	dasa := astro.NewVimshottariDasha(chart.BirthData.Tob.Tobtime.AsTime())
	// curStar := chart.Planets.PlanetsMap[astro.PLANET_MOON].Star - 1
	//dasa.AdjustTimePeriodsWithBase(chart.AstroObject, astro.PLANET_MOON)
	dasa.AdjustTimePeriodsWithBase(chart.AstroObject, astro.PLANET_ASC)
	// dasa.AdjustStartTime(chart.Planets.PlanetsMap[astro.PLANET_MOON].Longitude - float64(curStar)*float64(astro.STAR_DEGREES))
	dasa.Calculate()
	dasa.Print()
}
