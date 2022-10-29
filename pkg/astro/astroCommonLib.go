package astro

func CalculateRasi(degree float64) int64 {
	return 1 + (int64)(degree/RASI_DEGREES)
}

func CalculateStar(degree float64) int64 {
	return 1 + (int64)(degree/STAR_DEGREES)
}

func CalculatePaada(degree float64) int64 {

	//starLongitude := math.Mod(float64(degree), STAR_DEGREES)
	paada := int64(degree / PAADA_DEGREES)

	return 1 + paada%4

}

func CalculateNthHarmonic(degree float64, harmonic int64) int64 {

	upScaledDegree := float64(harmonic) * degree

	rasi := upScaledDegree / 30

	return 1 + int64(rasi)%TOTAL_RASIS

}
