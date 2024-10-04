package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()

	validate.RegisterValidation("uniqueEmail", UniqueEmail)
	validate.RegisterValidation("positiveAmount", PositiveAmount)
}

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err != nil {
		errs := map[string]string{}

		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			errs[field] = getErrorMessage(field, tag)
		}

		return errs
	}

	return nil
}

func getErrorMessage(field, tag string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "uniqueEmail":
		return field + " this email is already taken"
	case "positiveAmount":
		return field + " must be greater than 0"
	default:
		return field + " is not valid"
	}
}
