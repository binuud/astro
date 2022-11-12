package astro

import (
	"fmt"
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AstroDasha struct {
	astroV1.DashaSystem
}

func NewDasha(meta *astroV1.DashaMeta) *AstroDasha {
	// TODO validate basic names and duration
	return &AstroDasha{
		astroV1.DashaSystem{
			Meta: meta,
		},
	}
}

func (d *AstroDasha) Calculate() {

	startTime := d.Meta.StartTime.AsTime()
	currentDasha := d.Meta.StartDasa
	endTime := startTime.Add(time.Second * time.Duration(DASHA_VIMSHATTOTRI_SECONDS))
	d.CalculateLevel(1, 3, startTime, endTime, int(currentDasha))
}

func (d *AstroDasha) AdjustStartTime(degreesMoved float64) {
	currentDasha := d.Meta.StartDasa
	adjustTime := SECONDS_IN_YEAR * float64(d.Meta.SubDashaMetas[currentDasha].Duration) * degreesMoved / d.Meta.DasaDegrees

	startTime := d.Meta.StartTime.AsTime().Add(-(time.Second * time.Duration(adjustTime)))
	logrus.Infof("Correction time %v %v", startTime, d.Meta.StartTime.AsTime())
	d.Meta.StartTime = timestamppb.New(startTime)
}

func (d *AstroDasha) CalculateLevel(level int, totalLevels int, startTime time.Time, endTime time.Time, currentDasha int) []*astroV1.DashaItem {

	dashas := make([]*astroV1.DashaItem, 0)

	dashaCounter := 0

	numDashas := len(d.Meta.SubDashaMetas)

	timeDiff := endTime.Sub(startTime).Seconds()
	logrus.Infof("Time difference %v", int(timeDiff))

	for dashaCounter < numDashas {
		if currentDasha >= numDashas {
			currentDasha = 0
		}
		durationInSeconds := d.getDurationRatioInSeconds(int(timeDiff), int(d.Meta.SubDashaMetas[currentDasha].Duration))
		endTime := startTime.Add(time.Second * time.Duration(durationInSeconds))
		newDasha := &astroV1.DashaItem{
			Name:      d.Meta.SubDashaMetas[currentDasha].Name,
			Level:     int32(level),
			StartTime: timestamppb.New(startTime),
			EndTime:   timestamppb.New(endTime),
		}
		if level < totalLevels {
			newDasha.Dashas = d.CalculateLevel(level+1, totalLevels, startTime, endTime, currentDasha)
		}
		dashas = append(dashas, newDasha)

		startTime = newDasha.EndTime.AsTime()
		dashaCounter++
		currentDasha++
	}

	if level == 1 {
		d.Dashas = dashas
		return nil
	} else {
		return dashas
	}

}

func (d *AstroDasha) getDurationRatioInSeconds(totalSeconds int, ratio int) int64 {
	return int64(float64(totalSeconds) * float64(ratio) / float64(d.Meta.Duration))
}

func (d *AstroDasha) Print() {
	d.PrintDashas("", d.Dashas)
}

func (d *AstroDasha) PrintDashas(intend string, dashas []*astroV1.DashaItem) {
	zone, _ := time.LoadLocation("Asia/Kolkata")

	for _, dasha := range dashas {
		fmt.Printf(" %s %10s %v %v \n", intend, dasha.Name, dasha.StartTime.AsTime().In(zone), dasha.EndTime.AsTime().In(zone))
		if len(dasha.Dashas) > 0 {
			d.PrintDashas(intend+"-----", dasha.Dashas)
		}
	}
}
