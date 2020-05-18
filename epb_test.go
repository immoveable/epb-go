package epb

import "testing"

func TestEnergyClass(t *testing.T) {
	testCases := []struct {
		reg  Regulator
		cons float64
		want string
	}{
		{RegulatorBrussels, 100, "C+"},
		{RegulatorWallonia, 200, "C"},
	}
	for _, tc := range testCases {
		if got := EnergyClass(tc.reg, tc.cons); got != tc.want {
			t.Errorf("%q, %v: want %q, got %q", tc.reg, tc.cons, tc.want, got)
		}
	}
}
