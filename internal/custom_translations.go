package internal

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

func customTranslations(v *validator.Validate, t ut.Translator) {
	_ = v.RegisterTranslation("required", t, func(ut ut.Translator) error {
		return ut.Add("required", "cannot be empty", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
}

func TranslateErrors(err error, t ut.Translator) error {
	var trans []string
	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		trans = append(trans, fmt.Sprintf(`field '%s' with value '%s' %s`, e.Field(), e.Value(), e.Translate(t)))
	}

	return fmt.Errorf(strings.Join(trans, ", "))
}
