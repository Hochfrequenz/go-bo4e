package kontaktart

// Kontaktart ist die Art des Kontaktes zwischen bo.Geschaeftspartner n.
//go:generate stringer --type Kontaktart
//go:generate jsonenums --type Kontaktart
type Kontaktart int

const (
	// ANSCHREIBEN means mail (on paper)
	ANSCHREIBEN Kontaktart = iota + 1
	TELEFONAT              // TELEFONAT means phone
	FAX                    // FAX means Fax
	EMAIL                  // EMAIL means E-Mail
	SMS                    // SMS means text message
)
