package astro

func CalculateRasi(degree float32) int32 {
	return 1 + (int32)(degree/RASI_DEGREES)
}

func CalculateStar(degree float32) int32 {
	return 1 + (int32)(degree/STAR_DEGREES)
}

func CalculatePaada(degree float32) int32 {

	//starLongitude := math.Mod(float64(degree), STAR_DEGREES)
	paada := int32(degree / PAADA_DEGREES)

	return 1 + paada%4

}

func CalculateNthHarmonic(degree float32, harmonic int32) int32 {

	upScaledDegree := float32(harmonic) * degree

	rasi := upScaledDegree / 30

	return 1 + int32(rasi)%TOTAL_RASIS

}
