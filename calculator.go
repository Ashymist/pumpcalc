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

func CalculateLossesOnPipe(pipeLength float64, pipeDiameter float64, frictionCoefficient float64,
	frictionSum float64, fluidVelocity float64, gravity float64) (float64, error) {
	if frictionCoefficient <= 0 {
		return 0, NotPositiveFrictionCoefficientError
	}
	if gravity <= 0 {
		return 0, NotPositiveGravityError
	}

	return (frictionCoefficient*(pipeLength/pipeDiameter) + frictionSum) * ((math.Pow(fluidVelocity, 2)) / (2 * gravity)), nil
}

func CalculatePumpFlow(fluidHeightReserv float64, fluidHeightNeeded float64, pressureReserv float64,
	pressureNeeded float64, density float64, gravity float64, velocityBeforePump float64, velocityAfterPump float64,
	lossesBeforePump float64, lossesAfterPump float64) (float64, error) {
	if gravity == 0 || density == 0 {
		return 0, DivideByZeroError
	}

	return (fluidHeightNeeded - fluidHeightReserv) +
		((pressureNeeded - pressureReserv) / (density * gravity)) +
		((math.Pow(velocityAfterPump, 2) - math.Pow(velocityBeforePump, 2)) / (2 * gravity)) +
		lossesAfterPump + lossesBeforePump, nil
}

func CalculatePumpPressure(pumpFlow float64, density float64, gravity float64) float64 {
	return density * gravity * pumpFlow
}

func CalculatePumpNozzlePressureBefore(pressureReserv float64, density float64, gravity float64,
	heightBefore float64, lossesBefore float64, velocityBefore float64) (float64, error) {
	if gravity == 0 {
		return 0, DivideByZeroError
	}

	return pressureReserv + density*gravity*(heightBefore-lossesBefore-(math.Pow(velocityBefore, 2)/(2*gravity))), nil
}
