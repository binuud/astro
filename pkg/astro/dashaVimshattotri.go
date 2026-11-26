package astro

import (
	"time"

	astroV1 "github.com/binuud/astro/gen/go/v1/astro"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DASHA_VIMSHOTTARI_YEARS        = 120
	DASHA_VIMSHOTTARI_YEAR_SECONDS = 31558149 // Number of seconds in a year of 365.256363 sidereal days
	DASHA_VIMSHOTTARI              = "VIMSHOTTARI"
)

// Notes
// 3786834240 // 365.2425 * 120 * numseconds in a day // 365.256363

func NewVimshottariDasha(startTime time.Time) *AstroDasha {

	// TODO validate basic names and duration
	// subDashaMeta := make([]*astroV1.SubDashaMeta, 0)

	subDashaMeta := []*astroV1.SubDashaMeta{
		&astroV1.SubDashaMeta{
			Name:     PLANET_KETU,
			Duration: 7,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_VENUS,
			Duration: 20,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_SUN,
			Duration: 6,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_MOON,
			Duration: 10,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_MARS,
			Duration: 7,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_RAHU,
			Duration: 18,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_JUPITER,
			Duration: 16,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_SATURN,
			Duration: 19,
		},
		&astroV1.SubDashaMeta{
			Name:     PLANET_MERCURY,
			Duration: 17,
		},
	}

	meta := &astroV1.DashaMeta{
		Name:          DASHA_VIMSHOTTARI,
		StartTime:     timestamppb.New(startTime),
		DurationYears: DASHA_VIMSHOTTARI_YEARS,
		YearInSeconds: DASHA_VIMSHOTTARI_YEAR_SECONDS,
		DasaDegrees:   STAR_DEGREES,
		SubDashaMetas: subDashaMeta,
		Type:          astroV1.DashaMeta_Vimshottari,
	}

	meta.StartDasa = 0

	dasha := NewDasha(meta)

	return dasha
}
