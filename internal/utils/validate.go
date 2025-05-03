package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate

func ValidateStruct(input any) error {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	if err := validate.Struct(input); err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				translatedErr := fmt.Errorf(e.Translate(trans))

				return translatedErr
			}
		}
	}

	return nil
}

func IsEmpty(param string) bool {
	return param == ""
}
