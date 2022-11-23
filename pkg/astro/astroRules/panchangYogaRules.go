package astroRules

import (
	"fmt"

	"github.com/dronasys-com/astro/pkg/astro"
	aV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
)

var VaarStarYogaNature = []int32{
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Malefic), //0
	int32(aV1.AstroMetas_Malefic), int32(aV1.AstroMetas_Benefic), //2
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Malefic), //4
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Benefic), //8
	int32(aV1.AstroMetas_Malefic), int32(aV1.AstroMetas_Malefic), //10
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Benefic), //12
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Malefic), //14
	int32(aV1.AstroMetas_Malefic), int32(aV1.AstroMetas_Malefic), //16
	int32(aV1.AstroMetas_Malefic), int32(aV1.AstroMetas_Malefic), //18
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Benefic), //20
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Malefic), //22
	int32(aV1.AstroMetas_Malefic), int32(aV1.AstroMetas_Benefic), //24
	int32(aV1.AstroMetas_Malefic), int32(aV1.AstroMetas_Benefic), //26
	int32(aV1.AstroMetas_Benefic), int32(aV1.AstroMetas_Benefic), //28
}

func init() {

}

func GetVaarStarYogaNum(weekDay aV1.AstroConstants_Weekdays, starNum int32) (int32, error) {

	if weekDay < 0 || weekDay > 6 {
		err := fmt.Errorf("wrong weekday (range 0-6)")
		return 0, err
	}

	if starNum < 1 || starNum > 28 {
		err := fmt.Errorf("wrong star number (range 1-28)")
		return 0, err
	}

	starNum = starNum - (4 * int32(weekDay))

	if starNum < 1 {
		starNum = astro.TOTAL_STARS_WITH_ABHIJIT - starNum
	}
	return starNum, nil

}

func GetVaarStarYogaNature(yogaNum int32) (int32, error) {
	if yogaNum < 1 || yogaNum > 28 {
		err := fmt.Errorf("wrong vaar star yoga number (range 1 to 28)")
		return 0, err
	}
	return VaarStarYogaNature[yogaNum-1], nil
}
