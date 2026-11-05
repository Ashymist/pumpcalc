package pumpcalc_test

import (
	"errors"
	"math"
	"testing"

	. "github.com/Ashymist/pumpcalc"
)

const TOLERANCE = 1e-6

func TestGetVolumeFlowKg(t *testing.T) {
	massFlowKgS := 10.0
	density := 730.0
	volumeFlowM3S, err := GetVolumeFlowM3S(massFlowKgS, density)

	expected := 0.013699
	if err != nil || math.Abs(volumeFlowM3S-expected) > TOLERANCE {
		t.Errorf("Expected %f,nil with %f tolerance, got %f,%v instead", expected, TOLERANCE, volumeFlowM3S, err)
	}

	density = 0
	volumeFlowM3S, err = GetVolumeFlowM3S(massFlowKgS, density)
	if !errors.Is(err, DivideByZeroError) || volumeFlowM3S != 0.0 {
		t.Errorf("Expected %q error, got %v error instead", DivideByZeroError.Error(), err)
	}

	density = 730
	massFlowKgS = -1
	volumeFlowM3S, err = GetVolumeFlowM3S(massFlowKgS, density)
	if !errors.Is(err, NegativeMassFlowKgSError) || volumeFlowM3S != 0.0 {
		t.Errorf("Expected %q error, got %v error instead", NegativeMassFlowKgSError.Error(), err)
	}

	density = -1
	massFlowKgS = 10.0
	_, err = GetVolumeFlowM3S(massFlowKgS, density)
	if !errors.Is(err, NegativeDensityError) {
		t.Errorf("Expected %q error, got %v error instead", NegativeMassFlowKgSError.Error(), err)
	}
}

func TestGetPipeDiameter(t *testing.T) {
	volumeFlowM3S := 0.013699
	velocity := 0.8

	diameter, err := GetPipeDiameter(volumeFlowM3S, velocity)
	expected := 0.147657
	if err != nil || math.Abs(expected-diameter) > TOLERANCE {
		t.Errorf("Expected %f,nil with %f tolerance, got %f,%v instead", expected, TOLERANCE, diameter, err)
	}

	_, err = GetPipeDiameter(volumeFlowM3S, 0)
	if !errors.Is(err, DivideByZeroError) {
		t.Errorf("Expected %q error, got %v error instead", DivideByZeroError.Error(), err)
	}

	_, err = GetPipeDiameter(volumeFlowM3S, -1)
	if !errors.Is(err, NegativeVelocityError) {
		t.Errorf("Expected %q error, got %v error instead", NegativeVelocityError.Error(), err)
	}
}

func TestGetReynoldsNumber(t *testing.T) {
	velocity := 0.8
	pipeDiameter := 0.1477
	viscosity := 0.7 * 1e-6

	reynold, err := GetReynoldsNumber(velocity, pipeDiameter, viscosity)
	expected := 1.688 * 1e+5
	if err != nil || math.Abs(expected-reynold) > TOLERANCE {
		t.Errorf("Expected %f,nil with %f tolerance, got %f,%v instead", expected, TOLERANCE, reynold, err)
	}

	_, err = GetReynoldsNumber(velocity, pipeDiameter, 0)
	if !errors.Is(err, DivideByZeroError) {
		t.Errorf("Expected %q error, got %v error instead", DivideByZeroError.Error(), err)
	}

	_, err = GetReynoldsNumber(velocity, pipeDiameter, -1)
	if !errors.Is(err, NegativeKinematicViscosityError) {
		t.Errorf("Expected %q error, got %v error instead", DivideByZeroError.Error(), err)
	}
}

func TestCalculatePipeLosses(t *testing.T) {
	frictionCoef := 0.022
	length := 28.0
	diameter := 0.1477
	frictionSum := 7.8
	velocity := 0.8
	gravity := 9.81

	losses, err := CalculateLossesOnPipe(length, diameter, frictionCoef, frictionSum, velocity, gravity)
	expected := 0.390479
	if err != nil || math.Abs(expected-losses) > TOLERANCE {
		t.Errorf("Expected %f,nil with %f tolerance, got %f,%v instead", expected, TOLERANCE, losses, err)
	}

	_, err = CalculateLossesOnPipe(length, diameter, 0, frictionSum, velocity, gravity)
	if !errors.Is(err, NotPositiveFrictionCoefficientError) {
		t.Errorf("Expected %q error, got %v error instead", NotPositiveFrictionCoefficientError.Error(), err)
	}

	_, err = CalculateLossesOnPipe(length, diameter, frictionCoef, frictionSum, velocity, 0)
	if !errors.Is(err, NotPositiveGravityError) {
		t.Errorf("Expected %q error, got %v error instead", NotPositiveGravityError.Error(), err)
	}
}

func TestCalculatePumpFlow(t *testing.T) {
	heightAfter := 26.0
	heightBefore := 4.0
	pressureBefore := 101300.0
	pressureAfter := 122500.0
	density := 730.0
	gravity := 9.81
	velocityAfter := 2.0
	lossesBeforePump := 0.391
	lossesAfterPump := 5.087

	flow, err := CalculatePumpFlow(heightBefore, heightAfter, pressureBefore,
		pressureAfter, density, gravity,
		0, velocityAfter, lossesBeforePump, lossesAfterPump)
	expected := 30.642230
	if err != nil || math.Abs(expected-flow) > TOLERANCE {
		t.Errorf("Expected %f,nil with %f tolerance, got %f,%v instead", expected, TOLERANCE, flow, err)
	}

	_, err = CalculatePumpFlow(heightBefore, heightAfter, pressureBefore,
		pressureAfter, 0, 0,
		0, velocityAfter, lossesBeforePump, lossesAfterPump)
	if !errors.Is(err, DivideByZeroError) {
		t.Errorf("Expected %q error, got %v error instead", DivideByZeroError.Error(), err)
	}
}

func TestCalculatePumpPressure(t *testing.T) {
	density := 730.0
	gravity := 9.81
	pumpFlow := 30.642

	pressure := CalculatePumpPressure(pumpFlow, density, gravity)
	expected := 219436.5546
	if math.Abs(expected-pressure) > TOLERANCE {
		t.Errorf("Expected %f with %f tolerance, got %f instead", expected, TOLERANCE, pressure)
	}
}

func TestCalculatePumpNozzlePressureBefore(t *testing.T) {
	pressure := 101300.0
	density := 730.0
	gravity := 9.81
	height := 4.0
	losses := 0.391
	velocity := 0.8

	pressureBefore, err := CalculatePumpNozzlePressureBefore(pressure, density, gravity, height, losses, velocity)
	expected := 126911.5317
	if err != nil || math.Abs(expected-pressureBefore) > TOLERANCE {
		t.Errorf("Expected %f,nil with %f tolerance, got %f,%v instead", expected, TOLERANCE, pressureBefore, err)
	}

	_, err = CalculatePumpNozzlePressureBefore(pressure, density, 0, height, losses, velocity)
	if !errors.Is(err, DivideByZeroError) {
		t.Errorf("Expected %q error, got %v error instead", DivideByZeroError.Error(), err)
	}
}

func TestCalculatePumpNozzlePressureAfter(t *testing.T) {
	pressure := 122500.0
	density := 730.0
	gravity := 9.81
	height := 26.0
	losses := 5.087

	pressureAfter := CalculatePumpNozzlePressureAfter(pressure, density, gravity, height, losses)
	expected := 345123.3331
	if math.Abs(expected-pressureAfter) > TOLERANCE {
		t.Errorf("Expected %f with %f tolerance, got %f instead", expected, TOLERANCE, pressureAfter)
	}
}
