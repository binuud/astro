package astro

import (
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DASHA_ASHTATTOTRI_SECONDS = 3408150816 // 365.2425 * 108 * numseconds in a day
	DASHA_ASHTATTOTRI         = "ASHTATTOTRI"
)

func NewAsthathotriDasha(startTime time.Time) *AstroDasha {

	// TODO validate basic names and duration
	// subDashaMeta := make([]*astroV1.SubDashaMeta, 0)

	subDashaMeta := []*astroV1.SubDashaMeta{
		&astroV1.SubDashaMeta{
			Name:     "KETU",
			Duration: 7,
		},
		&astroV1.SubDashaMeta{
			Name:     "VENUS",
			Duration: 20,
		},
		&astroV1.SubDashaMeta{
			Name:     "SUN",
			Duration: 6,
		},
		&astroV1.SubDashaMeta{
			Name:     "MOON",
			Duration: 10,
		},
		&astroV1.SubDashaMeta{
			Name:     "MARS",
			Duration: 7,
		},
		&astroV1.SubDashaMeta{
			Name:     "RAHU",
			Duration: 18,
		},
		&astroV1.SubDashaMeta{
			Name:     "JUPITER",
			Duration: 16,
		},
		&astroV1.SubDashaMeta{
			Name:     "SATURN",
			Duration: 19,
		},
		&astroV1.SubDashaMeta{
			Name:     "MERCURY",
			Duration: 17,
		},
	}

	meta := &astroV1.DashaMeta{
		Name:          DASHA_ASHTATTOTRI,
		StartTime:     timestamppb.New(startTime),
		Duration:      DASHA_ASHTATTOTRI_SECONDS,
		SubDashaMetas: subDashaMeta,
		Type:          astroV1.DashaMeta_Ashtottatri,
	}

	dasha := NewDasha(meta)
	//	dasha.CalculateLevel(1)

	return dasha
}
