package citiesServer

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	citiesV1 "github.com/binuud/astro/gen/go/v1/astro"
	"github.com/sirupsen/logrus"
)

const DB_FILE = "cities.csv"

type CityDB struct {
	cities []*citiesV1.CityInfo
}

func NewCityDB() *CityDB {
	cityDBInstance := &CityDB{
		cities: make([]*citiesV1.CityInfo, 0),
	}
	cityDBInstance.cities = cityDBInstance.initDB()
	return cityDBInstance
}

func (c *CityDB) initDB() []*citiesV1.CityInfo {

	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "./data/"
	}
	csvFile := fmt.Sprintf("%s%s", dataPath, DB_FILE)
	csvData, err := c.readCsvFile(csvFile)
	if err != nil {
		logrus.Errorf("err cannot load cities database csv file %s", csvFile)
		panic(err)
	}
	return c.createInMemDB(csvData)
}

func (c *CityDB) createInMemDB(data [][]string) []*citiesV1.CityInfo {
	var cities []*citiesV1.CityInfo
	// 0 id,name,
	// 2 state_id,state_code,state_name,
	// 5 country_id,country_code,country_name,
	// 8 latitude,longitude
	// 10 wikiDataId

	for i, line := range data {
		if i > 0 { // omit header line
			var rec = citiesV1.CityInfo{}
			rec.Index = int64(i - 1)
			for col, field := range line {

				if col == 0 { // id
					id, err := strconv.Atoi(field)
					if err != nil {
						logrus.Fatalf("Error parsing id field from (%s) (%v)", line, err)
					}
					rec.Id = int64(id)
				} else if col == 1 { // name
					rec.Name = field
				} else if col == 2 { // state_id
					id, err := strconv.Atoi(field)
					if err != nil {
						logrus.Fatalf("Error parsing state_id field from (%s) (%v)", line, err)
					}
					rec.StateId = int64(id)
				} else if col == 3 { // state_code
					rec.StateCode = field
				} else if col == 4 { // state_name
					rec.StateName = field
				} else if col == 5 { // country_id
					id, err := strconv.Atoi(field)
					if err != nil {
						logrus.Fatalf("Error parsing country_id field from (%s) (%v)", line, err)
					}
					rec.CountryId = int64(id)
				} else if col == 6 { // country_code
					rec.CountryCode = field
				} else if col == 7 { // country_name
					rec.CountryName = field
				} else if col == 8 { // latitude
					lat, err := strconv.ParseFloat(field, 64)
					if err != nil {
						logrus.Fatalf("Error parsing latitude field from (%s) (%v)", line, err)
					}
					rec.Latitude = lat
				} else if col == 9 { // longitude
					lng, err := strconv.ParseFloat(field, 64)
					if err != nil {
						logrus.Fatalf("Error parsing longitude field from (%s) (%v)", line, err)
					}
					rec.Longitude = lng
				} else if col == 10 { // wikiDataId
					rec.WikiDataId = field
				}
			}
			// logrus.Infof("Parsed (%v)", rec)
			cities = append(cities, &rec)
		}
	}
	return cities
}

func (c *CityDB) readCsvFile(filename string) ([][]string, error) {
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer fileContent.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func (c *CityDB) GetCityInfo(id int) (*citiesV1.CityInfo, error) {
	if id < 0 || id >= len(c.cities) {
		err := fmt.Errorf("Invalid id or out of range data")
		return nil, err
	}
	return c.cities[id], nil
}

func (c *CityDB) SearchCity(city string, countryCode string) ([]*citiesV1.CityInfo, error) {
	if len(countryCode) != 2 {
		err := fmt.Errorf("Invalid country code, country code is a 2 letter string, eg: IN, US")
		return nil, err
	}
	result := make([]*citiesV1.CityInfo, 0)
	count := 1
	limit := 10
	for _, item := range c.cities {
		if item.CountryCode == countryCode {
			if strings.Contains(item.Name, city) {
				count++
				result = append(result, item)
			}

		}
		if count > limit {
			break
		}
	}
	return result, nil
}
