package kundengruppeKA

//go:generate stringer --type KundengruppeKA
//go:generate jsonenums --type KundengruppeKA
type KundengruppeKA int

const (
	S_TARIF_25000 KundengruppeKA = iota + 1
	S_TARIF_100000
	S_TARIF_500000
	S_TARIF_G_500000
	S_SONDERKUNDE
	G_KOWA_25000
	G_KOWA_100000
	G_KOWA_500000
	G_KOWA_G_500000
	G_TARIF_25000
	G_TARIF_100000
	G_TARIF_500000
	G_TARIF_G_500000
	G_SONDERKUNDE
	SONDER_KAS
	SONDER_SAS
	SONDER_TAS
	SONDER_TKS
	SONDER_TSS
)
