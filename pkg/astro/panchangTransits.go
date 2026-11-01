package astro

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/dronasys-com/astro/pkg/astro/ephimerisClient"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transits struct {
	*astroV1.Transits
}

// NewTransitsMoon
// create a transit array holder, with start time set to 2 days prior in a month
func NewTransits(year int32, month int32) (*Transits, error) {

	if year < 0 || year > 2050 {
		err := fmt.Errorf("year has to be in the range of 0 to 2050")
		return nil, err
	}

	if month < 1 || month > 12 {
		err := fmt.Errorf("month has to be in the range of 1 to 12")
		return nil, err
	}

	startTime := time.Date(int(year), time.Month(month), int(1), int(0), int(0), int(0), 0, time.UTC)
	startTime = startTime.AddDate(0, 0, -4) // reducing 2 day to get the overlapping star and rasi

	myTransits := &astroV1.Transits{
		Year:      year,
		Month:     month,
		StartTime: timestamppb.New(startTime),
	}

	return &Transits{myTransits}, nil
}

// getSunMoonEphimeris
// arguments: year - the year for which panchang is calculated
// returns a list of longitudes for moon for every 12hrs
// from 1st of the year to the last date of the given year
func (transits *Transits) getSunMoonEphimeris() ([]*astroV1.SunMoonEphimeris, error) {

	// Approach A: Extract all three components at once
	year, month, date := transits.StartTime.AsTime().Date()

	birthData := &astroV1.BirthData{
		Tob: &astroV1.TobData{
			Y:    int32(year),
			D:    int32(date),
			M:    int32(month),
			Hour: 0,
			Min:  0,
			Sec:  0,
		},
	}
	birthData.Tob.Tobtime = timestamppb.New(transits.StartTime.AsTime())
	// birthData, err := MakeBirthData(birthData)
	// if err != nil {
	// 	logrus.Errorf("Cannot parse time data for astro chart err(%v)", err)
	// }

	ephimeris := ephimerisClient.NewEphimerisClient()
	moonEphimeris, err := ephimeris.GetPanchangMonth(birthData)
	if err != nil {
		return nil, err
	}
	// logrus.Fatalf("Moon data %v", moonEphimeris)

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

func (transits *Transits) CalculateTransits() error {

	var err error
	transits.Ephimeris, err = transits.getSunMoonEphimeris()
	if err != nil {
		logrus.Panic(err)
	}

	transits.getMoonPanchang()
	transits.getRasiPanchang()
	transits.getPaadaPanchang()
	transits.getTithiPanchang()
	transits.getKaranaPanchang()
	transits.getYogaPanchang()
	return nil
}

func (transits *Transits) getMoonPanchang() {

	// zone, _ := time.LoadLocation("Asia/Kolkata")

	stars := make([]*astroV1.Transit, 0)
	counter := 0
	t := transits.StartTime.AsTime()
	numItems := len(transits.Ephimeris) - 1
	currentStar := GetStarNumber(transits.Ephimeris[0].Moon)

	for counter < numItems {

		nextStar := currentStar + 1
		if nextStar > TOTAL_STARS {
			nextStar = 1
		}

		nextStarDegrees := STAR_DEGREES * float64(nextStar-1)
		//logrus.Infof("Checking star current (%d), next (%d), transit degrees (%f)", currentStar, nextStar, nextStarDegrees)

		transitMinutes, err := GetTransitTime(transits.Ephimeris[counter].Moon, transits.Ephimeris[counter+1].Moon, HOUR_MINUTES, nextStarDegrees)
		if err == nil {
			currentStar++
			if currentStar > TOTAL_STARS {
				currentStar = 1
			}

			t1 := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*HOUR_MINUTES))

			starTransit := &astroV1.Transit{
				Name:    int32(currentStar),
				Transit: timestamppb.New(t1),
			}

			stars = append(stars, starTransit)

			// logrus.Infof("Star %s, %f<->%f,  %v", STAR_NAMES[currentStar], transits.Ephimeris[counter].Moon, transits.Ephimeris[counter+1].Moon, t.In(zone))
		}

		counter++

	}

	transits.Stars = stars

}

func (transits *Transits) getPaadaPanchang() {

	// zone, _ := time.LoadLocation("Asia/Kolkata")

	paadas := make([]*astroV1.Transit, 0)
	counter := 0
	t := transits.StartTime.AsTime()
	numItems := len(transits.Ephimeris) - 1
	currentItem := GetPaadaNumber(transits.Ephimeris[0].Moon)
	// logrus.Infof("Checking Paadas %d total ", numItems)

	for counter < numItems {

		nextItem := currentItem + 1
		if nextItem > TOTAL_PAADAS {
			nextItem = 1
		}

		nextItemDegrees := PAADA_DEGREES * float64(nextItem-1)
		// logrus.Infof("Checking paada counter (%d ), Item (%d - %d), transit degrees (%f)", counter, currentItem, nextItem, nextItemDegrees)

		transitMinutes, err := GetTransitTime(transits.Ephimeris[counter].Moon, transits.Ephimeris[counter+1].Moon, HOUR_MINUTES, nextItemDegrees)
		if err == nil {
			currentItem++
			if currentItem > TOTAL_PAADAS {
				currentItem = 1
			}

			t1 := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*HOUR_MINUTES))

			currentTransit := &astroV1.Transit{
				Name:    int32(currentItem),
				Transit: timestamppb.New(t1),
			}

			paadas = append(paadas, currentTransit)

			// logrus.Infof("Paada counter (%d) %s, %f<->%f,  %v", counter, currentItem, transits.Ephimeris[counter].Moon, transits.Ephimeris[counter+1].Moon, t1.In(zone))
		}

		counter++

	}
	// logrus.Infof("Checked Paadas %d/%d total ", counter, numItems)
	transits.Paadas = paadas

}

// getRasiPanchang()
// get the first rasi number, by getting the rasi associated with first moon-star in panchange
// take the next item, and calculate the degrees associated with next rasi
// since angular motion of earth and planet is constant in a day,
// calculate the transit time for next item (here rasi)
func (transits *Transits) getRasiPanchang() {

	// zone, _ := time.LoadLocation("Asia/Kolkata")

	rasis := make([]*astroV1.Transit, 0)
	counter := 0
	t := transits.StartTime.AsTime()
	numItems := len(transits.Ephimeris) - 1
	currentItem := GetRasiNumber(transits.Ephimeris[0].Moon)
	// logrus.Infof("Checking rasi %d total ", numItems)

	for counter < numItems {

		nextItem := currentItem + 1
		if nextItem > TOTAL_RASIS {
			nextItem = 1
		}

		nextItemDegrees := RASI_DEGREES * float64(nextItem-1)
		// logrus.Infof("Checking rasi counter (%d ), Item (%d - %d), transit degrees (%f)", counter, currentItem, nextItem, nextItemDegrees)

		transitMinutes, err := GetTransitTime(transits.Ephimeris[counter].Moon, transits.Ephimeris[counter+1].Moon, HOUR_MINUTES, nextItemDegrees)
		if err == nil {
			currentItem++
			if currentItem > TOTAL_RASIS {
				currentItem = 1
			}

			t1 := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*HOUR_MINUTES))

			currentTransit := &astroV1.Transit{
				Name:    int32(currentItem),
				Transit: timestamppb.New(t1),
			}

			rasis = append(rasis, currentTransit)

			// logrus.Infof("Rasi counter (%d) %d, %f<->%f,  %v", counter, currentItem, transits.Ephimeris[counter].Moon, transits.Ephimeris[counter+1].Moon, t1.In(zone))
		}

		counter++

	}
	logrus.Infof("Checked Rasis %d/%d total ", counter, numItems)
	transits.Rasis = rasis

}

func (transits *Transits) getTithiPanchang() {
	// zone, _ := time.LoadLocation("Asia/Kolkata")

	tithis := make([]*astroV1.Transit, 0)

	counter := 0
	numItems := len(transits.Ephimeris) - 1
	currentTithi := GetTithiNumber(transits.Ephimeris[0].Sun, transits.Ephimeris[0].Moon)

	for counter < numItems {

		t := transits.StartTime.AsTime()

		nextTithi := currentTithi + 1
		if nextTithi > TOTAL_TITHIS {
			nextTithi = 1
		}

		nextTithiDegrees := TITHI_DEGREES * float64(nextTithi-1)
		//logrus.Infof("Checking star current (%d), next (%d), transit degrees (%f)", currentStar, nextStar, nextStarDegrees)

		startTithiDegree := GetTithiDegree(transits.Ephimeris[counter].Sun, transits.Ephimeris[counter].Moon)
		endTithiDegree := GetTithiDegree(transits.Ephimeris[counter+1].Sun, transits.Ephimeris[counter+1].Moon)

		transitMinutes, err := GetTransitTime(startTithiDegree, endTithiDegree, HOUR_MINUTES, nextTithiDegrees)
		if err == nil {
			currentTithi++
			if currentTithi > TOTAL_TITHIS {
				currentTithi = 1
			}
			// paksha := GetPaksha(currentTithi)
			t := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*HOUR_MINUTES))

			tithiTransit := &astroV1.Transit{
				Name:    int32(currentTithi),
				Transit: timestamppb.New(t),
			}

			tithis = append(tithis, tithiTransit)

			// logrus.Infof("%s(%s), %f<->%f @ %v", PAKSHA_NAMES[paksha], TITHI_NAMES[currentTithi], startTithiDegree, endTithiDegree, t.In(zone))
		}

		counter++

	}

	transits.Tithis = tithis

}

func (transits *Transits) getKaranaPanchang() {
	// zone, _ := time.LoadLocation("Asia/Kolkata")

	karanas := make([]*astroV1.Transit, 0)

	counter := 0
	numItems := len(transits.Ephimeris) - 1
	currentKarana := GetKaranaNumber(transits.Ephimeris[0].Sun, transits.Ephimeris[0].Moon)

	for counter < numItems {
		t := transits.StartTime.AsTime()
		nextKarana := currentKarana + 1
		if nextKarana > TOTAL_KARANAS {
			nextKarana = 1
		}

		nextKaranaDegrees := KARANA_DEGREES * float64(nextKarana-1)
		//logrus.Infof("Checking star current (%d), next (%d), transit degrees (%f)", currentStar, nextStar, nextStarDegrees)

		startKaranaDegree := GetTithiDegree(transits.Ephimeris[counter].Sun, transits.Ephimeris[counter].Moon)
		endKaranaDegree := GetTithiDegree(transits.Ephimeris[counter+1].Sun, transits.Ephimeris[counter+1].Moon)

		transitMinutes, err := GetTransitTime(startKaranaDegree, endKaranaDegree, HOUR_MINUTES, nextKaranaDegrees)
		if err == nil {
			currentKarana++
			if currentKarana > TOTAL_KARANAS {
				currentKarana = 1
			}
			t := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*HOUR_MINUTES))

			karanas = append(karanas, &astroV1.Transit{
				Name:    int32(currentKarana),
				Transit: timestamppb.New(t),
			})

			// logrus.Infof("%s, %f<->%f @ %v", KARANA_NAMES[currentKarana], startKaranaDegree, endKaranaDegree, t.In(zone))
		}

		counter++

	}

	transits.Karanas = karanas

}

func (transits *Transits) getYogaPanchang() {

	// zone, _ := time.LoadLocation("Asia/Kolkata")

	yogas := make([]*astroV1.Transit, 0)

	counter := 0
	numItems := len(transits.Ephimeris) - 1
	currentYoga := GetYogaNumber(transits.Ephimeris[0].Sun, transits.Ephimeris[0].Moon)
	// logrus.Infof("Starting Yoga (%d)", currentYoga)
	for counter < numItems {
		t := transits.StartTime.AsTime()
		nextYoga := currentYoga + 1
		if nextYoga > TOTAL_YOGAS {
			nextYoga = 1
		}

		nextYogaDegrees := YOGA_DEGREES * float64(nextYoga-1)

		startYogaDegree := GetYogaDegree(transits.Ephimeris[counter].Sun, transits.Ephimeris[counter].Moon)
		endYogaDegree := GetYogaDegree(transits.Ephimeris[counter+1].Sun, transits.Ephimeris[counter+1].Moon)

		transitMinutes, err := GetTransitTime(startYogaDegree, endYogaDegree, HOUR_MINUTES, nextYogaDegrees)
		if err == nil {
			currentYoga++
			if currentYoga > TOTAL_YOGAS {
				currentYoga = 1
			}

			t := t.Add(time.Minute*time.Duration(transitMinutes) + time.Minute*time.Duration(counter*HOUR_MINUTES))

			yogas = append(yogas, &astroV1.Transit{
				Name:    int32(currentYoga),
				Transit: timestamppb.New(t),
			})

			// logrus.Infof(" %s, %f<->%f @ %v", YOGA_NAMES[currentYoga], startYogaDegree, endYogaDegree, t.In(zone))
		}

		counter++

	}

	transits.Yogas = yogas

}
