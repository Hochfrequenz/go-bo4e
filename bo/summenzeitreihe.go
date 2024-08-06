package bo

// Summenzeitreihe is a minimalistic implementation of the BO Summenzeitreihe. But this small implementation alone, allows use to unmarshall boneycombs that contain Summenzeitreihe-BOs
type Summenzeitreihe struct {
	Geschaeftsobjekt
}

func (_ Summenzeitreihe) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
