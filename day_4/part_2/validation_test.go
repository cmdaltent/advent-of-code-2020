package part_2

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hairValidation(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		passed bool
	}{
		{name: "ok", input: "#0912af", passed: true},
		{name: "no hash", input: "0000ff", passed: false},
		{name: "too long", input: "#0000ffab", passed: false},
		{name: "too short", input: "#00ff", passed: false},
		{name: "beyond a-f range", input: "#0gh00f", passed: false},
	}

	validate := validator.New()
	require.NoError(t, validate.RegisterValidation("hair", hairValidation))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := struct {
				Field string `validate:"hair"`
			}{
				Field: test.input,
			}
			err := validate.Struct(&is)
			if test.passed {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
