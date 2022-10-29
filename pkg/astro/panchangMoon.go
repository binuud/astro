package astro

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"github.com/sirupsen/logrus"
)

type PanchangMoon struct {
	Year          int64     // year for which the panchang has to be calculated
	StarDegrees   []float64 // helper data - list of degree boundaries for 27 stars
	MoonEphimeris []float64 // moon ephimeris for each day of the year (only longitudes)
}

func initPanchangMoon(year int64) ([]float64, error) {

	moonDegrees, err := getMoonEphimeris(year)
	if err != nil {
		return nil, err
	}
	return moonDegrees, nil

}

// getMoonEphimeris
// arguments: year - the year for which panchang is calculated
// returns a list of longitudes for moon for every 12hrs
// from 1st of the year to the last date of the given year
func getMoonEphimeris(year int64) ([]float64, error) {
	birthData := &astroV1.BirthData{
		Tob: &astroV1.TobData{
			Y:    year,
			D:    1,
			M:    1,
			Hour: 0,
			Min:  0,
			Sec:  0,
		},
	}
	birthData, err := MakeBirthData(birthData)
	if err != nil {
		logrus.Errorf("Cannot parse time data for astro chart err(%v)", err)
	}

	ephimeris := NewEphimeris((birthData))
	moonEphimeris, err := ephimeris.GetMoonEphimerisForYear()
	if err != nil {
		return nil, err
	}
	// logrus.Infof("Moon data %v", moonEphimeris)
	moonDegreesString := strings.Split(moonEphimeris, "\n")
	moonDegrees := make([]float64, 0)
	for _, item := range moonDegreesString {
		// logrus.Infof("Parsing %s", item)
		fields := strings.Fields(item)
		if len(fields) == 0 {
			logrus.Errorf("Cannot parse longitudes item ::%s:: err(%v)", item, fields)
			continue
		}
		longitude, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			logrus.Errorf("Cannot parse longitudes err(%v)", err)
			return nil, err
		}
		moonDegrees = append(moonDegrees, longitude)
	}
	return moonDegrees, nil
}

func NewPanchangMoon(year int64) (*PanchangMoon, error) {
	if year < 0 || year > 2050 {
		err := fmt.Errorf("year has to be in the range of 0 to 2050")
		return nil, err
	}
	myPanchang := &PanchangMoon{
		Year:        year,
		StarDegrees: GetStarDegreesArray(),
	}

	moonDegrees, err := getMoonEphimeris(year)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Moon degrees %v", moonDegrees)
	myPanchang.getMoonPanchang(moonDegrees)

	return myPanchang, nil
}

func (moon *PanchangMoon) getMoonPanchang(moonDegrees []float64) {
	zone, _ := time.LoadLocation("Asia/Kolkata")
	// t := time.Date(int(moon.Year), time.Month(1), int(1), int(1), int(1), int(1), 0, nil)
	halfDayCounter := 0
	dayCounter := 0
	moonDegreesNumItems := len(moonDegrees) - 1
	currentStar := GetStarNumber(moonDegrees[0])

	for halfDayCounter < moonDegreesNumItems {

		t := time.Date(int(moon.Year), time.Month(1), int(1), int(0), int(0), int(0), 0, time.UTC)

		nextStar := currentStar + 1
		if nextStar > TOTAL_STARS {
			nextStar = 1
		}

		nextStarDegrees := STAR_DEGREES * float64(nextStar-1)
		//logrus.Infof("Checking star current (%d), next (%d), transit degrees (%f)", currentStar, nextStar, nextStarDegrees)

		transitMinutes, err := GetTransitTime(moonDegrees[halfDayCounter], moonDegrees[halfDayCounter+1], 720, nextStarDegrees)
		if err == nil {
			currentStar++
			if currentStar > TOTAL_STARS {
				currentStar = 1
			}
			t = t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(halfDayCounter*720))
			logrus.Infof("Checking transit %d, %f, %f, %d, %v", nextStar, moonDegrees[halfDayCounter], moonDegrees[halfDayCounter+1], transitMinutes, t.In(zone))
		}

		halfDayCounter++
		if halfDayCounter%2 == 0 {
			dayCounter++
		}

	}

}
