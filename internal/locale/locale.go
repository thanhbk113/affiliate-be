package locale

import (
	"encoding/json"
	"path"
	"runtime"
	"strings"

	"github.com/thoas/go-funk"
)

// Langs ...
const (
	LangEn = "en"
	LangVi = "vi"
)

type (
	// Locale ...
	Locale struct {
		Key     string
		Code    int      `json:"code"`
		Message *Message `json:"message"`
	}

	// Message ...
	Message struct {
		En      string
		Vi      string
		Display string
	}
)

func getLocalePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}

// GetDisplay return text with language
func (msg *Message) GetDisplay(lang string) {
	text := funk.Get(msg, strings.Title(lang))
	if text != nil {
		msg.Display = text.(string)
	} else {
		msg.Display = "N/A"
	}
}

// MarshalJSON ...
func (msg *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(msg.Display)
}

var notFoundKey = Locale{
	Key: "NotFound",
	Message: &Message{
		En: "Not found",
		Vi: "Không tìm thấy",
	},
	Code: -1,
}

// Default key for each error
const (
	Default200 = CommonKeySuccess
	Default400 = CommonKeyBadRequest
	Default401 = CommonKeyUnauthorized
	Default403 = CommonKeyNoPermission
	Default404 = CommonKeyNotFound
)

// GetByKey give key and receive message + code
func GetByKey(lang string, key string) Locale {
	item := funk.Find(list, func(item Locale) bool {
		return item.Key == key
	})

	if item == nil {
		return notFoundKey
	}

	return item.(Locale)
}

// GetViMessageByKey ...
func GetViMessageByKey(key string) string {
	l := GetByKey(LangVi, key)
	return l.Message.Vi
}

// GetMessageByKey ...
func GetMessageByKey(lang, key string) string {
	l := GetByKey(lang, key)
	l.Message.GetDisplay(lang)
	return l.Message.Display
}

var list = make([]Locale, 0)

// LoadProperties ...
func LoadProperties() {
	// Assign locales
}
