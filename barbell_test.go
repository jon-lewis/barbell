package barbell

import "testing"

type barbellTest struct {
	name   string
	input  float64
	output []Plate
}

func isEqual(p1 []Plate, p2 []Plate) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}

	return true
}

// TestCalculation tests that the barbell calculates the correct plates.
func TestCalculation(t *testing.T) {
	tests := []barbellTest{
		{"90 == [45 45]", 90, []Plate{45, 45}},
		{"95 == [45 45 2.5 2.5]", 95, []Plate{45, 45, 2.5, 2.5}},
		{"115 == [45 45 10 10 2.5 2.5]", 115, []Plate{45, 45, 10, 10, 2.5, 2.5}},
		{"2.5 == []", 2.5, []Plate{}},
		{"5 == [2.5 2.5]", 5, []Plate{2.5, 2.5}},
		{"430 == [45 45 45 45 45 45 45 45 35 35]", 430, []Plate{45, 45, 45, 45, 45, 45, 45, 45, 35, 35}},
	}

	for _, test := range tests {
		actual := Plates(test.input)
		if !isEqual(actual, test.output) {
			t.Error("Actual value did not equal expected", test.name, "got", actual)
		}
	}
}
