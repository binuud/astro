package main

import (
	"fmt"
	"time"

	"github.com/dronasys-com/astro/cmd/db"

	"github.com/dronasys-com/astro/pkg/servers/astroServer"
	"github.com/sirupsen/logrus"
)

const (
	START_YEAR              = 1950
	START_MONTH             = 1
	START_DAY               = 1
	END_YEAR                = 2100
	INTERVAL                = 720
	BULK_UPDATE_BUCKET_SIZE = 300
)

func createRecordNameFromTime(t time.Time) string {
	return fmt.Sprintf("##TEST-%02d/%02d/%04d-%02d:%02d", t.UTC().Day(), t.UTC().Month(), t.UTC().Year(), t.UTC().Hour(), t.UTC().Minute())
}

func updateCharts() error {
	tz, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}
	startTime := time.Date(int(START_YEAR), time.Month(START_MONTH), int(START_DAY), int(0), int(0), int(0), 0, tz)
	nextTime := startTime
	year := START_YEAR
	bucketCount := 0

	astroBucket := make([]interface{}, 0)

	for year < END_YEAR {
		bucketCount++
		recordName := createRecordNameFromTime(nextTime)
		astroObject, err := astroServer.CreateAstroObjectForTime(recordName, nextTime)
		if err != nil {
			return err
		}
		astroBucket = append(astroBucket, astroObject)

		if bucketCount >= BULK_UPDATE_BUCKET_SIZE {
			logrus.Info("Bulk insert with last record (%v)", recordName)
			astroServer.BulkInsert(astroBucket)
			// reset counters
			bucketCount = 0
			astroBucket = make([]interface{}, 0)
		}

		nextTime = nextTime.Add(time.Minute * time.Duration(INTERVAL))
		year = nextTime.UTC().Year()

	}
	if bucketCount > 0 {

		logrus.Info("Bulk insert with remaining record ")
		astroServer.BulkInsert(astroBucket)
		// reset counters
		bucketCount = 0
		astroBucket = make([]interface{}, 0)
	}
	return nil
}

func main() {

	db.Init()

	logrus.Info("Astro Chart DB Creator - Trance !")
	logrus.Info("--------->>>> Creates charts for 15 minute intervals  <<<----------")
	updateCharts()

}
