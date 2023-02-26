package citiesServer

import "testing"

func TestNewCityDB(t *testing.T) {

	got := NewCityDB()
	t.Logf("Got result (%v) index 1 ", got.cities[1])
	t.Logf("Got result (%v) index 0 ", got.cities[0])
	if len(got.cities) != 150710 {
		t.Errorf("Expected  150710 entries but got '%v'", len(got.cities))
	}

}

func TestGetCityInfo(t *testing.T) {

	c := NewCityDB()
	got, err := c.GetCityInfo(-1)
	if err == nil {
		t.Errorf("Expected error for out of range id - 1 '%v'", got)
	}
	got, err = c.GetCityInfo(150711)
	if err == nil {
		t.Errorf("Expected error for out of range id  150711'%v'", got)
	}

	got, err = c.GetCityInfo(0)
	if err != nil {
		t.Errorf("Expected data but got error '%v'", err)
	} else if got.Id != 52 {
		t.Errorf("Expected - 52,Ashkāsham,3901,BDS,Badakhshan,1,AF,Afghanistan,36.68333000,71.53333000,Q4805192 -  but got '%v'", got)
	} else {
		t.Logf("Got expected result (%v)", got)
	}

	got, err = c.GetCityInfo(26)
	if err != nil {
		t.Errorf("Expected data but got error '%v'", err)
	} else if got.Id != 121 {
		t.Errorf("Expected - 121,Shahrak,3888,GHO,Ghōr,1,AF,Afghanistan,34.10737000,64.30520000,Q7461922 -  but got '%v'", got)
	} else {
		t.Logf("Got expected result (%v)", got)
	}

}

func TestSearchCity(t *testing.T) {

	c := NewCityDB()
	got, err := c.SearchCity("Che", "I")
	if err == nil {
		t.Errorf("Expected error for out of range Country ID '%v'", got)
	}
	got, err = c.SearchCity("Che", "INA")
	if err == nil {
		t.Errorf("Expected error for out of range Country ID '%v'", got)
	}

	got, err = c.SearchCity("Chennai", "IN")
	if err != nil {
		t.Errorf("Expected data but got error '%v'", err)
	} else if len(got) != 1 {
		t.Errorf("Expected exact one result '%v'", got)
	} else if got[0].Id != 131517 {
		t.Errorf("Expected - 131517,Chennai,4035,TN,\"Tamil Nadu\",101,IN,India,13.08784000,80.27847000,Q1352 - but got  '%v'", got)
	}

	got, err = c.SearchCity("Chenn", "IN")
	if err != nil {
		t.Errorf("Expected data but got error '%v'", err)
	} else if len(got) != 2 {
		t.Errorf("Expected exact two result '%v'", got)
	}

	got, err = c.SearchCity("Che", "IN")
	if err != nil {
		t.Errorf("Expected data but got error '%v'", err)
	} else if len(got) > 10 {
		t.Errorf("Expected less than 10 results, but got %v results ", len(got))
	}

}
