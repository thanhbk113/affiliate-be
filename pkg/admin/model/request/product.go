package request

import (
	"affiliate/internal/errorresponse"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ProductBody struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ProductBodyUpdate struct {
	Name        string `json:"name,omitempty"`
	Price       int    `json:"price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Description string `json:"description,omitempty"`
	Unit        string `json:"unit,omitempty"`
	Image       string `json:"image,omitempty"`
}

// ProductAll ...
type ProductAll struct {
	Page      int64  `query:"page"`
	Limit     int64  `query:"limit"`
	Keyword   string `query:"keyword"`
	Sort      string `query:"sort" enums:"newest,oldest"`
	FromPrice int    `query:"fromPrice"`
	ToPrice   int    `query:"toPrice"`
}

type ProductAllUser struct {
	NameID    string `query:"nameId"`
	Page      int64  `query:"page"`
	Limit     int64  `query:"limit"`
	Keyword   string `query:"keyword"`
	Sort      string `query:"sort" enums:"newest,oldest"`
	FromPrice int    `query:"fromPrice"`
	ToPrice   int    `query:"toPrice"`
}

// Validate ...
func (p *ProductBody) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Name, validation.Required.Error(errorresponse.CommonKeyNameIsRequired)),
		validation.Field(&p.Price, validation.Required.Error(errorresponse.CommonKeyPriceIsRequired)),
		validation.Field(&p.Quantity, validation.Required.Error(errorresponse.CommonKeyQuantityIsRequired)),
		validation.Field(&p.Description, validation.Required.Error(errorresponse.CommonKeyDescriptionIsRequired)),
		validation.Field(&p.Image, validation.Required.Error(errorresponse.CommonKeyImageIsRequired)),
		validation.Field(&p.Unit, validation.Required.Error(errorresponse.CommonKeyUnitIsRequired)),
	)
}

// Validate ...
func (p ProductAll) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Page,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyPageInvalid)),
		validation.Field(&p.Limit,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyLimitInvalid)),
	)
}

// Validate ...
func (p ProductAllUser) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Page,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyPageInvalid)),
		validation.Field(&p.Limit,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyLimitInvalid)),
		validation.Field(&p.NameID, validation.Required.Error(errorresponse.CommonKeyNameIDIsRequired)),
	)
}

// Validate ...
func (p *ProductBodyUpdate) Validate() error {
	return validation.ValidateStruct(p)
}
