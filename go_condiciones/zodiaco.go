package main

import "fmt"

func main() {
	var day, month int
	var zodiacSign string

	// Read day from user input
	fmt.Scanf("%d ", &day)

	// Read month from user input
	fmt.Scanf("%d", &month)

	// Switch for each month
	switch month {
	case 1:
		if day < 20 {
			zodiacSign = "Capricorn"
		} else {
			zodiacSign = "Aquarius"
		}
		break
	case 2:
		if day < 18 {
			zodiacSign = "Aquarius"
		} else {
			zodiacSign = "Pisces"
		}
		break
	case 3:
		if day < 21 {
			zodiacSign = "Pisces"
		} else {
			zodiacSign = "Aries"
		}
		break
	case 4:
		if day < 20 {
			zodiacSign = "Aries"
		} else {
			zodiacSign = "Taurus"
		}
		break
	case 5:
		if day < 21 {
			zodiacSign = "Taurus"
		} else {
			zodiacSign = "Gemini"
		}
		break
	case 6:
		if day < 21 {
			zodiacSign = "Gemini"
		} else {
			zodiacSign = "Cancer"
		}
		break
	case 7:
		if day < 23 {
			zodiacSign = "Cancer"
		} else {
			zodiacSign = "Leo"
		}
		break
	case 8:
		if day < 23 {
			zodiacSign = "Leo"
		} else {
			zodiacSign = "Virgo"
		}
		break
	case 9:
		if day < 23 {
			zodiacSign = "Virgo"
		} else {
			zodiacSign = "Libra"
		}
		break
	case 10:
		if day < 23 {
			zodiacSign = "Libra"
		} else {
			zodiacSign = "Scorpio"
		}
		break
	case 11:
		if day < 22 {
			zodiacSign = "Scorpio"
		} else {
			zodiacSign = "Sagittarius"
		}
		break
	case 12:
		if day < 22 {
			zodiacSign = "Sagittarius"
		} else {
			zodiacSign = "Capricorn"
		}
		break
	default:
		zodiacSign = "Unknown"
		break
	}
	fmt.Print(zodiacSign)
}
