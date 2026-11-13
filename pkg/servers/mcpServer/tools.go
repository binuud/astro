package mcpServer

import (
	"context"
	"fmt"

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

type Input_GetAstrologicalChart struct {
	DOB string `json:"dob" jsonschema:"The Date of birth in DD/MM/YYYY"`
	TOB string `json:"tob" jsonschema:"The Time of birth in HH:MM"`
}

type Output_GetAstrologicalChart struct {
	Text string `json:"text" jsonschema:"The astrological position of planets in the zodiac"`
}

func GetAstrologicalChart(ctx context.Context, req *mcp.CallToolRequest, input Input_GetAstrologicalChart) (
	*mcp.CallToolResult,
	Output_GetAstrologicalChart,
	error,
) {
	logrus.Infof("GetAstrologicalChart called with %v", input)
	dob, err := parseDOB(input.DOB)
	if err != nil {
		return nil, Output_GetAstrologicalChart{Text: "error Cannot parse date of birth, please enter DOB in DD/MM/YYYY format"}, nil
	}

	tob, err := parseTOB(input.TOB)
	if err != nil {
		return nil, Output_GetAstrologicalChart{Text: "error Cannot parse time of birth, please enter TOB in HH:MM format"}, nil
	}

	logrus.Infof("Creating chart for DOB %d/%d/%d :: TOB %d:%d", dob.Day(), dob.Month(), dob.Year(), tob.Hour(), tob.Minute())

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
			City: "Chennai",
			Lat:  13.08,
			Long: 80.7,
		},
	})
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	out := chart.GetChartDescription()
	return nil, Output_GetAstrologicalChart{Text: fmt.Sprintf("For your given date and time of birth, the position of planets in zodiac is %s", out)}, nil

}
