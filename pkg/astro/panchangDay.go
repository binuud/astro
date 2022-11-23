package astro

import (
	"time"

	astroV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
)

type PanchangDay struct {
	*astroV1.PanchangDay
}

func ProcessPanchangDay(day *astroV1.PanchangDay) {
	p := &PanchangDay{day}
	p.TagPanchangDay()
}

func (p *PanchangDay) TagPanchangDay() {

	zone, _ := time.LoadLocation(p.Zone)
	sunRiseTime := time.Date(int(p.Year), time.Month(p.Month), int(p.Day), int(6), int(0), int(0), 0, zone)
	sunSetTime := time.Date(int(p.Year), time.Month(p.Month), int(p.Day), int(18), int(0), int(0), 0, zone)

	for _, star := range p.Stars {
		if star.Name == int32(astroV1.AstroConstants_Krittika) &&
			star.Transit.AsTime().Before(sunSetTime) {
			p.AddTag(astroV1.AstroMetas_Krithika)
			break
		}
		//  else if star.Name == int32(astroV1.AstroConstants_Rohini) &&
		// 	star.Transit.AsTime().After(sunRiseTime) {
		// 	p.AddTag(astroV1.AstroMetas_Krithika)
		// 	break
		// }
	}

	for _, tithi := range p.Tithis {

		if tithi.Name <= int32(astroV1.AstroConstants_Purnima) {
			p.AddTag(astroV1.AstroMetas_ShuklaPaksha)
		} else if tithi.Name > int32(astroV1.AstroConstants_Purnima) {
			p.AddTag(astroV1.AstroMetas_KrishnaPaksha)
		}

		if tithi.Name == int32(astroV1.AstroConstants_S_Shasti) &&
			tithi.Transit.AsTime().Before(sunRiseTime) {
			p.AddTag(astroV1.AstroMetas_Shasti)
		} else if tithi.Name == int32(astroV1.AstroConstants_S_Saptami) &&
			tithi.Transit.AsTime().After(sunRiseTime) {
			p.AddTag(astroV1.AstroMetas_Shasti)
		}

		if tithi.Name == int32(astroV1.AstroConstants_Amavasya) &&
			tithi.Transit.AsTime().Before(sunSetTime) {
			p.AddTag(astroV1.AstroMetas_Amavasya)
		} else if tithi.Name == int32(astroV1.AstroConstants_S_Prathima) &&
			tithi.Transit.AsTime().After(sunRiseTime) {
			p.AddTag(astroV1.AstroMetas_Amavasya)
		}

		if tithi.Name == int32(astroV1.AstroConstants_Purnima) &&
			tithi.Transit.AsTime().Before(sunSetTime) {
			p.AddTag(astroV1.AstroMetas_Poornima)
		}
		// else if tithi.Name == int32(astroV1.AstroConstants_K_Prathima) &&
		// 	tithi.Transit.AsTime().After(sunRiseTime) {
		// 	p.AddTag(astroV1.AstroMetas_Poornima)
		// }

	}

	// tithiName := 0
	// tithiTime := time.Now()
	// if len(p.Tithis) > 0 {
	// 	tithiName = int(p.Tithis[0].Name)
	// 	tithiTime = p.Tithis[0].Transit.AsTime()
	// }

	// logrus.Infof(" Tagged %d (%v) (%v -- %v) %v", p.Day, tithiName, tithiTime.In(zone), sunRiseTime.In(zone), p.Tags)

}

func (p *PanchangDay) AddTag(tag astroV1.AstroMetas_PanchangTags) {
	if !p.IsTagged(tag) {
		p.Tags = append(p.Tags, tag)
	}
}

func (p *PanchangDay) IsTagged(tag astroV1.AstroMetas_PanchangTags) bool {
	for _, tagged := range p.Tags {
		if tag == tagged {
			return true
		}
	}
	return false
}
