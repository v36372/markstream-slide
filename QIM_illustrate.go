func QIMEncode(mag float64, phs float64, bit int) float64 {
	step := [5]float64{PI / 18, PI / 14, PI / 10, PI / 6, PI / 2}
	var stepsize = findStep(mag)
	if bit == 48 {
		return math.Floor(phs/step[stepsize]+0.5) * step[stepsize]
	} else {
		return math.Floor(phs/step[stepsize])*step[stepsize] + step[stepsize]/2
	}
}
