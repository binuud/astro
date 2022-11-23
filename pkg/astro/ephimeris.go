package astro

import (
	"fmt"
	"os/exec"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
	"github.com/sirupsen/logrus"
)

const (
	AYANAMSA_LAHIRI        = 1
	EPHIMERIS_SWISS        = "eswe"
	ARG_PLANET_ALL         = "d"
	ARG_PLANET_MOON        = "1"
	ARG_PLANET_SUN_MOON    = "01"
	FSEQALL                = "PLBRS"
	FLONGITUDE             = "Pl"
	FLONGITUDEONLY         = "l"
	PLANET_DEFAULT_COLUMNS = "Plbrds"
	ASTRO_CMD              = "/go/src/github.com/dronasys-com/astro/services/astroCalc/cmd/astroBin/astrobin"
)

type Ephimeris struct {
	BirthData *astroV1.BirthData
	CmdArgs   []string
}

func NewEphimeris(birthData *astroV1.BirthData) *Ephimeris {

	logrus.Infof(" Birthdata (%v)", birthData)

	return &Ephimeris{
		BirthData: birthData,
		CmdArgs:   make([]string, 0),
	}
}

func (e *Ephimeris) GetPlanetsEphimeris() (string, error) {

	e.NewArgs().SetAyanamsa(AYANAMSA_LAHIRI).
		SetDTPoB().SetCount(1).SetSteps("1").SetPlanets(ARG_PLANET_ALL).
		SetEphimerisType(EPHIMERIS_SWISS).SetHouse().SetColumns(PLANET_DEFAULT_COLUMNS)
	return e.ExecCommand()

}

func (e *Ephimeris) GetMoonEphimeris() (string, error) {

	e.NewArgs().SetAyanamsa(AYANAMSA_LAHIRI).
		SetDTPoB().SetCount(3).SetSteps("0.5").SetPlanets(ARG_PLANET_MOON).
		SetEphimerisType(EPHIMERIS_SWISS).SetColumns(FLONGITUDEONLY).NoHead()
	return e.ExecCommand()

}

func (e *Ephimeris) GetSunMoonEphimerisForYear() (string, error) {

	e.NewArgs().SetAyanamsa(AYANAMSA_LAHIRI).
		SetDTPoB().SetCount(1472).SetSteps("0.25").SetPlanets(ARG_PLANET_SUN_MOON).
		SetEphimerisType(EPHIMERIS_SWISS).SetColumns(FLONGITUDEONLY).NoHead().Horizontal()
	return e.ExecCommand()

}

func (e *Ephimeris) GetSunMoonEphimerisForMonth() (string, error) {

	numDays := 34
	count := numDays * 24
	e.NewArgs().SetAyanamsa(AYANAMSA_LAHIRI).
		SetDTPoB().SetCount(count).SetSteps("60m").SetPlanets(ARG_PLANET_SUN_MOON).
		SetEphimerisType(EPHIMERIS_SWISS).SetColumns(FLONGITUDEONLY).NoHead().Horizontal()
	return e.ExecCommand()

}

func (e *Ephimeris) ExecCommand() (string, error) {

	logrus.Infof("Executing command %s %v", ASTRO_CMD, e.CmdArgs)
	out, err := exec.Command(ASTRO_CMD, e.CmdArgs...).Output()

	//out, err := exec.Command("./astroBin/astrobin", dateString, timeString, posString, "-n1", "-s1", "-sid1", "-fPl", "-pd", "-eswe", "-house").Output()
	// ./astrobin -b4.2.1978 -ut05:15:00 -geopos80.27,13.08,0 -n1 -s1 -sid1 -fPl -pd -eswe -house
	//out, err := exec.Command("cat", "astroBin/sample.txt").Output()
	//  " -b$utdatenow -ut$utnow  -geopos$pob_long,$pob_lat,0 -n$num -s$steps -sid1 -fl -p1 -head";

	if err != nil {
		logrus.Errorf("Cannot get ephimeris data (%v)", err)
		return "", err
	}
	// logrus.Info("Output %v", string(out))
	return string(out), err

}

func (e *Ephimeris) NewArgs() *Ephimeris {
	e.CmdArgs = make([]string, 0)
	return e
}

func (e *Ephimeris) SetDTPoB() *Ephimeris {

	tob := e.BirthData.Tob.Tobtime.AsTime()
	dateString := fmt.Sprintf("-b%d.%d.%d", tob.Day(), tob.Month(), tob.Year())
	timeString := fmt.Sprintf("-ut%02d:%02d:%02d", tob.Hour(), tob.Minute(), tob.Second())

	e.CmdArgs = append(e.CmdArgs, dateString)
	e.CmdArgs = append(e.CmdArgs, timeString)

	if e.BirthData.Pob != nil {
		posString := fmt.Sprintf("-geopos%f,%f,0", e.BirthData.Pob.Long, e.BirthData.Pob.Lat)
		e.CmdArgs = append(e.CmdArgs, posString)
	}
	return e
}

func (e *Ephimeris) SetCount(count int) *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, fmt.Sprintf("-n%d", count))
	return e

}

func (e *Ephimeris) SetSteps(steps string) *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, fmt.Sprintf("-s%s", steps))
	return e

}

func (e *Ephimeris) SetAyanamsa(ayanamsa int) *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, fmt.Sprintf("-sid%d", ayanamsa))
	return e

}

func (e *Ephimeris) SetEphimerisType(ephimeris string) *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, fmt.Sprintf("-%s", ephimeris))
	return e

}

func (e *Ephimeris) SetPlanets(planets string) *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, fmt.Sprintf("-p%s", planets))
	return e

}

func (e *Ephimeris) SetColumns(seq string) *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, fmt.Sprintf("-f%s", seq))
	return e

}

func (e *Ephimeris) SetHouse() *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, "-house")
	return e

}

func (e *Ephimeris) NoHead() *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, "-head")
	return e

}

func (e *Ephimeris) Horizontal() *Ephimeris {

	e.CmdArgs = append(e.CmdArgs, "-hor")
	return e

}
