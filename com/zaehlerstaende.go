package com

// Zaehlerstaende is a table type of Zaehlerstand with defined descending sorting of the Zaehlerstand.Ablesedatum for sort.Sort(...)
type Zaehlerstaende []Zaehlerstand

func (z Zaehlerstaende) Len() int {
	return len(z)
}

func (z Zaehlerstaende) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}

func (z Zaehlerstaende) Less(i, j int) bool {
	return z[i].Ablesedatum.After(z[j].Ablesedatum)
}
