package epb

// Regulator represents an energy regulator.
type Regulator string

// Energy regulators
const (
	RegulatorBrussels Regulator = "brussels"
	RegulatorFlanders Regulator = "flanders"
	RegulatorFrance   Regulator = "france"
	RegulatorWallonia Regulator = "wallonia"
)

// EnergyClass returns the energy class of a building (like `A++`, `A+`, `A`, `B`...) following the regulator rules, given its energy consumption (in in kWh/m².year).
// If the energy class cannot be calculated, en empty string is returned.
func EnergyClass(reg Regulator, cons float64) string {
	switch reg {

	case RegulatorBrussels:
		if cons <= 0 {
			return "A++"
		} else if cons <= 15 {
			return "A+"
		} else if cons <= 30 {
			return "A"
		} else if cons <= 45 {
			return "A-"
		} else if cons <= 62 {
			return "B+"
		} else if cons <= 78 {
			return "B"
		} else if cons <= 95 {
			return "B-"
		} else if cons <= 113 {
			return "C+"
		} else if cons <= 132 {
			return "C"
		} else if cons <= 150 {
			return "C-"
		} else if cons <= 170 {
			return "D+"
		} else if cons <= 190 {
			return "D"
		} else if cons <= 210 {
			return "D-"
		} else if cons <= 232 {
			return "E+"
		} else if cons <= 253 {
			return "E"
		} else if cons <= 275 {
			return "E-"
		} else if cons <= 345 {
			return "F"
		}
		return "G"

	case RegulatorFlanders:
		return ""

	case RegulatorFrance:
		if cons <= 50 {
			return "A"
		} else if cons <= 90 {
			return "B"
		} else if cons <= 150 {
			return "C"
		} else if cons <= 230 {
			return "D"
		} else if cons <= 330 {
			return "E"
		} else if cons <= 450 {
			return "F"
		}
		return "G"

	case RegulatorWallonia:
		if cons <= 0 {
			return "A++"
		} else if cons <= 45 {
			return "A+"
		} else if cons <= 85 {
			return "A"
		} else if cons <= 170 {
			return "B"
		} else if cons <= 255 {
			return "C"
		} else if cons <= 340 {
			return "D"
		} else if cons <= 425 {
			return "E"
		} else if cons <= 510 {
			return "F"
		}
		return "G"
	}

	return ""
}

// TotalConsumption returns the total consumption of a building in a year, given its energy consumption (in kWh/m².year) and its total area (in m²).
func TotalConsumption(cons, area float64) float64 {
	return cons * area
}
