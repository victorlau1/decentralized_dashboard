package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
}

func TestClientDecentralizationRequiredFields(t *testing.T) {
	dc := &ClientDecentralization{}
	errs := validate.Struct(dc)
	for _, err := range errs.(validator.ValidationErrors) {
		assert.Equal(t, "has error", err)
	}
}
