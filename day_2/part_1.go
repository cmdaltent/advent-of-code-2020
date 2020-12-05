package main

import (
	"strings"
)

func applyRentalShopPolicy(pwd string, min, max int, char string) bool {
	c := strings.Count(pwd, char)
	return c >= min && c <= max
}
