package pumpcalc

func GetVolumeFlowKg(massFlowKgS float64, density float64) (float64, error) {
	if density <= 0 {
		return 0, DivideByZeroError
	}
	if massFlowKgS < 0 {
		return 0, NegativeMassFlowKgSError
	}
	return massFlowKgS / density, nil
}
