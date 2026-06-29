package main

import (
	"fmt"
	"time"

	"github.com/dronasys-com/astro/pkg/servers/astroServer"
	"github.com/sirupsen/logrus"
)

const (
	START_YEAR              = 2026
	START_MONTH             = 6
	START_DAY               = 29
	END_YEAR                = 2026
	INTERVAL                = 86400
	BULK_UPDATE_BUCKET_SIZE = 300
)

func createRecordNameFromTime(t time.Time) string {
	return fmt.Sprintf("##TEST-%02d/%02d/%04d-%02d:%02d", t.UTC().Day(), t.UTC().Month(), t.UTC().Year(), t.UTC().Hour(), t.UTC().Minute())
}

func getDeclinations() error {
	tz, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}
	startTime := time.Date(int(START_YEAR), time.Month(START_MONTH), int(START_DAY), int(0), int(0), int(0), 0, tz)
	nextTime := startTime
	year := START_YEAR
	bucketCount := 0

	for year < END_YEAR {
		bucketCount++
		recordName := createRecordNameFromTime(nextTime)
		astroObject, err := astroServer.CreateAstroObjectForTime(recordName, nextTime)
		if err != nil {
			return err
		}

		nextTime = nextTime.Add(time.Minute * time.Duration(INTERVAL))
		year = nextTime.UTC().Year()

	}

	return nil
}

func main() {

	logrus.Info("Get declination of planets !")
	getDeclinations()

}
