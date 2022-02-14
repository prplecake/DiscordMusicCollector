package dmc

import (
	"regexp"
)

// GetParams parses a URL for parameters and returns a map of named
// group matches.
func GetParams(regex *regexp.Regexp, message string) (paramsMap map[string]string) {
	match := regex.FindStringSubmatch(message)

	paramsMap = make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
