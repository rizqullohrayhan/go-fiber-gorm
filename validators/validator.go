package validators

import (
    "github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateStruct(data interface{}) map[string]string {
    err := Validate.Struct(data)
    if err == nil {
        return nil
    }

    errors := err.(validator.ValidationErrors)
    errorMessages := make(map[string]string)
    for _, e := range errors {
        errorMessages[e.Field()] = e.Tag()
    }
    return errorMessages
}