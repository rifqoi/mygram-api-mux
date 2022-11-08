package controller

import (
	"encoding/json"
	"io"
	"unicode"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ReadJSON(r io.Reader, data any) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(data)
	return err
}

func Validate(s interface{}) []string {
	validate := validator.New()
	validate.RegisterValidation("validatepassword", ValidatePassword)

	err := validate.Struct(s)
	if err != nil {
		errs := TranslateError(err, validate)
		return errs
	}
	return nil
}

func TranslateError(err error, validate *validator.Validate) []string {
	var errs []string

	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	// Password validation
	var validatepassword = "validatepassword"
	passwordErrMsg := "{0} must contains at least 1 upper case, numeric, and special character."
	addTranslation(validatepassword, passwordErrMsg, validate, trans)

	en_translations.RegisterDefaultTranslations(validate, trans)

	validatorErrs := translateAll(trans, err)
	for _, e := range validatorErrs {
		translatedErr := e
		errs = append(errs, translatedErr)
	}
	return errs
}

func translateAll(trans ut.Translator, err error) validator.ValidationErrorsTranslations {
	var errs validator.ValidationErrors
	if err != nil {

		// translate all error at once
		errs = err.(validator.ValidationErrors)

	}
	// translations are i18n aware!!!!
	// eg. '10 characters' vs '1 character'
	return errs.Translate(trans)
}

func addTranslation(tag string, errMessage string, validate *validator.Validate, trans ut.Translator) {

	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = validate.RegisterTranslation(tag, trans, registerFn, transFn)
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
