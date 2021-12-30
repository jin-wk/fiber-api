package util

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
	v   *validator.Validate
)

func Validate(model interface{}) []string {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	v = validator.New()
	en_translations.RegisterDefaultTranslations(v, trans)

	err := v.Struct(model)
	if err != nil {
		var messages []string
		errs := err.(validator.ValidationErrors)
		for _, val := range errs.Translate(trans) {
			messages = append(messages, val)
		}
		return messages
	}
	return nil
}
