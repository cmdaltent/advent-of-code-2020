package part_2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func heightValidator(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	if strings.HasSuffix(fieldValue, "cm") {
		raw := strings.TrimRight(fieldValue, "cm")
		height, err := strconv.Atoi(raw)
		if err != nil {
			return false
		}
		return height >= 150 && height <= 193
	}
	if strings.HasSuffix(fieldValue, "in") {
		raw := strings.TrimRight(fieldValue, "in")
		height, err := strconv.Atoi(raw)
		if err != nil {
			return false
		}
		return height >= 59 && height <= 76
	}
	return false
}

func hairValidation(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, fieldValue)
	return matched
}
