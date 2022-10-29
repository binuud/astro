package astro

import (
	"fmt"
	"math"
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"github.com/sirupsen/logrus"
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

// GetTransitTime
// arguments
// startDegree - the startDegree of the planet/object
// endDegree - the endDegree of the planet/object at the end of the given duration ( eg moon longitudes at start and end of day)
// duration - the duration in minutes when the start and end degree of the object was observed
// transitionDegree - the degree at which transition is expected to happen (eg: this can be boundaries like 30,60,90 for rasis)
// returns minutes(time) at which the transit happens
func GetTransitTime(startDegree float64, endDegree float64, duration int64, transitionDegree float64) (int64, error) {
	// if transitionDegree < startDegree {
	// 	err := fmt.Errorf("transition degree cannot be less than start degree")
	// 	return 0, err
	// }
	// if transitionDegree > endDegree {
	// 	err := fmt.Errorf("transition degree cannot be greater than end degree")
	// 	return 0, err
	// }

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
	// if (endDegree - transitionDegree) > 350 {
	// 	err := fmt.Errorf("preventing rollover error of longitude")
	// 	return 0, err
	// }
	logrus.Infof("(%f)<-(%f)->(%f)", startDegree, transitionDegree, endDegree)
	if startDegree < transitionDegree && endDegree > transitionDegree {

		t := int64(((transitionDegree - startDegree) / totalDegree) * float64(duration))
		return t, nil
	}
	err := fmt.Errorf("transition degree cannot be greater than end degree")
	return 0, err
}
