package barbell

// Plate represents a weight plate you put on a barbell.
type Plate float64

// Weight returns the weight of the plate
func (p Plate) Weight() float64 {
	return float64(p)
}

// Plates returns the plates you need to put on a barbell in order
// to achieve the given weight.
func Plates(w float64) []Plate {
	plates := []Plate{45, 35, 25, 10, 5, 2.5}
	result := []Plate{}

	for _, plate := range plates {
		// If we cannot fit these 2 plates on the bar, go to the next plate.
		if w < 2*plate.Weight() {
			continue
		}

		// Find the # of pairs that will fit on the barbell.
		n := int(w / plate.Weight())
		if n%2 == 1 {
			n--
		}

		w -= float64(n) * plate.Weight()

		for ; n > 0; n-- {
			result = append(result, plate)
		}
	}

	return result
}
