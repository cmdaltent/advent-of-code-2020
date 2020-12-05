package part_2

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/cmdaltent/advent-of-code-2020/day_4/common"
)

func GeneratePassport() common.Passport {
	return &passport{}
}

type passport struct {
	BirthYear      int    `validate:"required,min=1920,max=2002"`
	IssueYear      int    `validate:"required,min=2010,max=2020"`
	ExpirationYear int    `validate:"required,min=2020,max=2030"`
	Height         string `validate:"required,height"`
	HairColor      string `validate:"required,hair"`
	EyeColor       string `validate:"required,oneof=amb blu brn gry grn hzl oth"`
	ID             string `validate:"required,len=9"`
	CountryID      string
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("height", heightValidator); err != nil {
		log.Fatalln(err)
	}
	if err := validate.RegisterValidation("hair", hairValidation); err != nil {
		log.Fatalln(err)
	}
}

func (p *passport) IsValid() bool {
	return validate.Struct(p) == nil
}

func (p *passport) MergeWith(key, value string) {
	switch key {
	case "byr":
		p.BirthYear, _ = strconv.Atoi(value)
	case "iyr":
		p.IssueYear, _ = strconv.Atoi(value)
	case "eyr":
		p.ExpirationYear, _ = strconv.Atoi(value)
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
