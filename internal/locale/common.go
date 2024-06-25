package locale

import (
	"log"

	"github.com/magiconair/properties"
)

const (
	CommonKeySuccess         = "success"
	CommonKeyBadRequest      = "badRequest"
	CommonKeyUnauthorized    = "unauthorized"
	CommonKeyNotFound        = "notFound"
	CommonKeyNoPermission    = "noPermission"
	CommonKeyForbidden       = "forbidden"
	CommonKeyErrorWhenHandle = "errorWhenHandle"

	// IsRequired
	CommonKeyNameIsRequired = "nameIsRequired"
)

const (
	CommonKeySuccessCode = iota + 1
	CommonKeyBadRequestCode
	CommonKeyUnauthorizedCode
	CommonKeyNotFoundCode
	CommonKeyNoPermissionCode
	CommonKeyNameIsRequiredCode
	CommonKeyForbiddenCode
	CommonKeyErrorWhenHandleCode
)

type (
	commonLang struct {
		CommonKeySuccess         string `properties:"commonKeySuccess"`
		CommonKeyBadRequest      string `properties:"commonKeyBadRequest"`
		CommonKeyUnauthorized    string `properties:"commonKeyUnauthorized"`
		CommonKeyNotFound        string `properties:"commonKeyNotFound"`
		CommonKeyNoPermission    string `properties:"commonKeyNoPermission"`
		CommonKeyNameIsRequired  string `properties:"commonKeyNameIsRequired"`
		CommonKeyForbidden       string `properties:"commonKeyForbidden"`
		CommonKeyErrorWhenHandle string `properties:"commonKeyErrorWhenHandle"`
	}
)

var (
	commonEn commonLang
	commonVi commonLang
)

func init() {
	// Load properties
	p1 := properties.MustLoadFile(getLocalePath()+"/properties/en/common.properties", properties.UTF8)
	if err := p1.Decode(&commonEn); err != nil {
		log.Fatal(err)
	}
	p2 := properties.MustLoadFile(getLocalePath()+"/properties/vi/common.properties", properties.UTF8)
	if err := p2.Decode(&commonVi); err != nil {
		log.Fatal(err)
	}
}

func commonLoadLocales() []Locale {
	return []Locale{
		{
			Key: CommonKeySuccess,
			Message: &Message{
				En: commonEn.CommonKeySuccess,
				Vi: commonVi.CommonKeySuccess,
			},
			Code: CommonKeySuccessCode,
		},
		{
			Key: CommonKeyBadRequest,
			Message: &Message{
				En: commonEn.CommonKeyBadRequest,
				Vi: commonVi.CommonKeyBadRequest,
			},
			Code: CommonKeyBadRequestCode,
		},
		{
			Key: CommonKeyUnauthorized,
			Message: &Message{
				En: commonEn.CommonKeyUnauthorized,
				Vi: commonVi.CommonKeyUnauthorized,
			},
			Code: CommonKeyUnauthorizedCode,
		},
		{
			Key: CommonKeyNotFound,
			Message: &Message{
				En: commonEn.CommonKeyNotFound,
				Vi: commonVi.CommonKeyNotFound,
			},
			Code: CommonKeyNotFoundCode,
		},
		{
			Key: CommonKeyNoPermission,
			Message: &Message{
				En: commonEn.CommonKeyNoPermission,
				Vi: commonVi.CommonKeyNoPermission,
			},
			Code: CommonKeyNoPermissionCode,
		},
		{
			Key: CommonKeyNameIsRequired,
			Message: &Message{
				En: commonEn.CommonKeyNameIsRequired,
				Vi: commonVi.CommonKeyNameIsRequired,
			},
			Code: CommonKeyNameIsRequiredCode,
		},
		{
			Key: CommonKeyForbidden,
			Message: &Message{
				En: commonEn.CommonKeyForbidden,
				Vi: commonVi.CommonKeyForbidden,
			},
			Code: CommonKeyForbiddenCode,
		},
		{
			Key: CommonKeyErrorWhenHandle,
			Message: &Message{
				En: commonEn.CommonKeyErrorWhenHandle,
				Vi: commonVi.CommonKeyErrorWhenHandle,
			},
			Code: CommonKeyErrorWhenHandleCode,
		},
	}
}
