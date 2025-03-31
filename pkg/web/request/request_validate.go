package request

import (
	"errors"

	"github.com/gmalheiro/playground-golang-clean-arch/pkg/web/response"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func validateJSON(object any) []response.Validation {
	if err := validate.Struct(object); err != nil {
		var validateErrs validator.ValidationErrors
		fieldErrors := make(map[string][]string)

		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				fieldName := e.Field()
				fieldError := fieldName + " " + e.Tag()

				fieldErrors[fieldName] = append(fieldErrors[fieldName], fieldError)
			}
		}

		validates := make([]response.Validation, 0, len(fieldErrors))

		for k, v := range fieldErrors {
			validates = append(validates, response.Validation{Field: k, Messages: v})
		}

		return validates
	}

	return nil
}
