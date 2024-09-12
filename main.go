package main

import "fmt"

func main() {

	var i int
	var temperature float64

	fmt.Println("Converting Between Temperature Unit Types")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("1. C --> R (Reaumur)")
	fmt.Println("2. C --> F (Fahrenheit")
	fmt.Println("3. C --> K (Kelvin)")
	fmt.Println("Enter your choice (1/2/3) : ")
	fmt.Scan(&i) //asking to make a choice
	fmt.Println("Enter the temperature in Celsius : ")
	fmt.Scan(&temperature) //enter temperature
	switch i {
	case 1:
		reau := temperature * 0.8
		fmt.Printf("Temperature in R : %.2f", reau)
	case 2:
		fahr := (temperature * 9 / 5) + 32
		fmt.Printf("Temperature in F : %.2f", fahr)
	case 3:
		kelv := temperature + 273.15
		fmt.Printf("Temperature in K : %.2f", kelv)
	default:
		fmt.Println("Invalid. Please select 1,2, or 3")
	}
}
