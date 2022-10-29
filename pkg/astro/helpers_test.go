package astro

import "testing"

func TestGetTransitTime(t *testing.T) {

	got, err := GetTransitTime(231.3073792, 238.8580745, 720, 239.994)
	t.Logf("Got result %v", got)
	if err == nil {
		t.Errorf("Expected error, but got no error")
	}

	got, err = GetTransitTime(238.8580745, 246.4365072, 720, 239.994)
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}
	t.Logf("Got result %v", got)

}
