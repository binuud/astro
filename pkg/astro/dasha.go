package astro

import (
	"fmt"
	"math"
	"strings"
	"time"

	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
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

// Calculate
// Calculates the dasha persiods for 3 nested recursions
// maha, antar, pratyantar
func (d *AstroDasha) Calculate() {

	startTime := d.Meta.StartTime.AsTime()
	currentDasha := d.Meta.StartDasa
	endTime := startTime.Add(time.Second * time.Duration(d.Meta.YearInSeconds*d.Meta.DurationYears))
	d.CalculateLevel(1, 2, startTime, endTime, int(currentDasha))
}

// AdjustTimePeriods
// Instead of having fixed dasa bhukti time periods, calculate same from longitude moved by planets in corresponding rasis
func (d *AstroDasha) AdjustTimePeriods(chart *astroV1.AstroObject) {

}

func (d *AstroDasha) AdjustTimePeriodsWithBaseOld(chart *astroV1.AstroObject, basePlanet string) {

	// calculate new midpoint for reference
	// baseDegrees := chart.Planets.PlanetsMap[basePlanet].Longitude
	degreeCorrection := 0.0 // math.Mod(float64(baseDegrees), STAR_DEGREES) - 6.5
	totalDuration := 0
	for _, dasha := range d.Meta.SubDashaMetas {

		planet := chart.Planets.PlanetsMap[dasha.Name]
		bavaDeg := float64(planet.Longitude) - degreeCorrection
		//	print "$bavaDeg <br/>";
		if bavaDeg > 360.0 {
			bavaDeg = bavaDeg - 360.0
		}

		if bavaDeg < 0 {
			bavaDeg = 360 + bavaDeg
		}

		remainingDegrees := float64(planet.Star)*float64(STAR_DEGREES) - bavaDeg

		dasha.Duration = int32(math.Round(remainingDegrees))
		totalDuration += int(dasha.Duration)

		fmt.Printf(" %10s %02d %02d %06.2f %06.2f %06.2f %03d\n", planet.Name, planet.Star, planet.Paada, planet.Longitude, bavaDeg, remainingDegrees, dasha.Duration)
	}

	fmt.Printf("Total Duration %d Age - %d", totalDuration, totalDuration+int(chart.BirthData.Tob.Y))
	// d.Meta.DurationYears = int64(totalDuration)

}

// finds the long traversed by the base planet
// subtracts same from all planet longiture, this resets the clock to base planet.
// image the sky as clock, with each planet as hands like hr,min, second
// we want the clocks start time to sync with birth
// so rotate the clock counter clockwise so that the base planet points to 0 long of the corresponding star
// during birth

func (d *AstroDasha) AdjustTimePeriodsWithBase(chart *astroV1.AstroObject, basePlanetName string) {

	// calculate new midpoint for reference
	basePlanet := chart.Planets.PlanetsMap[basePlanetName]
	degreeCorrection := basePlanet.Longitude - float64(basePlanet.Star-1)*STAR_DEGREES
	// degreeCorrection := math.Mod(float64(basePlanet.Longitude), PAADA_DEGREES) // rounded at paada

	fmt.Println("Base Planet degree correction %f", degreeCorrection)

	totalDuration := 0
	for _, dasha := range d.Meta.SubDashaMetas {

		planet := chart.Planets.PlanetsMap[dasha.Name]
		bavaDeg := float64(planet.Longitude) - degreeCorrection
		//	print "$bavaDeg <br/>";
		if bavaDeg > 360.0 {
			bavaDeg = bavaDeg - 360.0
		}

		if bavaDeg < 0 {
			bavaDeg = 360 + bavaDeg
		}

		remainingDegrees := float64(planet.Star)*float64(STAR_DEGREES) - bavaDeg

		dasha.Duration = int32(math.Round(remainingDegrees))
		totalDuration += int(dasha.Duration)

		fmt.Printf(" %10s %02d %02d %06.2f %06.2f %06.2f %03d\n", planet.Name, planet.Star, planet.Paada, planet.Longitude, bavaDeg, remainingDegrees, dasha.Duration)
	}
	fmt.Printf("Total Duration %d Age - %d\n\n\n", totalDuration, totalDuration+int(chart.BirthData.Tob.Y))

	if totalDuration > DASHA_VIMSHOTTARI_YEARS {
		// now ratio the num years so that the total duration matches 120 (upscale/downscale)
		scaledTotalDuration := 0
		for _, dasha := range d.Meta.SubDashaMetas {
			planet := chart.Planets.PlanetsMap[dasha.Name]
			oldDuration := dasha.Duration
			dasha.Duration = int32(math.Round(float64(dasha.Duration) / float64(totalDuration) * DASHA_VIMSHOTTARI_YEARS))
			scaledTotalDuration += int(dasha.Duration)

			fmt.Printf(" %10s %02d %02d  %03d\n", planet.Name, planet.Star, oldDuration, dasha.Duration)
		}

		fmt.Printf("Total Scaled Duration %d  --  %d Age - %d\n\n", totalDuration, scaledTotalDuration, scaledTotalDuration+int(chart.BirthData.Tob.Y))

	} else {
		d.Meta.DurationYears = int64(totalDuration)
	}

}

func (d *AstroDasha) AdjustStartTime(degreesMoved float64) {
	currentDasha := d.Meta.StartDasa
	adjustTime := float64(d.Meta.YearInSeconds) * float64(d.Meta.SubDashaMetas[currentDasha].Duration) * degreesMoved / d.Meta.DasaDegrees

	startTime := d.Meta.StartTime.AsTime().Add(-(time.Second * time.Duration(adjustTime)))
	// logrus.Infof("Correction time %v %v", startTime, d.Meta.StartTime.AsTime())
	d.Meta.StartTime = timestamppb.New(startTime)
}

func (d *AstroDasha) CalculateLevel(level int, totalLevels int, startTime time.Time, endTime time.Time, currentDasha int) []*astroV1.DashaItem {

	dashas := make([]*astroV1.DashaItem, 0)

	dashaCounter := 0

	numDashas := len(d.Meta.SubDashaMetas)

	timeDiff := endTime.Sub(startTime).Seconds()
	// logrus.Infof("Time difference %v", int(timeDiff))

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
	return int64(float64(totalSeconds) * float64(ratio) / float64(d.Meta.DurationYears))
}

func (d *AstroDasha) Print() {
	d.PrintDashas(1, 2, d.Dashas)
}

func (d *AstroDasha) PrintDashas(level int, totalLevels int, dashas []*astroV1.DashaItem) {
	zone, _ := time.LoadLocation("Asia/Kolkata")

	intend := strings.Repeat("-----", level)
	for _, dasha := range dashas {
		startAge := dasha.StartTime.AsTime().In(zone).Year() - d.Meta.StartTime.AsTime().In(zone).Year()
		endAge := dasha.EndTime.AsTime().In(zone).Year() - d.Meta.StartTime.AsTime().In(zone).Year()
		fmt.Printf(" %s %10s %v %v %3d to %3d years \n", intend, dasha.Name, dasha.StartTime.AsTime().In(zone), dasha.EndTime.AsTime().In(zone), startAge, endAge)
		if len(dasha.Dashas) > 0 && level <= totalLevels {
			d.PrintDashas(level+1, totalLevels, dasha.Dashas)
		}
	}
}

func (d *AstroDasha) PrintAntarDashas(intend string, dashas []*astroV1.DashaItem) {
	zone, _ := time.LoadLocation("Asia/Kolkata")

	for _, dasha := range dashas {
		fmt.Printf(" %s %10s %v %v \n", intend, dasha.Name, dasha.StartTime.AsTime().In(zone), dasha.EndTime.AsTime().In(zone))
		// if len(dasha.Dashas) > 0 {
		// 	d.PrintDashas(intend+"-----", dasha.Dashas)
		// }
	}
}
