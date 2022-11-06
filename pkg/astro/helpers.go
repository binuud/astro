package astro

import (
	"fmt"
	"math"
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeBirthData(in *astroV1.BirthData) (*astroV1.BirthData, error) {
	tz, err := time.LoadLocation(in.Tob.Zone)
	if err != nil {
		return nil, err
	}
	t := time.Date(int(in.Tob.Y), time.Month(in.Tob.M), int(in.Tob.D), int(in.Tob.Hour), int(in.Tob.Min), int(in.Tob.Sec), 0, tz)
	in.Tob.Tobtime = timestamppb.New(t)
	return in, nil
}

// GetStarNumber
// arguments - degree (float) the longitude of the astronomical object
// returns the star number (int)
// eg: if degree is 14, the star number would be 2 -> Krittika
// eg: if degree is 10.2, the star number would be 1 -> Aswini
func GetStarNumber(degree float64) int {
	if degree > 360.0 {
		degree = math.Mod(degree, 360)
	}
	return int(degree/STAR_DEGREES) + 1
}

func GetPaksha(tithi int) int {
	if tithi <= 15 {
		return SHUKRA_PAKSHA
	}
	return KRISHNA_PAKSHA
}

func GetKaranaNumber(sunDegree float64, moonDegree float64) int {
	return (int(GetTithiDegree(sunDegree, moonDegree)) / KARANA_DEGREES) + 1
}

func GetTithiNumber(sunDegree float64, moonDegree float64) int {
	return (int(GetTithiDegree(sunDegree, moonDegree)) / TITHI_DEGREES) + 1
}

func GetTithiDegree(sunDegree float64, moonDegree float64) float64 {
	tithiDegree := moonDegree - sunDegree
	if tithiDegree < 0 {
		tithiDegree = 360 + tithiDegree
	}
	return tithiDegree

}

func GetYogaNumber(sunDegree float64, moonDegree float64) int {
	return (int(GetYogaDegree(sunDegree, moonDegree)) * 60 / 800) + 1
}

func GetYogaDegree(sunDegree float64, moonDegree float64) float64 {
	yogaDegree := moonDegree + sunDegree
	if yogaDegree > 360 {
		yogaDegree = yogaDegree - 360
	}
	return yogaDegree

}

// GetTransitTime
// arguments
// startDegree - the startDegree of the planet/object
// endDegree - the endDegree of the planet/object at the end of the given duration ( eg moon longitudes at start and end of day)
// duration - the duration in minutes when the start and end degree of the object was observed
// transitionDegree - the degree at which transition is expected to happen (eg: this can be boundaries like 30,60,90 for rasis)
// returns minutes(time) at which the transit happens
func GetTransitTime(startDegree float64, endDegree float64, duration int64, transitionDegree float64) (int64, error) {

	if endDegree < startDegree {
		endDegree = 360 + endDegree
		if transitionDegree == 0.0 {
			transitionDegree = 360.0
		}
	}

	totalDegree := endDegree - startDegree
	if totalDegree == 0 {
		err := fmt.Errorf("no movement of object seen")
		return 0, err
	}

	// logrus.Infof("(%f)<-(%f)->(%f)", startDegree, transitionDegree, endDegree)
	if startDegree < transitionDegree && endDegree > transitionDegree {

		t := int64(((transitionDegree - startDegree) / totalDegree) * float64(duration))
		return t, nil
	}
	err := fmt.Errorf("transition degree has to be in between start and end degree")
	return 0, err
}
