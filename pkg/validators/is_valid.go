package validators

import "regexp"

func IsValid(uuid string) bool {
	//regular expression used to check for validity of the ID given as URL parameter.
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	//if matching return true else false
	return r.MatchString(uuid)
}
