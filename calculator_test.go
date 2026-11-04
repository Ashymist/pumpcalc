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
	volumeFlowKg, err := GetVolumeFlowKg(massFlowKgS, density)

	expected := 0.013699
	if err != nil || math.Abs(volumeFlowKg-expected) > TOLERANCE {
		t.Errorf("Expected %f,nil with %f tolerance, got %f,%v instead", expected, TOLERANCE, volumeFlowKg, err)
	}

	density = 0
	volumeFlowKg, err = GetVolumeFlowKg(massFlowKgS, density)
	if !errors.Is(err, DivideByZeroError) || volumeFlowKg != 0.0 {
		t.Errorf("Expected %q error, got %q error instead", DivideByZeroError.Error(), err)
	}

	density = 730
	massFlowKgS = -1
	volumeFlowKg, err = GetVolumeFlowKg(massFlowKgS, density)
	if !errors.Is(err, NegativeMassFlowKgSError) || volumeFlowKg != 0.0 {
		t.Errorf("Expected %q error, got %q error instead", NegativeMassFlowKgSError.Error(), err)
	}
}
