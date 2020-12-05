package main

import "log"

func main() {
	log.Printf("First Part")
	checkInput(applyRentalShopPolicy)

	log.Println("Second Part")
	checkInput(applyTobogganCorporatePolicy)
}
