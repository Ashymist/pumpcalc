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
		t.Errorf("Expected %q error, got %q error instead", DivideByZeroError.Error(), err)
	}

	density = 730
	massFlowKgS = -1
	volumeFlowM3S, err = GetVolumeFlowM3S(massFlowKgS, density)
	if !errors.Is(err, NegativeMassFlowKgSError) || volumeFlowM3S != 0.0 {
		t.Errorf("Expected %q error, got %q error instead", NegativeMassFlowKgSError.Error(), err)
	}
}
