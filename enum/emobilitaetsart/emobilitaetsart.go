package emobilitaetsart

// Im Falle der E-Mobilität bei
//
//go:generate stringer --type EMobilitaetsart
//go:generate jsonenums --type EMobilitaetsart

type EMobilitaetsart int

const (
	WALLBOX                 EMobilitaetsart = iota + 1 // ZE6: Wallbox
	E_MOBILITAETSLADESAEULE                            // Z87: E-Mobilitätsladesäule
	LADEPARK                                           // ZE7: Ladepark
)
