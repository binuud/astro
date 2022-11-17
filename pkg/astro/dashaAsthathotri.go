package astro

import (
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DASHA_ASHTOTTARI_YEARS        = 108
	DASHA_ASHTOTTARI_YEAR_SECONDS = 3408150816
	// 365.2425 * 108 * numseconds in a day
	DASHA_ASHTOTTARI = "ASHTOTTARI"
)

func NewAsthathotriDasha(startTime time.Time) *AstroDasha {

	// TODO validate basic names and duration
	// subDashaMeta := make([]*astroV1.SubDashaMeta, 0)

	subDashaMeta := []*astroV1.SubDashaMeta{
		&astroV1.SubDashaMeta{
			Name:     PLANET_SUN,
			Duration: 6,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_MOON,
			Duration: 15,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_MARS,
			Duration: 8,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_RAHU,
			Duration: 17,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_JUPITER,
			Duration: 10,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_SATURN,
			Duration: 19,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_MERCURY,
			Duration: 12,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_VENUS,
			Duration: 21,
		},
	}

	meta := &astroV1.DashaMeta{
		Name:          DASHA_ASHTOTTARI,
		StartTime:     timestamppb.New(startTime),
		DurationYears: DASHA_ASHTOTTARI_YEARS,
		YearInSeconds: DASHA_ASHTOTTARI_YEAR_SECONDS,
		DasaDegrees:   STAR_DEGREES,
		SubDashaMetas: subDashaMeta,
		Type:          astroV1.DashaMeta_Ashtottari,
	}

	dasha := NewDasha(meta)
	//	dasha.CalculateLevel(1)

	return dasha
}
