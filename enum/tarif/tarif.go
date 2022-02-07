package tarif

// Tarif describes the tarif
//go:generate stringer --type Tarif
//go:generate jsonenums --type Tarif
type Tarif int

const (
	// T1_TARIF1 is "Tarif 1"
	T1_TARIF1 Tarif = iota + 1

	// T2_TARIF2 is "Tarif 2"
	T2_TARIF2

	// T3_TARIF3 is "Tarif 3"
	T3_TARIF3

	// T4_TARIF4 is "Tarif 4"
	T4_TARIF4

	// T5_TARIF5 is "Tarif 5"
	T5_TARIF5

	// T6_TARIF6 is "Tarif 6"
	T6_TARIF6

	// T7_TARIF7 is "Tarif 7"
	T7_TARIF7

	// T8_TARIF8 is "Tarif 8"
	T8_TARIF8

	// T9_TARIF9 is "Tarif 9"
	T9_TARIF9
)
