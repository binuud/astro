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
	got, err = GetVaarStarYogaNum(1, 1)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == 25 {
		t.Logf("Got Expected result (Monday/Aswini) 25 (%v)", got)
	} else {
		t.Errorf("Expected  (Monday/Aswini) 25, but got '%v'", got)
	}

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

	got, err = GetVaarStarYogaNum(4, 1)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if got == 13 {
		t.Logf("Got Expected result (Thursday/Aswhwini) 13 (%v)", got)
	} else {
		t.Errorf("Expected (Thursday/Aswhwini) 13, but got '%v'", got)
	}

}

func TestGetVaarStarYoga(t *testing.T) {

	// invalid test cases
	got, err := GetVaarStarYoga(-1, 24)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	got, err = GetVaarStarYoga(1, 29)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	// valid test cases
	got, err = GetVaarStarYoga(0, 10) // Mudgara
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if int32(got.Nature) == int32(aV1.AstroMetas_Malefic) && got.Yoga == aV1.StarWeekdayYoga_Mudgara {
		t.Logf("Got Expected result (Mudgara) (%v)", got)
	} else {
		t.Errorf("Expected (Mudgara) but got '%v'", got)
	}

	got, err = GetVaarStarYoga(6, 24) // Mudgara
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if int32(got.Nature) == int32(aV1.AstroMetas_Benefic) && got.Yoga == aV1.StarWeekdayYoga_Pravardhamama {
		t.Logf("Got Expected result (Pravardhamama) (%v)", got)
	} else {
		t.Errorf("Expected (Pravardhamama) but got '%v'", got)
	}

	got, err = GetVaarStarYoga(4, 1) // Manas
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	if int32(got.Nature) == int32(aV1.AstroMetas_Benefic) && got.Yoga == aV1.StarWeekdayYoga_Manas {
		t.Logf("Got Expected result (Manas) (%v)", got)
	} else {
		t.Errorf("Expected (Manas) but got '%v'", got)
	}

}
