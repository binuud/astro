package astro

// GetStarDegreesArray
// returns an array of star degree boundaries (float64)
// Since the star degrees will be used frequently, we populate the array
// instead of making lot of float divisions and multiplications
func GetStarDegreesArray() []float64 {
	starDegree := 0.0
	starCount := 0
	starDegrees := make([]float64, TOTAL_STARS)
	for starCount < TOTAL_STARS {
		starDegrees = append(starDegrees, starDegree)
		starDegree += STAR_DEGREES
		starCount++
	}
	return starDegrees
}
