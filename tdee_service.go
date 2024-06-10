package main

import (
	"math"
	"strings"
)

var ActivityLevels = map[string]float64{
	"SEDENTARY":       1.2,
	"LIGHT_ACTIVE":    1.375,
	"MODERATE_ACTIVE": 1.55,
	"ACTIVE":          1.725,
	"VERY_ACTIVE":     1.9,
}

func GetWeightInKg(weightInLbs float64) float64 {
	return math.Floor(weightInLbs / 2.20462)
}

func GetHeightInCm(heightInInches float64) float64 {
	return math.Floor(heightInInches * 2.54)
}

func GetMaleTdee(weightInKg, heightInCm float64, age int, activityLevel, equation string) float64 {
	base := GetMaleBaseCalculation(weightInKg, heightInCm, age, equation)
	return base * ActivityLevels[strings.ToUpper(activityLevel)]
}

func GetFemaleTdee(weightInKg, heightInCm float64, age int, activityLevel, equation string) float64 {
	base := GetFemaleBaseCalculation(weightInKg, heightInCm, age, equation)
	return base * ActivityLevels[strings.ToUpper(activityLevel)]
}

// Harris-Benedict equation for male BMR calculation
// (655 + (9.6 * weight [kg]) + (1.8 * height [cm]) – (4.7 * age [years]))

// Mifflin St. Jeor Male scale for TDEE calculation
// (10*weight [kg]) + (6.25*height [cm]) – (5*age [years]) + 5

func GetMaleBaseCalculation(weightInKg, heightInCm float64, age int, equation string) float64 {
	switch equation {
	case "harris-benedict":
		return math.Floor(655 + (9.6 * weightInKg) + (1.8 * heightInCm) - (4.7 * float64(age)))
	case "mifflin-st-jeor":
	default:
		return math.Floor((10 * weightInKg) + (6.25 * heightInCm) - (5 * float64(age)) + 5)
	}
	return 0
}

// Mifflin St. Jeor Female scale for TDEE calculation
// (10*weight [kg]) + (6.25*height [cm]) – (5*age [years]) – 161

// Harris-Benedict equation for female BMR calculation
// 66 + (13.7 * weight [kg]) + (5 * height [cm]) – (6.8 * age [years])

func GetFemaleBaseCalculation(weightInKg, heightInCm float64, age int, equation string) float64 {
	switch equation {
	case "harris-benedict":
		return math.Floor(66 + (13.7 * weightInKg) + (5 * heightInCm) - (6.8 * float64(age)))
	case "mifflin-st-jeor":
	default:
		return math.Floor((10 * weightInKg) + (6.25 * heightInCm) - (5 * float64(age)) - 161)
	}
	return 0
}
