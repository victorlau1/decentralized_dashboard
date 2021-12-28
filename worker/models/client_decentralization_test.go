package models

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var validate *validator.Validate

func TestClientDecentralizationRequiredFields(t *testing.T) {
	validate = validator.New()
	dc := &ClientDecentralization{}
	errs := validate.Struct(dc)
	found := []string{}
	for _, err := range errs.(validator.ValidationErrors) {
		found = append(found, err.Field())
	}

	assert.Equal(t, found, []string{"Country", "Region", "Blockchain", "Timestamp", "Client"})
}
