package pjson

import "encoding/json"

// ConvertToJSONString ...
func ConvertToJSONString(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
