package filter

import (
	"errors"
	"regexp"
)

func filter(input string, filters []string) (bool, error) {

	// return early with default setting
	if filters == nil {
		return true, nil
	}

	// check for matches
	for _, filter := range filters {
		re, err := regexp.Compile(filter)
		if err != nil {
			return false, errors.New("cannot compile malformed filter string")
		}

		if re.MatchString(input) {
			return false, nil
		}
	}

	// allowed if no matches found
	return true, nil
}
