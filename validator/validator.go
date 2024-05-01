package validator

import (
	"github.com/AliSahib998/QuotesAssesments/errhandler"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

var validate = validator.New()

func Validation(data interface{}) error {

	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")
	_ = entranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(data)

	var errorList = make([]errhandler.ValidationError, 0)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			var validationErr = errhandler.ValidationError{
				Path:    e.Field(),
				Message: e.Translate(trans),
			}
			errorList = append(errorList, validationErr)
		}

	}

	if len(errorList) > 0 {
		return errhandler.NewBadRequestError("invalid_request", errorList)
	}
	return nil
}
