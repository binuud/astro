package astro

import (
	"fmt"

	"golang.org/x/exp/slices"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
)

type Planet_Meta struct {
	Name           string
	Number         int
	Exalted        int
	Deblited       int
	OwnHouse       []int
	MoolTrikona    int
	FriendPlanets  []string
	EnemyPlanets   []string
	NeutralPlanets []string
}

// var PLANET_METAS = make(map[string]*Planet_Meta)

var PLANET_METAS = map[string]*Planet_Meta{
	PLANET_SUN: &Planet_Meta{
		Name:        "Sun",
		Number:      1,
		Exalted:     1,
		Deblited:    7,
		OwnHouse:    []int{5},
		MoolTrikona: 5,
	},
	PLANET_MOON: &Planet_Meta{
		Name:        "Moon",
		Number:      2,
		Exalted:     2,
		Deblited:    8,
		OwnHouse:    []int{4},
		MoolTrikona: 2,
	},
	PLANET_MERCURY: &Planet_Meta{
		Name:        "Mercury",
		Number:      3,
		Exalted:     6,
		Deblited:    12,
		OwnHouse:    []int{3, 6},
		MoolTrikona: 6,
	},
	PLANET_VENUS: &Planet_Meta{
		Name:        "Venus",
		Number:      4,
		Exalted:     12,
		Deblited:    6,
		OwnHouse:    []int{2, 7},
		MoolTrikona: 7,
	},
	PLANET_MARS: &Planet_Meta{
		Name:        "Mars",
		Number:      5,
		Exalted:     10,
		Deblited:    4,
		OwnHouse:    []int{1, 8},
		MoolTrikona: 1,
	},
	PLANET_JUPITER: &Planet_Meta{
		Name:        "Jupiter",
		Number:      6,
		Exalted:     4,
		Deblited:    10,
		OwnHouse:    []int{9, 12},
		MoolTrikona: 9,
	},
	PLANET_SATURN: &Planet_Meta{
		Name:        "Saturn",
		Number:      7,
		Exalted:     7,
		Deblited:    1,
		OwnHouse:    []int{10, 11},
		MoolTrikona: 11,
	},
}

func CalculatePlanetMeta(planet *astroV1.PlanetDetails) error {

	planetMeta, ok := PLANET_METAS[planet.Name]
	if ok == false {
		err := fmt.Errorf("planet not found in planet meta")
		return err
	}

	if planet.Rasi == int64(planetMeta.Exalted) {
		planet.Exalted = true
	} else {
		planet.Exalted = false
	}

	if planet.Rasi == int64(planetMeta.Deblited) {
		planet.Deblited = true
	} else {
		planet.Deblited = false
	}

	if planet.Rasi == int64(planetMeta.MoolTrikona) {
		planet.MoolTrikona = true
	} else {
		planet.MoolTrikona = false
	}

	if planet.Rasi == int64(planetMeta.MoolTrikona) {
		planet.MoolTrikona = true
	} else {
		planet.MoolTrikona = false
	}

	if slices.Index(planetMeta.OwnHouse, int(planet.Rasi)) != -1 {
		planet.OwnHouse = true
	}

	return nil
}
