package mcpServer_test

import (
	"testing"

	"github.com/dronasys-com/astro/pkg/servers/mcpServer"
)

func Test_GetAge(t *testing.T) {
	t.Log("Testing astro age")
	_, got, err := mcpServer.GetAge(t.Context(), nil, mcpServer.Input_GetAge{DOB: "01/01/1950"})
	if err != nil {
		t.Errorf("error Test_GetAge %v", err)
	}
	if got.Age != "" {
		t.Logf("got age %s", got.Age)
	}
}

func Test_GetLatLong(t *testing.T) {
	t.Log("Testing Get lang long of a city")
	_, got, err := mcpServer.GetLatLong(t.Context(), nil, mcpServer.Input_GetLatLong{City: "Chennai", CountryCode: "IN"})
	if err != nil {
		t.Errorf("error Test_GetAge %v", err)
	}
	if got.Lat != "" {
		t.Logf("got city: %s, lat: %s long: %s", got.City, got.Lat, got.Long)
	}
}

func Test_GetAstrologicalChart(t *testing.T) {
	t.Log("Testing astro age")
	_, got, err := mcpServer.GetAstrologicalChart(t.Context(), nil, mcpServer.Input_GetAstrologicalChart{DOB: "01/01/1950", TOB: "10:10", City: "Chennai", CountryCode: "IN"})
	if err != nil {
		t.Errorf("error Test_GetAge %v", err)
	}
	if got.Text != "" {
		t.Logf("got text %s", got.Text)
	}
}
