package pumpcalc_test

import (
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
	if err == nil || volumeFlowKg != 0.0 {
		t.Errorf("Can't divide by zero")
	}
}
