package pumpcalc

import "errors"

func GetVolumeFlowKg(massFlowKg float64, density float64) (float64, error) {
	if density <= 0 {
		return 0, DivideByZeroError
	}
	if massFlowKg < 0 {
		return 0, errors.New("MassFlowKg can't be negative")
	}
	return massFlowKg / density, nil
}
