package format

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"affiliate/internal/constants"

	"github.com/gosimple/slug"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// FormatPhoneCommon common phone number
func FormatPhoneCommon(phone string) string {
	// Return if there is no phone
	if phone == "" {
		return ""
	}

	// Split all spaces
	phone = strings.Replace(phone, " ", "", -1)

	// If first character is "0", replace to "+84", become "+84....."
	if string(phone[0]) == "0" {
		phone = strings.Replace(phone, "0", "+84", 1)
	}

	// If 2 first characters is "84", add "+"
	if phone[0:2] == "84" {
		phone = "+" + phone
	}

	return phone
}

// PhoneNumber ...
type PhoneNumber struct {
	Number      string
	CountryCode string
	Full        string
}

// TrimAndLowerCase ...
func TrimAndLowerCase(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

// RemoveSpace ...
func RemoveSpace(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

// PhoneNumberFormatFromPhone ...
func PhoneNumberFormatFromPhone(phone string) PhoneNumber {
	var phoneNumber PhoneNumber
	if len(phone) == 9 && string(phone[0]) != "0" {
		phone = "0" + phone
	}
	isValid := PhoneNumberIsValid(phone)
	if !isValid {
		return phoneNumber
	}

	phone = strings.Replace(phone, " ", "", -1)

	switch true {
	case string(phone[0]) == "0":
		phone = strings.Replace(phone, "0", "", 1)
	case phone[0:3] == "+84":
		phone = strings.Replace(phone, "+84", "", 1)
	case phone[0:2] == "84":
		phone = strings.Replace(phone, "84", "", 1)
	}

	phoneNumber.Number = phone
	phoneNumber.CountryCode = "+84"
	phoneNumber.Full = phoneNumber.CountryCode + phoneNumber.Number
	return phoneNumber
}

// RemoveEnterMultiValue ...
func RemoveEnterMultiValue(data []string) []string {
	var (
		response = make([]string, 0)
	)
	for _, value := range data {
		response = append(response, RemoveEnter(value))
	}
	return response
}

// RemoveEnter ...
func RemoveEnter(content string) string {
	return strings.ReplaceAll(strings.ReplaceAll(content, "\n", " "), "\r", " ")
}

// ConvertSlugProvince ...
func ConvertSlugProvince(slug string) string {
	if slug == "ho-chi-minh" {
		return "tp-ho-chi-minh"
	}
	return slug
}

// ConvertSlugProvinceFind ...
func ConvertSlugProvinceFind(slug string) string {
	if slug == "tp-ho-chi-minh" {
		return "ho-chi-minh"
	}
	return slug
}

// NonAccentVietnamese ...
func NonAccentVietnamese(str string) string {
	if str != "" {
		str = strings.ToLower(str)
		str = replaceStringWithRegex(str, `đ`, "d")
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		result, _, _ := transform.String(t, str)
		result = replaceStringWithRegex(result, `[^a-zA-Z0-9\s]`, "")

		return result
	}
	return ""
}

// RemoveAccentVietnamese ...
func RemoveAccentVietnamese(str string) string {
	if str != "" {
		str = strings.ToLower(str)
		str = replaceStringWithRegex(str, `đ`, "d")
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		result, _, _ := transform.String(t, str)
		return strings.TrimSpace(result)
	}
	return ""
}

// isMn ...
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// replaceStringWithRegex ...
func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}

// ConvertStringToSlug ...
func ConvertStringToSlug(text string) string {
	return slug.Make(text)
}

// PhoneNumberIsValid check phone number is valid
func PhoneNumberIsValid(phone string) bool {
	r := regexp.MustCompile(constants.RegexPhoneNumber)
	return r.MatchString(phone)
}

// RemoveDoubleSpace ...
func RemoveDoubleSpace(s string) string {
	s = strings.Trim(s, " ")
	var (
		values = make([]string, 0)
	)
	splits := strings.Split(s, " ")
	for _, value := range splits {
		if value != "" {
			values = append(values, value)
		}
	}
	return strings.Join(values, " ")
}

// StringToInt64 ...
func StringToInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

// SearchString ...
func SearchString(keyword string) bson.M {
	return bson.M{
		"$regex": primitive.Regex{
			Pattern: NonAccentVietnamese(keyword),
			Options: "i",
		},
	}
}
