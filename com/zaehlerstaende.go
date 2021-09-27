package com

// Zaehlerstaende is a table type of Zaehlerstand with defined descending sorting of the Zaehlerstand.Ablesedatum for sort.Sort(...)
type Zaehlerstaende []Zaehlerstand

// Len returns the size of z
func (z Zaehlerstaende) Len() int {
	return len(z)
}

// Swap swaps the entries at index i and j
func (z Zaehlerstaende) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}

// Less returns true iff the Zaehlerstand.Ablesedatum at index i is after the Zaehlerstand.Ablesedatum at index j
func (z Zaehlerstaende) Less(i, j int) bool {
	return z[i].Ablesedatum.After(z[j].Ablesedatum)
}
