package ut

import (
	"encoding/json"
	"strings"
)

func HasFeature(msg string, features ...string) bool {
	msg = strings.ToLower(msg)
	for _, feature := range features {
		if strings.Contains(msg, strings.ToLower(feature)) {
			return true
		}
	}
	return false
}

func JsonString(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(b)
}
