package domains

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func CustomValidationError(err error) map[string]string {
	errs := make(map[string]string)

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			fieldName := fieldErr.Field()
			switch fieldErr.Tag() {
			case "required":
				errs[fieldName] = fmt.Sprintf("%s is required", fieldName)
			case "email":
				errs[fieldName] = fmt.Sprintf("%s must be a valid email address", fieldName)
			case "min":
				errs[fieldName] = fmt.Sprintf("%s must be at least %s characters long", fieldName, fieldErr.Param())
			case "max":
				errs[fieldName] = fmt.Sprintf("%s must be at most %s characters long", fieldName, fieldErr.Param())
			default:
				errs[fieldName] = fmt.Sprintf("%s is invalid", fieldName)
			}
		}
	}
	return errs
}
