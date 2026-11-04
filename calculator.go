package pumpcalc

func GetVolumeFlowKg(massFlowKg float64, density float64) (float64, error) {
	if density <= 0 {
		return 0, DivideByZeroError
	}
	return massFlowKg / density, nil
}
