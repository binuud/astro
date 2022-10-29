package astro

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astrocalc/go"
	"github.com/sirupsen/logrus"
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

	var err error
	var planetMap = make(map[string]*astroV1.PlanetDetails)
	for _, item := range longitudes {
		// logrus.Infof("Parsing string (%v)", item)
		fields := strings.Fields(item)
		if len(fields) > 0 {
			longitude := 0.0
			planetName := ""
			switch fields[0] {
			case "Ascendant":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_ASC
			case "Sun":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_SUN
			case "Mercury":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_MERCURY

			case "Venus":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_VENUS

			case "Mars":
				longitude, err = strconv.ParseFloat(fields[1], 32)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_MARS

			case "Jupiter":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_JUPITER

			case "Moon":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_MOON

			case "Saturn":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_SATURN

			case "Uranus":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_URANUS

			case "Neptune":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_NEPTUNE

			case "Pluto":
				longitude, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					logrus.Errorf("Cannot parse longitudes err(%v)", err)
				}
				planetName = PLANET_PLUTO

			case "mean":
				if fields[1] == "Node" {
					// logrus.Printf("\n\n %%%%%%%% (%v)", fields)
					longitude, err = strconv.ParseFloat(fields[2], 64)
					if err != nil {
						logrus.Errorf("Cannot parse longitudes err(%v)", err)
					}
					planetName = PLANET_RAHU
				}

			}
			if planetName != "" {
				planetMap[planetName] = &astroV1.PlanetDetails{
					Name:      planetName,
					Longitude: longitude,
				}
			}
		}

	}
	fmt.Printf("Final data %v", planetMap)
	chart.Planets.PlanetsMap = planetMap

	chart.CalculateAuxillaryPlanets()
	chart.CalculateRasiChart()
	chart.CalculateBhavaChart()
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

func (chart *Chart) CalculateRasiChart() {
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
}

func (chart *Chart) PrintChartTable() {
	fmt.Printf("\n\n********************************************\n")
	fmt.Printf("    Planet   R   B  2H  9H 10H 27H 60H   S   P   Longitude\n")
	for _, planet := range chart.Planets.PlanetsMap {
		fmt.Printf("%10s  %02d  %02d  %02d  %02d  %02d  %02d  %02d  %02d  %02d   %03.2f\n", planet.Name, planet.Rasi, planet.Bhava, planet.Amsa2, planet.Amsa9, planet.Amsa10, planet.Amsa27, planet.Amsa60, planet.Star, planet.Paada, planet.Longitude)
	}
	fmt.Printf("\n\n********************************************\n")
}
