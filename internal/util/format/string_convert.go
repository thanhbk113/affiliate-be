package format

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// ToString ...
func ToString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// StringToInt ...
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StringsToInterface ...
func StringsToInterface(strings []string) (result []interface{}) {
	for _, s := range strings {
		result = append(result, s)
	}
	return
}

// NonAccentVietnamese ...
func NonAccentVietnamese(str string) string {
	str = strings.ToLower(str)
	str = replaceStringWithRegex(str, `đ`, "d")
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, str)
	result = replaceStringWithRegex(result, `[^a-zA-Z0-9\s]`, "")

	return result
}

// replaceStringWithRegex ...
func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}
