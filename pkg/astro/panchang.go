package astro

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Panchang struct {
	*astroV1.Panchang
}

func NewPanchangMoon(year int32, month int32) (*Panchang, error) {

	if year < 0 || year > 2050 {
		err := fmt.Errorf("year has to be in the range of 0 to 2050")
		return nil, err
	}

	if month < 1 || month > 12 {
		err := fmt.Errorf("month has to be in the range of 1 to 12")
		return nil, err
	}

	startTime := time.Date(int(year), time.Month(month), int(1), int(0), int(0), int(0), 0, time.UTC)
	startTime.AddDate(0, 0, -2) // reducing 1 day to get the overlapping star

	myPanchang := &astroV1.Panchang{
		Year:      year,
		Month:     month,
		StartTime: timestamppb.New(startTime),
	}

	// logrus.Infof("Moon degrees %v", myPanchang.Ephimeris)

	return &Panchang{myPanchang}, nil
}

// getMoonEphimeris
// arguments: year - the year for which panchang is calculated
// returns a list of longitudes for moon for every 12hrs
// from 1st of the year to the last date of the given year
func (panchang *Panchang) getSunMoonEphimeris(year int32, month int32) ([]*astroV1.SunMoonEphimeris, error) {
	birthData := &astroV1.BirthData{
		Tob: &astroV1.TobData{
			Y:    year,
			D:    1,
			M:    month,
			Hour: 0,
			Min:  0,
			Sec:  0,
		},
	}
	birthData.Tob.Tobtime = timestamppb.New(panchang.StartTime.AsTime())
	// birthData, err := MakeBirthData(birthData)
	// if err != nil {
	// 	logrus.Errorf("Cannot parse time data for astro chart err(%v)", err)
	// }

	ephimeris := NewEphimeris((birthData))
	moonEphimeris, err := ephimeris.GetSunMoonEphimerisForMonth()
	if err != nil {
		return nil, err
	}
	// logrus.Infof("Moon data %v", moonEphimeris)
	moonDegreesString := strings.Split(moonEphimeris, "\n")
	sunMoonEphimeris := make([]*astroV1.SunMoonEphimeris, 0)
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

		sunMoonLongitudes := &astroV1.SunMoonEphimeris{
			Moon: moonLongitude,
			Sun:  sunLongitude,
		}

		sunMoonEphimeris = append(sunMoonEphimeris, sunMoonLongitudes)
	}
	return sunMoonEphimeris, nil
}

func (panchang *Panchang) CalculatePanchang() error {

	var err error
	panchang.Ephimeris, err = panchang.getSunMoonEphimeris(panchang.Year, panchang.Month)
	if err != nil {
		return err
	}
	panchang.getMoonPanchang()
	panchang.getTithiPanchang()
	panchang.getKaranaPanchang()
	panchang.getYogaPanchang()
	return nil
}

func (panchang *Panchang) getMoonPanchang() {

	// zone, _ := time.LoadLocation("Asia/Kolkata")

	stars := make([]*astroV1.PanchangTransits, 0)
	counter := 0
	t := panchang.StartTime.AsTime()
	numItems := len(panchang.Ephimeris) - 1
	currentStar := GetStarNumber(panchang.Ephimeris[0].Moon)

	for counter < numItems {

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

			t1 := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			starTransit := &astroV1.PanchangTransits{
				Name:    int32(currentStar),
				Transit: timestamppb.New(t1),
			}

			stars = append(stars, starTransit)

			// logrus.Infof("Star %s, %f<->%f,  %v", STAR_NAMES[currentStar], panchang.Ephimeris[counter].Moon, panchang.Ephimeris[counter+1].Moon, t.In(zone))
		}

		counter++

	}

	panchang.Stars = stars

}

func (panchang *Panchang) getTithiPanchang() {
	// zone, _ := time.LoadLocation("Asia/Kolkata")

	tithis := make([]*astroV1.PanchangTransits, 0)

	counter := 0
	numItems := len(panchang.Ephimeris) - 1
	currentTithi := GetTithiNumber(panchang.Ephimeris[0].Sun, panchang.Ephimeris[0].Moon)

	for counter < numItems {

		t := panchang.StartTime.AsTime()

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
			// paksha := GetPaksha(currentTithi)
			t := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			tithiTransit := &astroV1.PanchangTransits{
				Name:    int32(currentTithi),
				Transit: timestamppb.New(t),
			}

			tithis = append(tithis, tithiTransit)

			// logrus.Infof("%s(%s), %f<->%f @ %v", PAKSHA_NAMES[paksha], TITHI_NAMES[currentTithi], startTithiDegree, endTithiDegree, t.In(zone))
		}

		counter++

	}

	panchang.Tithis = tithis

}

func (panchang *Panchang) getKaranaPanchang() {
	// zone, _ := time.LoadLocation("Asia/Kolkata")

	karanas := make([]*astroV1.PanchangTransits, 0)

	counter := 0
	numItems := len(panchang.Ephimeris) - 1
	currentKarana := GetKaranaNumber(panchang.Ephimeris[0].Sun, panchang.Ephimeris[0].Moon)

	for counter < numItems {
		t := panchang.StartTime.AsTime()
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
			t := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			karanas = append(karanas, &astroV1.PanchangTransits{
				Name:    int32(currentKarana),
				Transit: timestamppb.New(t),
			})

			// logrus.Infof("%s, %f<->%f @ %v", KARANA_NAMES[currentKarana], startKaranaDegree, endKaranaDegree, t.In(zone))
		}

		counter++

	}

	panchang.Karanas = karanas

}

func (panchang *Panchang) getYogaPanchang() {

	// zone, _ := time.LoadLocation("Asia/Kolkata")

	yogas := make([]*astroV1.PanchangTransits, 0)

	counter := 0
	numItems := len(panchang.Ephimeris) - 1
	currentYoga := GetYogaNumber(panchang.Ephimeris[0].Sun, panchang.Ephimeris[0].Moon)
	logrus.Infof("Starting Yoga (%d)", currentYoga)
	for counter < numItems {
		t := panchang.StartTime.AsTime()
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

			t := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*QUARTER_DAY_MINUTES))

			yogas = append(yogas, &astroV1.PanchangTransits{
				Name:    int32(currentYoga),
				Transit: timestamppb.New(t),
			})

			// logrus.Infof(" %s, %f<->%f @ %v", YOGA_NAMES[currentYoga], startYogaDegree, endYogaDegree, t.In(zone))
		}

		counter++

	}

	panchang.Yogas = yogas

}
