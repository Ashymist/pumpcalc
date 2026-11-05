package pumpcalc

import "math"

func GetVolumeFlowM3S(massFlowKgS float64, density float64) (float64, error) {
	if density == 0 {
		return 0, DivideByZeroError
	} else if density < 0 {
		return 0, NegativeDensityError
	}
	if massFlowKgS < 0 {
		return 0, NegativeMassFlowKgSError
	}
	return massFlowKgS / density, nil
}

func GetPipeDiameter(volumeFlowM3S float64, fluidVelocity float64) (float64, error) {
	if fluidVelocity == 0 {
		return 0, DivideByZeroError
	}
	if fluidVelocity < 0 {
		return 0, NegativeVelocityError
	}

	return math.Sqrt((4 * volumeFlowM3S) / (math.Pi * fluidVelocity)), nil
}

func GetReynoldsNumber(fluidVelocity float64, pipeDiameter float64, kinematicViscosity float64) (float64, error) {
	if kinematicViscosity == 0 {
		return 0, DivideByZeroError
	} else if kinematicViscosity < 0 {
		return 0, NegativeKinematicViscosityError
	}

	return (fluidVelocity * pipeDiameter) / kinematicViscosity, nil
}
