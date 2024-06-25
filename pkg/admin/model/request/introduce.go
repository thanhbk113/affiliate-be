package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// IntroduceAll ...
type IntroduceAll struct {
	ParentNameId string `query:"parentNameId"`
}

type IntroduceBodyUpdate struct {
	Content string `json:"content"`
}

// Validate validates the IntroduceBodyUpdate fields.
func (p IntroduceBodyUpdate) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Content, validation.Required.Error("Content is required")),
	)
}

// Validate validates the IntroduceAll fields.
func (p IntroduceAll) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ParentNameId, validation.Required.Error("ParentNameId is required")),
	)
}
