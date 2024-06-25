package pjson

import "encoding/json"

// ConvertToBytes ...
func ConvertToBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
