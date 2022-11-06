package astro

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
)

const (
	FIELD_PLANET      = 0
	FIELD_LONGITUDE   = 1
	FIELD_LATITUDE    = 2
	FIELD_DISTANCE    = 3
	FIELD_DECLINATION = 4
	FIELD_SPEED       = 5
)

type Chart struct {
	BirthData *astroV1.BirthData
	Planets   *astroV1.Planets
}

func NewAstroChart(birthData *astroV1.BirthData) (*Chart, error) {

	birthData, err := MakeBirthData(birthData)
	if err != nil {
		logrus.Errorf("Cannot parse time data for astro chart err(%v)", err)
	}

	ephimeris := NewEphimeris(birthData)
	planetsEphimeris, err := ephimeris.GetPlanetsEphimeris()
	if err != nil {
		return nil, err
	}

	moonEphimeris, err := ephimeris.GetMoonEphimeris()
	if err != nil {
		return nil, err
	}

	chart := &Chart{
		BirthData: birthData,
		Planets: &astroV1.Planets{
			Ephimeris: planetsEphimeris,
		},
	}
	logrus.Infof(" Moon (%v)", moonEphimeris)
	logrus.Infof(" Planet (%v)", planetsEphimeris)
	chart.ParseEphimeris()
	return chart, nil

}

func (chart *Chart) ParseEphimeris() {

	longitudes := strings.Split(chart.Planets.Ephimeris, "\n")
	if len(longitudes) < 3 {
		logrus.Errorf("No ephimeris data available")
		return
	}
	// fmt.Printf("data %v", longitudes)

	var planetMap = make(map[string]*astroV1.PlanetDetails)
	for _, item := range longitudes {
		// logrus.Infof("Parsing string (%v)", item)
		fields := strings.Fields(item)
		if len(fields) > 0 {
			planetName := ""
			switch fields[0] {
			case "Ascendant":
				planetName = PLANET_ASC
			case "Sun":
				planetName = PLANET_SUN
			case "Mercury":
				planetName = PLANET_MERCURY
			case "Venus":
				planetName = PLANET_VENUS
			case "Mars":
				planetName = PLANET_MARS
			case "Jupiter":
				planetName = PLANET_JUPITER
			case "Moon":
				planetName = PLANET_MOON
			case "Saturn":
				planetName = PLANET_SATURN
			case "Uranus":
				planetName = PLANET_URANUS
			case "Neptune":
				planetName = PLANET_NEPTUNE
			case "Pluto":
				planetName = PLANET_PLUTO
			case "mean":
				if fields[1] == "Node" {
					fields = fields[1:] // shift the fields array
					planetName = PLANET_RAHU
				}
			}
			if planetName != "" {
				longitude, err := strconv.ParseFloat(fields[FIELD_LONGITUDE], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}

				latitude := 0.0
				declination := 0.0
				speed := 0.0
				distance := 0.0

				if planetName != PLANET_ASC {
					latitude, err = strconv.ParseFloat(fields[FIELD_LATITUDE], 64)
					if err != nil {
						logrus.Errorf("Cannot parse latitude err(%v)", err)
					}

					if planetName != PLANET_MOON {
						distance, err = strconv.ParseFloat(fields[FIELD_DISTANCE], 64)
						if err != nil {
							logrus.Errorf("Cannot parse distance err(%v)", err)
						}
					}

					declination, err = strconv.ParseFloat(fields[FIELD_DECLINATION], 64)
					if err != nil {
						logrus.Errorf("Cannot parse declination err(%v)", err)
					}

					speed, err = strconv.ParseFloat(fields[FIELD_SPEED], 64)
					if err != nil {
						logrus.Errorf("Cannot parse speed err(%v)", err)
					}
				}

				planetMap[planetName] = &astroV1.PlanetDetails{
					Name:        planetName,
					Longitude:   longitude,
					Latitude:    latitude,
					Distance:    distance,
					Declination: declination,
					Speed:       speed,
				}
			}
		}

	}
	// fmt.Printf("Final data %v", planetMap)
	chart.Planets.PlanetsMap = planetMap

	chart.CalculateAuxillaryPlanets()
	chart.CalculateCharts()
	chart.CalculateBhavaChart()
	chart.CalculateMeta()
}

func (chart *Chart) CalculateAuxillaryPlanets() {
	ketuLongitude := chart.Planets.PlanetsMap[PLANET_RAHU].Longitude + 180.0
	if ketuLongitude > 360.0 {
		ketuLongitude = ketuLongitude - 360.0
	}
	chart.Planets.PlanetsMap[PLANET_KETU] = &astroV1.PlanetDetails{
		Name:      PLANET_KETU,
		Longitude: float64(ketuLongitude),
	}
}

func (chart *Chart) CalculateCharts() {
	for _, planet := range chart.Planets.PlanetsMap {

		planet.Rasi = CalculateRasi(float64(planet.Longitude))
		planet.Star = CalculateStar(float64(planet.Longitude))
		planet.Paada = CalculatePaada(float64(planet.Longitude))
		planet.Amsa2 = CalculateNthHarmonic(planet.Longitude, HARMONIC_HORA)
		planet.Amsa9 = CalculateNthHarmonic(planet.Longitude, HARMONIC_NAVAMSA)
		planet.Amsa10 = CalculateNthHarmonic(planet.Longitude, HARMONIC_DASAMSA)
		planet.Amsa27 = CalculateNthHarmonic(planet.Longitude, HARMONIC_NAKSATHRAMSA)
		planet.Amsa60 = CalculateNthHarmonic(planet.Longitude, HARMONIC_SASHTYAMSA)
	}
}

func (chart *Chart) CalculateMeta() {
	for _, planet := range chart.Planets.PlanetsMap {
		CalculatePlanetMeta(planet)
	}
}

func (chart *Chart) CalculateBhavaChart() {
	ascDegrees := chart.Planets.PlanetsMap[PLANET_ASC].Longitude
	bhavaCorrection := math.Mod(float64(ascDegrees), RASI_DEGREES) - 15.0
	for _, planet := range chart.Planets.PlanetsMap {
		bavaDeg := float64(planet.Longitude) - bhavaCorrection
		//	print "$bavaDeg <br/>";
		if bavaDeg > 360.0 {
			bavaDeg = bavaDeg - 360.0
		}

		if bavaDeg < 0 {
			bavaDeg = 360 + bavaDeg
		}

		planet.Bhava = CalculateRasi(float64(bavaDeg))
	}
	asc := chart.Planets.PlanetsMap[PLANET_ASC].Bhava
	for _, planet := range chart.Planets.PlanetsMap {
		bhava := planet.Bhava - asc + 1
		if bhava < 1 {
			bhava = TOTAL_RASIS + bhava
		}
		planet.BhavaH = bhava

	}
}

func (chart *Chart) PrintChartTable() {
	fmt.Printf("\n\n********************************************\n")
	fmt.Printf("    Planet   R   B  BH 2H  9H 10H 27H 60H   S   P   Long  Dec Sp\n")
	for _, planet := range chart.Planets.PlanetsMap {
		fmt.Printf("%10s  %02d  %02d  %02d  %02d  %02d  %02d  %02d  %02d  %02d  %02d   %09.2f %09.2f %09.2f\n", planet.Name, planet.Rasi, planet.Bhava, planet.BhavaH, planet.Amsa2, planet.Amsa9, planet.Amsa10, planet.Amsa27, planet.Amsa60, planet.Star, planet.Paada, planet.Longitude, planet.Declination, planet.Speed)
	}
	fmt.Printf("\n\n********************************************\n")
}
