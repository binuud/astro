package astro

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"github.com/sirupsen/logrus"
)

type SunMoonEphimeris struct {
	Moon float64
	Sun  float64
}

type PanchangTransits struct {
	Name    int
	Transit time.Time
}

type Panchang struct {
	Year      int64               // year for which the panchang has to be calculated
	Ephimeris []*SunMoonEphimeris // moon ephimeris for each day of the year (only longitudes)
	StartTime time.Time           // start date when panchang calculation is started
	Stars     []*PanchangTransits
	Yogas     []*PanchangTransits
	Karanas   []*PanchangTransits
	Tithis    []*PanchangTransits
}

// getMoonEphimeris
// arguments: year - the year for which panchang is calculated
// returns a list of longitudes for moon for every 12hrs
// from 1st of the year to the last date of the given year
func getSunMoonEphimeris(year int64) ([]*SunMoonEphimeris, error) {
	birthData := &astroV1.BirthData{
		Tob: &astroV1.TobData{
			Y:    year - 1,
			D:    31,
			M:    12,
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
	moonEphimeris, err := ephimeris.GetSunMoonEphimerisForYear()
	if err != nil {
		return nil, err
	}
	// logrus.Infof("Moon data %v", moonEphimeris)
	moonDegreesString := strings.Split(moonEphimeris, "\n")
	sunMoonEphimeris := make([]*SunMoonEphimeris, 0)
	for _, item := range moonDegreesString {
		// logrus.Infof("Parsing %s", item)
		fields := strings.Fields(item)
		if len(fields) != 2 {
			// logrus.Errorf("Cannot parse longitudes item ::%s:: err(%v)", item, fields)
			continue
		}

		sunLongitude, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			logrus.Errorf("Cannot parse longitudes err(%v)", err)
			return nil, err
		}

		moonLongitude, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			logrus.Errorf("Cannot parse longitudes err(%v)", err)
			return nil, err
		}

		sunMoonLongitudes := &SunMoonEphimeris{
			Moon: moonLongitude,
			Sun:  sunLongitude,
		}

		sunMoonEphimeris = append(sunMoonEphimeris, sunMoonLongitudes)
	}
	return sunMoonEphimeris, nil
}

func NewPanchangMoon(year int64) (*Panchang, error) {
	var err error
	if year < 0 || year > 2050 {
		err := fmt.Errorf("year has to be in the range of 0 to 2050")
		return nil, err
	}
	myPanchang := &Panchang{
		Year:      year,
		StartTime: time.Date(int(year-1), time.Month(12), int(31), int(0), int(0), int(0), 0, time.UTC),
	}

	myPanchang.Ephimeris, err = getSunMoonEphimeris(year)
	if err != nil {
		return nil, err
	}
	// logrus.Infof("Moon degrees %v", myPanchang.Ephimeris)

	myPanchang.getMoonPanchang()
	myPanchang.getTithiPanchang()
	myPanchang.getKaranaPanchang()
	myPanchang.getYogaPanchang()
	return myPanchang, nil
}

func (panchang *Panchang) getMoonPanchang() {

	zone, _ := time.LoadLocation("Asia/Kolkata")

	stars := make([]*PanchangTransits, 0)
	counter := 0
	numItems := len(panchang.Ephimeris) - 1
	currentStar := GetStarNumber(panchang.Ephimeris[0].Moon)

	for counter < numItems {

		t := panchang.StartTime

		nextStar := currentStar + 1
		if nextStar > TOTAL_STARS {
			nextStar = 1
		}

		nextStarDegrees := STAR_DEGREES * float64(nextStar-1)
		//logrus.Infof("Checking star current (%d), next (%d), transit degrees (%f)", currentStar, nextStar, nextStarDegrees)

		transitMinutes, err := GetTransitTime(panchang.Ephimeris[counter].Moon, panchang.Ephimeris[counter+1].Moon, QUARTER_DAY_MINUTES, nextStarDegrees)
		if err == nil {
			currentStar++
			if currentStar > TOTAL_STARS {
				currentStar = 1
			}

			t = t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			starTransit := &PanchangTransits{
				Name:    currentStar,
				Transit: t,
			}

			stars = append(stars, starTransit)

			logrus.Infof("Star %s, %f<->%f,  %v", STAR_NAMES[currentStar], panchang.Ephimeris[counter].Moon, panchang.Ephimeris[counter+1].Moon, t.In(zone))
		}

		counter++

	}

	panchang.Stars = stars

}

func (panchang *Panchang) getTithiPanchang() {
	zone, _ := time.LoadLocation("Asia/Kolkata")

	tithis := make([]*PanchangTransits, 0)

	counter := 0
	numItems := len(panchang.Ephimeris) - 1
	currentTithi := GetTithiNumber(panchang.Ephimeris[0].Sun, panchang.Ephimeris[0].Moon)

	for counter < numItems {

		t := panchang.StartTime

		nextTithi := currentTithi + 1
		if nextTithi > TOTAL_TITHIS {
			nextTithi = 1
		}

		nextTithiDegrees := TITHI_DEGREES * float64(nextTithi-1)
		//logrus.Infof("Checking star current (%d), next (%d), transit degrees (%f)", currentStar, nextStar, nextStarDegrees)

		startTithiDegree := GetTithiDegree(panchang.Ephimeris[counter].Sun, panchang.Ephimeris[counter].Moon)
		endTithiDegree := GetTithiDegree(panchang.Ephimeris[counter+1].Sun, panchang.Ephimeris[counter+1].Moon)

		transitMinutes, err := GetTransitTime(startTithiDegree, endTithiDegree, QUARTER_DAY_MINUTES, nextTithiDegrees)
		if err == nil {
			currentTithi++
			if currentTithi > TOTAL_TITHIS {
				currentTithi = 1
			}
			paksha := GetPaksha(currentTithi)
			t = t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			tithiTransit := &PanchangTransits{
				Name:    currentTithi,
				Transit: t,
			}

			tithis = append(tithis, tithiTransit)

			logrus.Infof("%s(%s), %f<->%f @ %v", PAKSHA_NAMES[paksha], TITHI_NAMES[currentTithi], startTithiDegree, endTithiDegree, t.In(zone))
		}

		counter++

	}

	panchang.Tithis = tithis

}

func (panchang *Panchang) getKaranaPanchang() {
	zone, _ := time.LoadLocation("Asia/Kolkata")

	karanas := make([]*PanchangTransits, 0)

	counter := 0
	numItems := len(panchang.Ephimeris) - 1
	currentKarana := GetKaranaNumber(panchang.Ephimeris[0].Sun, panchang.Ephimeris[0].Moon)

	for counter < numItems {

		t := panchang.StartTime

		nextKarana := currentKarana + 1
		if nextKarana > TOTAL_KARANAS {
			nextKarana = 1
		}

		nextKaranaDegrees := KARANA_DEGREES * float64(nextKarana-1)
		//logrus.Infof("Checking star current (%d), next (%d), transit degrees (%f)", currentStar, nextStar, nextStarDegrees)

		startKaranaDegree := GetTithiDegree(panchang.Ephimeris[counter].Sun, panchang.Ephimeris[counter].Moon)
		endKaranaDegree := GetTithiDegree(panchang.Ephimeris[counter+1].Sun, panchang.Ephimeris[counter+1].Moon)

		transitMinutes, err := GetTransitTime(startKaranaDegree, endKaranaDegree, QUARTER_DAY_MINUTES, nextKaranaDegrees)
		if err == nil {
			currentKarana++
			if currentKarana > TOTAL_KARANAS {
				currentKarana = 1
			}
			t = t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			karanas = append(karanas, &PanchangTransits{
				Name:    currentKarana,
				Transit: t,
			})

			logrus.Infof("%s, %f<->%f @ %v", KARANA_NAMES[currentKarana], startKaranaDegree, endKaranaDegree, t.In(zone))
		}

		counter++

	}

	panchang.Karanas = karanas

}

func (panchang *Panchang) getYogaPanchang() {

	zone, _ := time.LoadLocation("Asia/Kolkata")

	yogas := make([]*PanchangTransits, 0)

	counter := 0
	numItems := len(panchang.Ephimeris) - 1
	currentYoga := GetYogaNumber(panchang.Ephimeris[0].Sun, panchang.Ephimeris[0].Moon)
	logrus.Infof("Starting Yoga (%d)", currentYoga)
	for counter < numItems {

		t := panchang.StartTime

		nextYoga := currentYoga + 1
		if nextYoga > TOTAL_YOGAS {
			nextYoga = 1
		}

		nextYogaDegrees := YOGA_DEGREES * float64(nextYoga-1)

		startYogaDegree := GetYogaDegree(panchang.Ephimeris[counter].Sun, panchang.Ephimeris[counter].Moon)
		endYogaDegree := GetYogaDegree(panchang.Ephimeris[counter+1].Sun, panchang.Ephimeris[counter+1].Moon)

		transitMinutes, err := GetTransitTime(startYogaDegree, endYogaDegree, QUARTER_DAY_MINUTES, nextYogaDegrees)
		if err == nil {
			currentYoga++
			if currentYoga > TOTAL_YOGAS {
				currentYoga = 1
			}

			t = t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			yogas = append(yogas, &PanchangTransits{
				Name:    currentYoga,
				Transit: t,
			})

			logrus.Infof(" %s, %f<->%f @ %v", YOGA_NAMES[currentYoga], startYogaDegree, endYogaDegree, t.In(zone))
		}

		counter++

	}

	panchang.Yogas = yogas

}
