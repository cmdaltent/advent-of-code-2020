package main

import "github.com/go-playground/validator/v10"

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

func (p *passport) isValid() bool {
	return validate.Struct(p) == nil
}

func (p *passport) mergeWith(key, value string) {
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
