package astro

func CalculateRasi(degree float32) int32 {
	return 1 + (int32)(degree/RASI_DEGREES)
}
