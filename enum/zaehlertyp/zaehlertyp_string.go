// Code generated by "stringer --type Zaehlertyp"; DO NOT EDIT.

package zaehlertyp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Drehstromzaehler-1]
	_ = x[Balgengaszaehler-2]
	_ = x[Drehkolbenzaehler-3]
	_ = x[SmartMeter-4]
	_ = x[Leistungszaehler-5]
	_ = x[Maximumgaehler-6]
	_ = x[Turbinenradgaszaehler-7]
	_ = x[Ultraschallgaszaehler-8]
	_ = x[Wechselstromzaehler-9]
}

const _Zaehlertyp_name = "DrehstromzaehlerBalgengaszaehlerDrehkolbenzaehlerSmartMeterLeistungszaehlerMaximumgaehlerTurbinenradgaszaehlerUltraschallgaszaehlerWechselstromzaehler"

var _Zaehlertyp_index = [...]uint8{0, 16, 32, 49, 59, 75, 89, 110, 131, 150}

func (i Zaehlertyp) String() string {
	i -= 1
	if i < 0 || i >= Zaehlertyp(len(_Zaehlertyp_index)-1) {
		return "Zaehlertyp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Zaehlertyp_name[_Zaehlertyp_index[i]:_Zaehlertyp_index[i+1]]
}