package definitionennotwendigkeit

//go:generate stringer --type DefinitionenNotwendigkeit
//go:generate jsonenums --type DefinitionenNotwendigkeit
type DefinitionenNotwendigkeit int

const (
	ZAEHLZEITDEFINITIONEN_WERDEN_VERWENDET DefinitionenNotwendigkeit = iota + 1
	ZAEHLZEITDEFINITIONEN_WERDEN_NICHT_VERWENDET
	DEFINITIONEN_WERDEN_VERWENDET
	DEFINITIONEN_WERDEN_NICHT_VERWENDET
)
