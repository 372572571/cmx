package buildconfig

import "strings"

type ReferenceInformation struct {
	Object string
	Field  string
	Route  string
}

func NewReferenceInformation(reference string) *ReferenceInformation {
	referenceSplit := strings.Split(reference, ".")
	if len(referenceSplit) == 0 {
		return nil
	}

	route := strings.Join(referenceSplit[:len(referenceSplit)-1], ".")
	route = strings.ReplaceAll(route, ".", "/")
	object := referenceSplit[len(referenceSplit)-2]
	field := referenceSplit[len(referenceSplit)-1]

	return &ReferenceInformation{
		Object: object,
		Field:  field,
		Route:  route,
	}
}
