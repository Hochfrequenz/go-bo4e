package fernsteuerbarkeitstatus

// Status der Fernsteuerbarkeit einer Marktlokation gemäß den Vorgaben für die Direktvermarktung.

//go:generate stringer --type FernsteuerbarkeitStatus
//go:generate jsonenums --type FernsteuerbarkeitStatus
type FernsteuerbarkeitStatus int

const (
	// Marktlokation ist technisch nicht fernsteuerbar.
	// Der NB bestätigt mit der Anmeldung einer erzeugenden Marktlokation zur Direktvermarktung, dass die Anlage nicht mit einer Fernsteuerung ausgestattet ist und nicht fernsteuerbar ist. Die Voraussetzung zur Zahlung der Managementprämie für fernsteuerbare Anlagen ist nicht gegeben. Z96
	NICHT_FERNSTEUERBAR     FernsteuerbarkeitStatus = iota + 1
	TECHNISCH_FERNSTEUERBAR                         // Marktlokation ist technisch fernsteuerbar. Der NB bestätigt mit der Anmeldung einer erzeugenden Marktlokation zur Direktvermarktung, dass die Marktlokation mit einer Fernsteuerung ausgestattet, aber dem NB keine Information darüber vorliegt, dass der LF die Marktlokation fernsteuern kann. Die Voraussetzung zur Zahlung der Managementprämie für fernsteuerbare Marktlokation ist nicht gegeben. Z97
	LIEFERANT_FERNSTEUERBAR                         // Marktlokation ist durch den Lieferanten fernsteuerbar. Der NB bestätigt mit der Anmeldung einer Marktlokation zur Direktvermarktung, dass die Marktlokation mit einer Fernsteuerung ausgestattet ist und der LF diese auch fernsteuern kann. Die Voraussetzung zur Zahlung der Managementprämie für fernsteuerbare Marktlokationen ist gegeben. Z98
)
