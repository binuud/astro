package astro

import (
	"fmt"

	"golang.org/x/exp/slices"

	aV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
)

const (
	GENDER_MALE    = "GENDER_MALE"
	GENDER_FEMALE  = "GENDER_FEMALE"
	GENDER_NEUTRAL = "GENDER_NEUTRAL"

	STATE_AIR   = "STATE_AIR"
	STATE_EARTH = "STATE_EARTH"
	STATE_FIRE  = "STATE_FIRE"
	STATE_WATER = "STATE_WATER"
)

type Planet_Meta struct {
	Name           string
	Number         int
	Exalted        int
	Deblited       int
	OwnHouse       []int
	MoolTrikona    int
	Aspects        []int
	FriendPlanets  []int
	EnemyPlanets   []int
	NeutralPlanets []int
	FriendHouses   []int
	EnemyHouses    []int
	NeutralHouses  []int
	Tags           []string
}

// var PLANET_METAS = make(map[string]*Planet_Meta)

var PLANET_METAS = map[string]*Planet_Meta{
	PLANET_SUN: &Planet_Meta{
		Name:           "Sun",
		Number:         int(aV1.AstroConstants_Sun),
		Exalted:        int(aV1.AstroConstants_Aries),
		Deblited:       int(aV1.AstroConstants_Libra),
		OwnHouse:       []int{int(aV1.AstroConstants_Leo)},
		MoolTrikona:    int(aV1.AstroConstants_Leo),
		FriendPlanets:  []int{int(aV1.AstroConstants_Moon), int(aV1.AstroConstants_Mars), int(aV1.AstroConstants_Jupiter)},
		EnemyPlanets:   []int{int(aV1.AstroConstants_Saturn), int(aV1.AstroConstants_Venus)},
		NeutralPlanets: []int{int(aV1.AstroConstants_Mercury)},
		FriendHouses:   []int{int(aV1.AstroConstants_Cancer), int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio), int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces)},
		EnemyHouses:    []int{int(aV1.AstroConstants_Capricon), int(aV1.AstroConstants_Aquarius), int(aV1.AstroConstants_Taurus), int(aV1.AstroConstants_Libra)},
		NeutralHouses:  []int{int(aV1.AstroConstants_Gemini), int(aV1.AstroConstants_Virgo)},
	},
	PLANET_MOON: &Planet_Meta{
		Name:           "Moon",
		Number:         int(aV1.AstroConstants_Moon),
		Exalted:        int(aV1.AstroConstants_Taurus),
		Deblited:       int(aV1.AstroConstants_Scorpio),
		OwnHouse:       []int{int(aV1.AstroConstants_Cancer)},
		MoolTrikona:    int(aV1.AstroConstants_Taurus),
		FriendPlanets:  []int{int(aV1.AstroConstants_Sun), int(aV1.AstroConstants_Mercury)},
		EnemyPlanets:   []int{},
		NeutralPlanets: []int{int(aV1.AstroConstants_Mars), int(aV1.AstroConstants_Jupiter), int(aV1.AstroConstants_Venus), int(aV1.AstroConstants_Saturn)},
		FriendHouses:   []int{int(aV1.AstroConstants_Leo), int(aV1.AstroConstants_Gemini), int(aV1.AstroConstants_Virgo)},
		EnemyHouses:    []int{},
		NeutralHouses: []int{int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio),
			int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces),
			int(aV1.AstroConstants_Taurus), int(aV1.AstroConstants_Libra),
			int(aV1.AstroConstants_Capricon), int(aV1.AstroConstants_Aquarius)},
	},
	PLANET_MERCURY: &Planet_Meta{
		Name:           "Mercury",
		Number:         int(aV1.AstroConstants_Mercury),
		Exalted:        int(aV1.AstroConstants_Virgo),
		Deblited:       int(aV1.AstroConstants_Pisces),
		OwnHouse:       []int{int(aV1.AstroConstants_Gemini), int(aV1.AstroConstants_Virgo)},
		MoolTrikona:    int(aV1.AstroConstants_Virgo),
		FriendPlanets:  []int{int(aV1.AstroConstants_Sun), int(aV1.AstroConstants_Venus)},
		EnemyPlanets:   []int{int(aV1.AstroConstants_Moon)},
		NeutralPlanets: []int{int(aV1.AstroConstants_Mars), int(aV1.AstroConstants_Jupiter), int(aV1.AstroConstants_Saturn)},
		FriendHouses:   []int{int(aV1.AstroConstants_Leo), int(aV1.AstroConstants_Taurus), int(aV1.AstroConstants_Libra)},
		EnemyHouses:    []int{int(aV1.AstroConstants_Cancer)},
		NeutralHouses: []int{int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio),
			int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces),
			int(aV1.AstroConstants_Capricon), int(aV1.AstroConstants_Aquarius)},
	},
	PLANET_VENUS: &Planet_Meta{
		Name:           "Venus",
		Number:         int(aV1.AstroConstants_Venus),
		Exalted:        int(aV1.AstroConstants_Pisces),
		Deblited:       int(aV1.AstroConstants_Virgo),
		OwnHouse:       []int{int(aV1.AstroConstants_Taurus), int(aV1.AstroConstants_Libra)},
		MoolTrikona:    int(aV1.AstroConstants_Libra),
		FriendPlanets:  []int{int(aV1.AstroConstants_Mercury), int(aV1.AstroConstants_Saturn)},
		EnemyPlanets:   []int{int(aV1.AstroConstants_Sun), int(aV1.AstroConstants_Moon)},
		NeutralPlanets: []int{int(aV1.AstroConstants_Mars), int(aV1.AstroConstants_Jupiter)},
		FriendHouses:   []int{int(aV1.AstroConstants_Gemini), int(aV1.AstroConstants_Virgo), int(aV1.AstroConstants_Capricon), int(aV1.AstroConstants_Aquarius)},
		EnemyHouses: []int{int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio),
			int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces)},
		NeutralHouses: []int{int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio),
			int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces)},
	},
	PLANET_MARS: &Planet_Meta{
		Name:           "Mars",
		Number:         int(aV1.AstroConstants_Mars),
		Exalted:        int(aV1.AstroConstants_Capricon),
		Deblited:       int(aV1.AstroConstants_Cancer),
		OwnHouse:       []int{int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio)},
		Aspects:        []int{4, 8},
		MoolTrikona:    int(aV1.AstroConstants_Aries),
		FriendPlanets:  []int{int(aV1.AstroConstants_Sun), int(aV1.AstroConstants_Moon), int(aV1.AstroConstants_Jupiter)},
		EnemyPlanets:   []int{int(aV1.AstroConstants_Mercury)},
		NeutralPlanets: []int{int(aV1.AstroConstants_Venus), int(aV1.AstroConstants_Saturn)},
		FriendHouses:   []int{int(aV1.AstroConstants_Leo), int(aV1.AstroConstants_Cancer), int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces)},
		EnemyHouses:    []int{int(aV1.AstroConstants_Gemini), int(aV1.AstroConstants_Virgo)},
		NeutralHouses: []int{int(aV1.AstroConstants_Taurus), int(aV1.AstroConstants_Libra),
			int(aV1.AstroConstants_Capricon), int(aV1.AstroConstants_Aquarius)},
	},
	PLANET_JUPITER: &Planet_Meta{
		Name:           "Jupiter",
		Number:         int(aV1.AstroConstants_Jupiter),
		Exalted:        int(aV1.AstroConstants_Cancer),
		Deblited:       int(aV1.AstroConstants_Capricon),
		OwnHouse:       []int{int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces)},
		Aspects:        []int{5, 9},
		MoolTrikona:    int(aV1.AstroConstants_Saggitarius),
		FriendPlanets:  []int{int(aV1.AstroConstants_Sun), int(aV1.AstroConstants_Moon), int(aV1.AstroConstants_Mars)},
		EnemyPlanets:   []int{int(aV1.AstroConstants_Mercury), int(aV1.AstroConstants_Venus)},
		NeutralPlanets: []int{int(aV1.AstroConstants_Saturn)},
		FriendHouses:   []int{int(aV1.AstroConstants_Leo), int(aV1.AstroConstants_Cancer), int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio)},
		EnemyHouses:    []int{int(aV1.AstroConstants_Gemini), int(aV1.AstroConstants_Virgo), int(aV1.AstroConstants_Taurus), int(aV1.AstroConstants_Libra)},
		NeutralHouses:  []int{int(aV1.AstroConstants_Capricon), int(aV1.AstroConstants_Aquarius)},
	},
	PLANET_SATURN: &Planet_Meta{
		Name:           "Saturn",
		Number:         int(aV1.AstroConstants_Saturn),
		Exalted:        int(aV1.AstroConstants_Libra),
		Deblited:       int(aV1.AstroConstants_Aries),
		OwnHouse:       []int{int(aV1.AstroConstants_Capricon), int(aV1.AstroConstants_Aquarius)},
		Aspects:        []int{3, 10},
		MoolTrikona:    int(aV1.AstroConstants_Aquarius),
		FriendPlanets:  []int{int(aV1.AstroConstants_Mercury), int(aV1.AstroConstants_Venus)},
		EnemyPlanets:   []int{int(aV1.AstroConstants_Sun), int(aV1.AstroConstants_Moon), int(aV1.AstroConstants_Mars)},
		NeutralPlanets: []int{int(aV1.AstroConstants_Jupiter)},
		FriendHouses:   []int{int(aV1.AstroConstants_Gemini), int(aV1.AstroConstants_Virgo), int(aV1.AstroConstants_Taurus), int(aV1.AstroConstants_Libra)},
		EnemyHouses:    []int{int(aV1.AstroConstants_Leo), int(aV1.AstroConstants_Cancer), int(aV1.AstroConstants_Aries), int(aV1.AstroConstants_Scorpio)},
		NeutralHouses:  []int{int(aV1.AstroConstants_Saggitarius), int(aV1.AstroConstants_Pisces)},
	},
}

func CalculatePlanetMeta(planet *aV1.PlanetDetails) error {

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

	// find ownwer of house

	// logrus.Infof(" Planet house/owner %v %v", planet.Name, planet.Rasi)
	if slices.Index(planetMeta.FriendHouses, int(planet.Rasi)) != -1 {
		planet.FriendHouse = true
	} else if slices.Index(planetMeta.EnemyHouses, int(planet.Rasi)) != -1 {
		planet.EnemyHouse = true
	} else if slices.Index(planetMeta.NeutralHouses, int(planet.Rasi)) != -1 {
		planet.NeutralHouse = true
	}

	return nil
}
