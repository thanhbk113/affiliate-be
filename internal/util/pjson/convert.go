package pjson

// ConvertStringsToInterfaces ...
func ConvertStringsToInterfaces(data []string) (res []interface{}) {
	for _, d := range data {
		res = append(res, d)
	}
	return res
}
