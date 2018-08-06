package utils

import "encoding/json"

// AnyTypeToStr ...
func AnyTypeToStr(value interface{}) (s string) {
	switch v := value.(type) {
	case []byte:
		s = string(v)
	default:
		tmp, _ := json.Marshal(value)
		s = string(tmp)
	}
	return s
}
