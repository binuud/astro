package astro

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_Pacnhang_Month(t *testing.T) {
	p, err := NewPanchang(2026, 12, 1)
	if err != nil {
		t.Errorf("Cannot prepare chart err(%v)", err)
	}
	day := p.PanchangMonth.Days
	got := STAR_NAMES[day[0].Stars[0].Name]
	logrus.Infof("Stars %v", got)
	if got == "Makha" {
		t.Logf("Panchang working fine")
	} else {
		t.Error("Panchang returning wrong data")
	}
	// p.PrintMonth()
}
