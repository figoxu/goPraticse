package ut

import "strings"

func HasFeature(msg string, features ...string) bool {
	msg = strings.ToLower(msg)
	for _, feature := range features {
		if strings.Contains(msg, strings.ToLower(feature)) {
			return true
		}
	}
	return false
}
