package astro

import (
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeBirthData(in *astroV1.BirthData) (*astroV1.BirthData, error) {
	tz, err := time.LoadLocation(in.Dob.Zone)
	if err != nil {
		return nil, err
	}
	t := time.Date(int(in.Dob.Y), time.Month(in.Dob.M), int(in.Dob.D), int(in.Dob.Hour), int(in.Dob.Min), int(in.Dob.Sec), 0, tz)
	in.Dob.Tobtime = timestamppb.New(t)
	return in, nil
}
