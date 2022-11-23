package astroRules

import (
	"testing"

	aV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
)

func TestGetVaarStarYogaNum(t *testing.T) {

	// invalid test cases
	got, err := GetVaarStarYogaNum(-1, 24)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	got, err = GetVaarStarYogaNum(7, 24)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	got, err = GetVaarStarYogaNum(2, 0)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	got, err = GetVaarStarYogaNum(2, 29) // monday
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	// valid test cases
	got, err = GetVaarStarYogaNum(2, 28)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == 20 {
		t.Logf("Got Expected result (Tuesday/Revthi) 20 (%v)", got)
	} else {
		t.Errorf("Expected (Tuesday/Revthi) 20, but got '%v'", got)
	}

	got, err = GetVaarStarYogaNum(0, 1)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == 1 {
		t.Logf("Got Expected (Sunday/Ashwini) result 1 (%v)", got)
	} else {
		t.Errorf("Expected (Sunday/Ashwini) 1, but got '%v'", got)
	}

	got, err = GetVaarStarYogaNum(6, 24)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == 28 {
		t.Logf("Got Expected result (Saturday/Dhanista) 28 (%v)", got)
	} else {
		t.Errorf("Expected (Saturday/Dhanista) 28, but got '%v'", got)
	}

}

func TestGetVaarStarYogaNature(t *testing.T) {

	// invalid test cases
	got, err := GetVaarStarYogaNature(0)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	got, err = GetVaarStarYogaNature(29)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	// valid test cases
	got, err = GetVaarStarYogaNature(10) // Mudgara
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == int32(aV1.AstroMetas_Malefic) {
		t.Logf("Got Expected result (Mudgara) %d (%v)", int32(aV1.AstroMetas_Malefic), got)
	} else {
		t.Errorf("Expected (Mudgara) %d, but got '%v'", int32(aV1.AstroMetas_Malefic), got)
	}

	got, err = GetVaarStarYogaNature(28) // Pravardhamama
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == int32(aV1.AstroMetas_Benefic) {
		t.Logf("Got Expected result (Pravardhamama) %d (%v)", int32(aV1.AstroMetas_Benefic), got)
	} else {
		t.Errorf("Expected (Pravardhamama) %d, but got '%v'", int32(aV1.AstroMetas_Benefic), got)
	}

	got, err = GetVaarStarYogaNature(22) // Mushala
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == int32(aV1.AstroMetas_Malefic) {
		t.Logf("Got Expected result (Mushala) %d (%v)", int32(aV1.AstroMetas_Malefic), got)
	} else {
		t.Errorf("Expected (Mushala) %d, but got '%v'", int32(aV1.AstroMetas_Malefic), got)
	}

}
