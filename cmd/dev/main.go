package main

import (
	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/dronasys-com/astro/pkg/astro"
	log "github.com/sirupsen/logrus"
)

func main() {

	test_dasa()
	// test_newdasa()
	// test_calendar()
	// test_chart()
	// test_calendar()
}

func test_calendar() {
	p, err := astro.NewPanchang(2022, 2, 1)
	if err != nil {
		log.Errorf("Cannot prepare chart err(%v)", err)
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
		log.Error("Cannot prepare chart err(%v)", err)
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
		log.Error("Cannot prepare chart err(%v)", err)
	}

	dasa := astro.NewVimshottariDasha(chart.BirthData.Tob.Tobtime.AsTime())
	curStar := chart.Planets.PlanetsMap[astro.PLANET_MOON].Star - 1
	dasa.AdjustStartTime(chart.Planets.PlanetsMap[astro.PLANET_MOON].Longitude - float64(curStar)*float64(astro.STAR_DEGREES))
	dasa.Calculate()
	dasa.Print()
}

func test_newdasa() {
	// log.Infof("Config is (%v)", config.Config)
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
		log.Error("Cannot prepare chart err(%v)", err)
	}

	dasa := astro.NewVimshottariDasha(chart.BirthData.Tob.Tobtime.AsTime())
	// curStar := chart.Planets.PlanetsMap[astro.PLANET_MOON].Star - 1
	//dasa.AdjustTimePeriodsWithBase(chart.AstroObject, astro.PLANET_MOON)
	dasa.AdjustTimePeriodsWithBase(chart.AstroObject, astro.PLANET_ASC)
	// dasa.AdjustStartTime(chart.Planets.PlanetsMap[astro.PLANET_MOON].Longitude - float64(curStar)*float64(astro.STAR_DEGREES))
	dasa.Calculate()
	dasa.Print()
}
