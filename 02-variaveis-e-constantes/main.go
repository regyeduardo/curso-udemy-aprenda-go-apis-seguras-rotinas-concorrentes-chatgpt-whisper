package main

import "fmt"

func main() {
	var name string = "Drogo"
	var anotherName = "Carl"
	oneMoreName := "Mongo"

	const firstConstName = "Aang"

	fmt.Println(name)
	fmt.Println(anotherName)
	fmt.Println(oneMoreName)
	fmt.Println(firstConstName)

	var age int = 43
	var ageEightBits int8 = 43
	var ageSixteenBits int16 = 43
	var ageThirtTwoBits int32 = 43
	var ageSixtyFourBits int32 = 43

	// Natural numbers
	var ageUIntEightBits uint8 = 1
	var ageUIntSixteenBits uint16 = 1
	var ageUIntThirtTwoBits uint32 = 1
	var ageUIntSixtyFourBits uint32 = 1

	var pi float32 = 3.14
	var piSixtyFourBits float64 = 3.14

	fmt.Println("age", age)
	fmt.Println("ageEightBits", ageEightBits)
	fmt.Println("ageSixteenBits", ageSixteenBits)
	fmt.Println("ageThirtTwoBits", ageThirtTwoBits)
	fmt.Println("ageSixtyFourBits", ageSixtyFourBits)
	fmt.Println("ageUIntEightBits", ageUIntEightBits)
	fmt.Println("ageUIntSixteenBits", ageUIntSixteenBits)
	fmt.Println("ageUIntThirtTwoBits", ageUIntThirtTwoBits)
	fmt.Println("ageUIntSixtyFourBits", ageUIntSixtyFourBits)
	fmt.Println("pi", pi)
	fmt.Println("piSixtyFourBits", piSixtyFourBits)

	var isEnable bool = true

	fmt.Println(isEnable)

	var a, b, c, d = "a", 50, true, 6.34

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
