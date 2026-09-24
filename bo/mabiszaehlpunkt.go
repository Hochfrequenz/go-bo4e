package bo

import "regexp"

// MabisZaehlpunkt enthält Informationen zu einem MABIS-Zählpunkt.
type MabisZaehlpunkt struct {
	Geschaeftsobjekt
	// Id ist die MABIS-Zählpunktbezeichnung, z.B. DE 47108151234567
	Id *string `json:"Id,omitempty" validate:"omitempty,alphanum,len=33"`
}

var mabisZaehlpunktIdRegex = regexp.MustCompile(`[A-Z\d]{33}`)

// ValidateId prüft, ob die Id ein gültiges Format hat.
func (m *MabisZaehlpunkt) ValidateId() bool {
	if m.Id == nil {
		return false
	}
	return mabisZaehlpunktIdRegex.MatchString(*m.Id)
}
