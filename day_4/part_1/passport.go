package part_1

import (
	"github.com/go-playground/validator/v10"

	"github.com/cmdaltent/advent-of-code-2020/day_4/common"
)

func GeneratePassport() common.Passport {
	return &passport{}
}

type passport struct {
	BirthYear      string `validate:"required"`
	IssueYear      string `validate:"required"`
	ExpirationYear string `validate:"required"`
	Height         string `validate:"required"`
	HairColor      string `validate:"required"`
	EyeColor       string `validate:"required"`
	ID             string `validate:"required"`
	CountryID      string
}

var validate = validator.New()

func (p *passport) IsValid() bool {
	return validate.Struct(p) == nil
}

func (p *passport) MergeWith(key, value string) {
	switch key {
	case "byr":
		p.BirthYear = value
	case "iyr":
		p.IssueYear = value
	case "eyr":
		p.ExpirationYear = value
	case "hgt":
		p.Height = value
	case "hcl":
		p.HairColor = value
	case "ecl":
		p.EyeColor = value
	case "pid":
		p.ID = value
	case "cid":
		p.CountryID = value
	}
}
