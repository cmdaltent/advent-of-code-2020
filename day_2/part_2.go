package main

func applyTobogganCorporatePolicy(password string, first, second int, char string) bool {
	n1 := string(password[first-1]) == char
	n2 := string(password[second-1]) == char

	return (n1 == true && n2 == false) || (n1 == false && n2 == true)
}
