package mcpServer

import (
	"context"
	"fmt"
	"strconv"

	astroV1 "github.com/dronasys-com/astro/gen/go/v1/astro"
	"github.com/dronasys-com/astro/pkg/astro"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sirupsen/logrus"
)

type Input_GetAge struct {
	DOB string `json:"dob" jsonschema:"The Date of birth in DD/MM/YYYY"`
}

type Output_GetAge struct {
	Age string `json:"age" jsonschema:"The age of the user"`
}

func GetAge(ctx context.Context, req *mcp.CallToolRequest, input Input_GetAge) (
	*mcp.CallToolResult,
	Output_GetAge,
	error,
) {
	age, err := CalculateAge(input.DOB)
	if err != nil {
		return nil, Output_GetAge{Age: "Cannot calculate age, please enter DOB in DD/MM/YYYY format"}, nil
	}
	return nil, Output_GetAge{Age: fmt.Sprintf("Your age is %d", age)}, nil
}

type Input_GetLatLong struct {
	City        string `json:"city" jsonschema:"The City for which latitude, longitude has to be found"`
	CountryCode string `json:"countryCode" jsonschema:"The Country Code, where the city is located. Eg: IN, US"`
}

type Output_GetLatLong struct {
	City string `json:"city" jsonschema:"The closest matching city"`
	Lat  string `json:"lat" jsonschema:"The latitude of the city"`
	Long string `json:"long" jsonschema:"The longitude of the city"`
}

func GetLatLong(ctx context.Context, req *mcp.CallToolRequest, input Input_GetLatLong) (
	*mcp.CallToolResult,
	*Output_GetLatLong,
	error,
) {
	city, lat, long, err := GetLatLongForCity(input.City, input.CountryCode)
	if err != nil {
		return nil, nil, err
	}
	return nil, &Output_GetLatLong{City: city, Lat: strconv.FormatFloat(lat, 'f', 2, 64), Long: strconv.FormatFloat(long, 'f', 2, 64)}, nil
}

type Input_GetAstrologicalChart struct {
	DOB         string `json:"dob" jsonschema:"The Date of birth in DD/MM/YYYY"`
	TOB         string `json:"tob" jsonschema:"The Time of birth in HH:MM"`
	City        string `json:"city" jsonschema:"The City of birth"`
	CountryCode string `json:"countryCode" jsonschema:"The Country Code, for the country of birth. Eg: IN, US"`
}

//	type Output_GetAstrologicalChart struct {
//		Text string `json:"text" jsonschema:"The astrological position of planets in the zodiac"`
//	}
type Output_GetAstrologicalChart struct {
	Zodiac map[string]string `json:"zodiac" jsonschema:"The zodiac where each planet is placed"`
	House  map[string]string `json:"house" jsonschema:"The house number where each planet is placed"`
	Text   string            `json:"text" jsonschema:"The astrological position of planets in the zodiac"`
}

func GetAstrologicalChart(ctx context.Context, req *mcp.CallToolRequest, input Input_GetAstrologicalChart) (
	*mcp.CallToolResult,
	*Output_GetAstrologicalChart,
	error,
) {
	logrus.Infof("GetAstrologicalChart called with %v", input)

	dob, err := parseDOB(input.DOB)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "error Cannot parse date of birth, please enter DOB in DD/MM/YYYY format"},
			},
		}, nil, fmt.Errorf("error Cannot parse date of birth, please enter DOB in DD/MM/YYYY format")
	}

	tob, err := parseTOB(input.TOB)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "error Cannot parse time of birth, please enter TOB in HH:MM format"},
			},
		}, nil, fmt.Errorf("error Cannot parse time of birth, please enter TOB in HH:MM format")

	}

	city, lat, long, err := GetLatLongForCity(input.City, input.CountryCode)
	if err != nil {
		return nil, nil, err
	}

	logrus.Infof("Creating chart for DOB %d/%d/%d :: TOB %d:%d", dob.Day(), dob.Month(), dob.Year(), tob.Hour(), tob.Minute())
	logrus.Infof("Place of Birth %s, %s [%.2f] [%.2ff]", city, input.CountryCode, lat, long)

	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    int32(dob.Day()),
			M:    int32(dob.Month()),
			Y:    int32(dob.Year()),
			Hour: int32(tob.Hour()),
			Min:  int32(tob.Minute()),
			Sec:  0,
			Zone: "Asia/Kolkata",
		},
		Pob: &astroV1.PobData{
			City: city,
			Lat:  float32(lat),
			Long: float32(long),
		},
	})
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	out := _getAstrologicalChart(chart)
	return nil, out, nil
	//return nil, Output_GetAstrologicalChart{Text: fmt.Sprintf("For your given date and time of birth, the position of planets in zodiac is %s", out)}, nil

}
