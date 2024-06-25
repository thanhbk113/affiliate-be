package format

import "strconv"

// IntToString ...
func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}
