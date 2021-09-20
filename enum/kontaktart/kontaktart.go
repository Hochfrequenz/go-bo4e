package kontaktart

//go:generate stringer --type Kontaktart
//go:generate jsonenums --type Kontaktart
// Kontaktart ist die Art des Kontaktes zwischen Geschäftspartnern.
type Kontaktart int

const (
	Anschreiben Kontaktart = iota + 1 // Anschreiben
	Telefonat                         // Telefonat
	Fax                               // Fax
	Email                             // E-Mail
	SMS                               // SMS
)
