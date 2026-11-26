package astro

import (
	"testing"

	astroV1 "github.com/binuud/astro/gen/go/v1/astro"
)

func TestNewAstroChart(t *testing.T) {

	// chart, err := NewAstroChart(&astroV1.BirthData{
	// 	Tob: &astroV1.TobData{
	// 		D:    04,
	// 		M:    02,
	// 		Y:    1978,
	// 		Hour: 10,
	// 		Min:  45,
	// 		Sec:  0,
	// 		Zone: "Asia/Kolkata",
	// 	},
	// 	Pob: &astroV1.PobData{
	// 		City: "Chennai",
	// 		Lat:  13.08,
	// 		Long: 80.7,
	// 	},
	// })

	chart, err := NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    14,
			M:    07,
			Y:    2026,
			Hour: 12,
			Min:  15,
			Sec:  0,
			Zone: "Asia/Kolkata",
		},
		Pob: &astroV1.PobData{
			City: "Chennai",
			Lat:  27.08,
			Long: 80.7,
		},
	})
	if err != nil {
		t.Errorf("Cannot prepare chart err(%v)", err)
	}
	t.Logf("Got result (%v) ", chart)
	chart.PrintChartTable()

}
