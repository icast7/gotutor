package main

import "fmt"

func main() {
	crv, _ := getCar("crv")
	bolt, _ := getCar("bolt")

	printDetails(crv)
	printDetails(bolt)
}

func printDetails(c iCar) {
	fmt.Printf("Car: %s", c.getName())
	fmt.Println()
	fmt.Printf("Power: %d", c.getPower())
	fmt.Println()
}
