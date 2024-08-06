package bo

// Einspeisung is a minimalistic implementation of the BO Einspeisung. But this small implementation alone, allows use to unmarshall boneycombs that contain Einspeisung-BOs
type Einspeisung struct {
	Geschaeftsobjekt
}

func (_ Einspeisung) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
