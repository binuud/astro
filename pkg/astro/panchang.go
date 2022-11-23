package astro

import (
	"fmt"
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
)

type Panchang struct {
	Transits      *astroV1.Transits
	PanchangMonth *astroV1.PanchangMonth
	InDate        *astroV1.TobData
}

func NewPanchang(year int32, month int32, day int32) (*Panchang, error) {

	transits, err := NewTransitsMoon(year, month)
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}

	transits.CalculateTransits()
	transits.Transits.Ephimeris = nil

	myPanchang := &Panchang{
		Transits: transits.Transits,
		InDate: &astroV1.TobData{
			D:    day,
			M:    month,
			Y:    year,
			Hour: 0,
			Min:  0,
			Sec:  0,
			Zone: "Asia/Kolkata",
		},
	}

	myPanchang.BuildPanchang()

	return myPanchang, nil
}

func (p *Panchang) BuildPanchang() error {

	var err error
	zone, _ := time.LoadLocation(p.InDate.Zone)

	panchangMonth := &astroV1.PanchangMonth{
		Year:  p.InDate.Y,
		Month: p.InDate.M,
		Zone:  p.InDate.Zone,
		Days:  make([]*astroV1.PanchangDay, 0),
	}

	// create the array for required dates
	p.InDate, err = MakeBirthTob(p.InDate)
	if err != nil {
		return err
	}

	startTime := p.InDate.Tobtime.AsTime()
	dayCount := 1

	curStarIndex := 0
	curTithiIndex := 0
	curYogaIndex := 0
	curKaranasIndex := 0
	curPaadaIndex := 0

	for {

		var stars []*astroV1.Transit
		var tithis []*astroV1.Transit
		var yogas []*astroV1.Transit
		var karanas []*astroV1.Transit
		var paadas []*astroV1.Transit

		endTime := startTime.Add(time.Hour * 24)

		curStarIndex, stars = p.GetTransitingObjects(curStarIndex, startTime, endTime, p.Transits.Stars)
		curTithiIndex, tithis = p.GetTransitingObjects(curTithiIndex, startTime, endTime, p.Transits.Tithis)
		curYogaIndex, yogas = p.GetTransitingObjects(curYogaIndex, startTime, endTime, p.Transits.Yogas)
		curKaranasIndex, karanas = p.GetTransitingObjects(curKaranasIndex, startTime, endTime, p.Transits.Karanas)
		curPaadaIndex, paadas = p.GetTransitingObjects(curPaadaIndex, startTime, endTime, p.Transits.Paadas)

		panchangDay := &astroV1.PanchangDay{
			Year:    int32(startTime.In(zone).Year()),
			Month:   int32(startTime.In(zone).Month()),
			Day:     int32(startTime.In(zone).Day()),
			Weekday: int32(startTime.In(zone).Weekday()),
			Zone:    p.InDate.Zone,
			Stars:   stars,
			Tithis:  tithis,
			Yogas:   yogas,
			Karanas: karanas,
			Paadas:  paadas,
		}

		ProcessPanchangDay(panchangDay)
		panchangMonth.Days = append(panchangMonth.Days, panchangDay)

		dayCount++
		startTime = startTime.Add(time.Hour * 24)

		if int32(startTime.In(zone).Month()) != p.InDate.M {
			break
		}
	}

	p.PanchangMonth = panchangMonth

	return nil

}

func (p *Panchang) PrintMonth() {
	zone, _ := time.LoadLocation(p.InDate.Zone)
	for _, day := range p.PanchangMonth.Days {
		if p.InDate.M != day.Month {
			break
		}
		fmt.Printf("\n\n Day                          %d", day.Day)
		for _, star := range day.Stars {
			fmt.Printf("\n  Star    %2d %25s  %v", star.Name, STAR_NAMES[star.Name], star.Transit.AsTime().In(zone).Format(time.Kitchen))
		}
		for _, tithi := range day.Tithis {
			fmt.Printf("\n  Tithi   %2d %25s  %v", tithi.Name, TITHI_NAMES[tithi.Name], tithi.Transit.AsTime().In(zone).Format(time.Kitchen))
		}
		for _, yoga := range day.Yogas {
			fmt.Printf("\n  Yoga    %2d %25s  %v", yoga.Name, YOGA_NAMES[yoga.Name], yoga.Transit.AsTime().In(zone).Format(time.Kitchen))
		}
		for _, karana := range day.Karanas {
			fmt.Printf("\n  Karana  %2d %25s  %v", karana.Name, KARANA_NAMES[karana.Name], karana.Transit.AsTime().In(zone).Format(time.Kitchen))
		}
		for _, paada := range day.Paadas {
			fmt.Printf("\n  Paada   %2d %25s  %v", paada.Name%9+1, RASI_NAMES[paada.Name/9], paada.Transit.AsTime().In(zone).Format(time.Kitchen))
		}
		fmt.Printf("\n  Tags   %v\n\n", day.Tags)
		fmt.Println()
	}
}

func (p *Panchang) GetTransitingObjects(curIndex int, startTime time.Time, endTime time.Time, transits []*astroV1.Transit) (int, []*astroV1.Transit) {

	// zone, _ := time.LoadLocation(p.InDate.Zone)
	newTransits := make([]*astroV1.Transit, 0)
	safeCounter := 0
	for {

		safeCounter++
		if safeCounter > 24 {
			break
		}

		if curIndex >= len(transits) {
			break
		}

		// fmt.Printf(" Checking item %d @ %v  %v <== ==> %v\n", transits[curIndex].Name, transits[curIndex].Transit.AsTime().In(zone), startTime.In(zone), endTime.In(zone))

		if transits[curIndex].Transit.AsTime().After(startTime) {
			if transits[curIndex].Transit.AsTime().Before(endTime) {
				// fmt.Printf(" item %s @ %v \n", transits[curIndex].Name, transits[curIndex].Transit.AsTime().In(zone))
				newTransits = append(newTransits, transits[curIndex])
			} else {
				break
			}
		}
		curIndex++

	}
	return curIndex, newTransits
}
