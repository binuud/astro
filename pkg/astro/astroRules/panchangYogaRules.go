package astroRules

import (
	"fmt"

	aV1 "github.com/dronasys-com/astro/proto/v1/astro/go"
)

const (
	TOTAL_STARS_WITH_ABHIJIT = 28
)

func init() {

}

func GetVaarStarYogaNum(weekDay int32, starNum int32) (int32, error) {

	if weekDay < 0 || weekDay > 6 {
		err := fmt.Errorf("wrong weekday (range 0-6)")
		return 0, err
	}

	if starNum < 1 || starNum > 28 {
		err := fmt.Errorf("wrong star number (range 1-28)")
		return 0, err
	}

	yogaNum := starNum - (4 * int32(weekDay))

	if yogaNum < 1 {
		yogaNum = TOTAL_STARS_WITH_ABHIJIT + yogaNum
	}
	return yogaNum, nil

}

func GetVaarStarYoga(weekDay int32, starNum int32) (*aV1.StarWeekdayYoga, error) {
	yogaNum, err := GetVaarStarYogaNum(weekDay, starNum)
	if err != nil {
		return nil, err
	}
	return VaarStarYogas[yogaNum-1], nil
}
