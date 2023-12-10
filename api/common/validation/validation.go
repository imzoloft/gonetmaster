package validation

import (
	"github.com/go-playground/validator/v10"
)

type errorValidation struct {
	failedField string
	tag         string
}

func (err *errorValidation) Error() string {
	return err.failedField + " " + err.tag
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(data interface{}) []error {
	var errors []error
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorValidation

			element.failedField = err.StructNamespace()
			element.tag = err.Tag()

			errors = append(errors, &element)
		}
	}
	return errors
}
