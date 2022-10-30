package astro

import (
	"fmt"
	"time"
)

type CalendarPrinter struct {
	Panchang      *Panchang
	StarCounter   int
	TithiCounter  int
	YogaCounter   int
	KaranaCounter int
	Zone          *time.Location
}

func NewCalendarPrinter(panchang *Panchang) *CalendarPrinter {
	cal := &CalendarPrinter{
		Panchang:      panchang,
		StarCounter:   0,
		TithiCounter:  0,
		YogaCounter:   0,
		KaranaCounter: 0,
	}
	cal.Zone, _ = time.LoadLocation("Asia/Kolkata")
	return cal
}

func (cal *CalendarPrinter) PrintCalendar() {

	numDays := 366
	dayCounter := 1

	t := time.Date(int(cal.Panchang.Year), time.Month(1), int(1), int(0), int(0), int(0), 0, cal.Zone)

	for dayCounter <= numDays {

		hoursElapsed := 24

		fmt.Printf("**************************************\n")
		fmt.Printf("Year %d, Month %d, Day %d (%s) \n", t.In(cal.Zone).Year(), t.In(cal.Zone).Month(), t.In(cal.Zone).Day(), t.In(cal.Zone).Weekday())

		dayEnd := t.Add(time.Second * time.Duration(86399))

		// now display the star transits for the day
		cal.printStars(t, dayEnd)
		cal.printTithis(t, dayEnd)
		cal.printYogas(t, dayEnd)
		cal.printKaranas(t, dayEnd)

		t = t.Add(time.Hour * time.Duration(hoursElapsed))
		dayCounter++

	}

}

func (cal *CalendarPrinter) printStars(dayStart time.Time, dayEnd time.Time) {

	safeCounter := 0
	if cal.StarCounter == 0 {
		fmt.Printf(" Star %s \n", STAR_NAMES[cal.Panchang.Stars[cal.StarCounter].Name])
	} else {
		fmt.Printf(" Star %s \n", STAR_NAMES[cal.Panchang.Stars[cal.StarCounter-1].Name])
	}

	for {
		if cal.Panchang.Stars[cal.StarCounter].Transit.Before(dayStart) {
			cal.StarCounter++
		} else {
			if cal.Panchang.Stars[cal.StarCounter].Transit.Before(dayEnd) {
				fmt.Printf(" Star %s @ %v \n", STAR_NAMES[cal.Panchang.Stars[cal.StarCounter].Name], cal.Panchang.Stars[cal.StarCounter].Transit.In(cal.Zone))
				cal.StarCounter++
			}
		}
		safeCounter++
		if safeCounter > 5 {
			break
		}
	}

}

func (cal *CalendarPrinter) printTithis(dayStart time.Time, dayEnd time.Time) {

	safeCounter := 0
	if cal.TithiCounter == 0 {
		fmt.Printf(" Tithi %s \n", TITHI_NAMES[cal.Panchang.Tithis[cal.TithiCounter].Name])
	} else {
		fmt.Printf(" Tithi %s \n", TITHI_NAMES[cal.Panchang.Tithis[cal.TithiCounter-1].Name])
	}

	for {
		if cal.Panchang.Tithis[cal.TithiCounter].Transit.Before(dayStart) {
			cal.TithiCounter++
		} else {
			if cal.Panchang.Tithis[cal.TithiCounter].Transit.Before(dayEnd) {
				fmt.Printf(" Tithi %s @ %v \n", TITHI_NAMES[cal.Panchang.Tithis[cal.TithiCounter].Name], cal.Panchang.Tithis[cal.TithiCounter].Transit.In(cal.Zone))
				cal.TithiCounter++
			}
		}
		safeCounter++
		if safeCounter > 5 {
			break
		}
	}

}

func (cal *CalendarPrinter) printYogas(dayStart time.Time, dayEnd time.Time) {

	safeCounter := 0
	if cal.YogaCounter == 0 {
		fmt.Printf(" Yoga %s \n", YOGA_NAMES[cal.Panchang.Yogas[cal.YogaCounter].Name])
	} else {
		fmt.Printf(" Yoga %s \n", YOGA_NAMES[cal.Panchang.Yogas[cal.YogaCounter-1].Name])
	}

	for {
		if cal.Panchang.Yogas[cal.YogaCounter].Transit.Before(dayStart) {
			cal.YogaCounter++
		} else {
			if cal.Panchang.Yogas[cal.YogaCounter].Transit.Before(dayEnd) {
				fmt.Printf(" Yoga %s @ %v \n", YOGA_NAMES[cal.Panchang.Yogas[cal.YogaCounter].Name], cal.Panchang.Yogas[cal.YogaCounter].Transit.In(cal.Zone))
				cal.YogaCounter++
			}
		}
		safeCounter++
		if safeCounter > 5 {
			break
		}
	}

}

func (cal *CalendarPrinter) printKaranas(dayStart time.Time, dayEnd time.Time) {

	safeCounter := 0
	if cal.KaranaCounter == 0 {
		fmt.Printf(" Karana %s \n", KARANA_NAMES[cal.Panchang.Karanas[cal.KaranaCounter].Name])
	} else {
		fmt.Printf(" Karana %s \n", KARANA_NAMES[cal.Panchang.Karanas[cal.KaranaCounter-1].Name])
	}

	for {
		if cal.Panchang.Karanas[cal.KaranaCounter].Transit.Before(dayStart) {
			cal.KaranaCounter++
		} else {
			if cal.Panchang.Karanas[cal.KaranaCounter].Transit.Before(dayEnd) {
				fmt.Printf(" Karana %s @ %v \n", KARANA_NAMES[cal.Panchang.Karanas[cal.KaranaCounter].Name], cal.Panchang.Karanas[cal.KaranaCounter].Transit.In(cal.Zone))
				cal.KaranaCounter++
			}
		}
		safeCounter++
		if safeCounter > 5 {
			break
		}
	}

}
