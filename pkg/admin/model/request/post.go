package request

import (
	"affiliate/internal/errorresponse"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ParPost struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
type SubPost struct {
	Pass    string `json:"pass"`
	ParID   string `json:"parId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SubPostUpdate struct {
	Pass    string `json:"pass"`
	Content string `json:"content"`
}

type SubPostAll struct {
	Page  int64  `query:"page"`
	Limit int64  `query:"limit"`
	Sort  string `query:"sort" enums:"newest,oldest"`
}

type ParPostAll struct {
	Page  int64  `query:"page"`
	Limit int64  `query:"limit"`
	Sort  string `query:"sort" enums:"newest,oldest"`
}

func (p *ParPost) Validate() error {

	return validation.ValidateStruct(p,
		validation.Field(&p.Name, validation.Required.Error(errorresponse.CommonKeyNameIsRequired)),
		validation.Field(&p.Pass, validation.Required.Error(errorresponse.CommonKeyPassIsRequired)),
	)
}

func (p *SubPostUpdate) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Pass, validation.Required.Error(errorresponse.CommonKeyPassIsRequired)),
		validation.Field(&p.Content, validation.Required.Error(errorresponse.CommonKeyContentIsRequired)),
	)
}

func (p *SubPost) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Pass, validation.Required.Error(errorresponse.CommonKeyPassIsRequired)),
		validation.Field(&p.ParID, validation.Required.Error(errorresponse.CommonKeyParIDIsRequired)),
		validation.Field(&p.Title, validation.Required.Error(errorresponse.CommonKeyTitleIsRequired)),
		validation.Field(&p.Content, validation.Required.Error(errorresponse.CommonKeyContentIsRequired)),
	)
}

func (p *ParPostAll) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Page,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyPageInvalid)),
		validation.Field(&p.Limit,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyLimitInvalid)),
	)
}

func (p *SubPostAll) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Page,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyPageInvalid)),
		validation.Field(&p.Limit,
			validation.Min(int64(0)).Error(errorresponse.CommonKeyLimitInvalid)),
	)
}
