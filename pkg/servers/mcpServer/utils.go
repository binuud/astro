package mcpServer

import (
	"fmt"
	"time"

	"github.com/binuud/astro/pkg/astro"
	"github.com/binuud/astro/pkg/servers/citiesServer"
	"github.com/sirupsen/logrus"
)

func parseDOB(dobStr string) (time.Time, error) {
	const layout = "02/01/2006"

	return time.Parse(layout, dobStr)
}

func parseTOB(tobStr string) (time.Time, error) {
	const layout = "15:04"

	return time.Parse(layout, tobStr)
}

func CalculateAge(dobStr string) (int, error) {
	// Go uses the specific reference date Mon Jan 2 15:04:05 MST 2006 for layouts
	const layout = "02/01/2006"

	dob, err := parseDOB(dobStr)
	if err != nil {
		return 0, fmt.Errorf("invalid date format (expected dd/mm/yyyy): %w", err)
	}

	now := time.Now()
	if dob.After(now) {
		return 0, fmt.Errorf("date of birth cannot be in the future")
	}

	age := now.Year() - dob.Year()

	// Subtract 1 if the birthday hasn't occurred yet in the current year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age, nil
}

func GetLatLongForCity(city string, countryCode string) (string, float64, float64, error) {
	db := citiesServer.NewCityDB()
	cities, err := db.SearchCity(city, countryCode)
	if err != nil {
		return "", 0, 0, err
	}
	logrus.Infof("found cities for search %s %d", city, len(cities))
	// return the first cities lat long
	return cities[0].Name, cities[0].Latitude, cities[0].Longitude, nil
}

// PrintChartDescription
// Print the chart description in plain english
func _getAstrologicalChart(chart *astro.Chart) *Output_GetAstrologicalChart {
	planetNames := []string{"asc", "sun", "moon", "mars", "mercury", "jupiter", "venus", "saturn", "rahu", "ketu", "neptune", "uranus", "pluto"}
	output := ""
	zodiac := make(map[string]string, 0)
	house := make(map[string]string, 0)
	for _, planetName := range planetNames {
		planet := chart.Planets.PlanetsMap[planetName]
		motion := ""
		if planet.Speed < 0 {
			motion = "Retrograde"
		}
		if motion == "Retrograde" {
			zodiac[planet.Name] = fmt.Sprintf("Retrograde in %s at %.2f", astro.ENGLISH_RASI_NAMES[planet.Rasi], planet.RasiLongitude)
			house[planet.Name] = fmt.Sprintf(" Retrograde in %d house", planet.BhavaH)
		} else {
			zodiac[planet.Name] = fmt.Sprintf("Direct in %s at %.2f", astro.ENGLISH_RASI_NAMES[planet.Rasi], planet.RasiLongitude)
			house[planet.Name] = fmt.Sprintf("Direct in %d house", planet.BhavaH)
		}

		output = output + fmt.Sprintf("%s %s in %s. ", planet.Name, motion, astro.ENGLISH_RASI_NAMES[planet.Rasi])

	}
	return &Output_GetAstrologicalChart{
		Text:   output,
		Zodiac: zodiac,
		House:  house,
	}
}
