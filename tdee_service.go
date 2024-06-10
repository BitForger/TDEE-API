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

// Mifflin St. Jeor Male scale for TDEE calculation
// (10*weight [kg]) + (6.25*height [cm]) – (5*age [years]) + 5

func GetMaleBaseCalculation(weightInKg, heightInCm float64, age int) float64 {
	return math.Floor((10 * weightInKg) + (6.25 * heightInCm) - (5 * float64(age)) + 5)
}

// Mifflin St. Jeor Female scale for TDEE calculation
// (10*weight [kg]) + (6.25*height [cm]) – (5*age [years]) – 161

func GetFemaleBaseCalculation(weightInKg, heightInCm float64, age int) float64 {
	return math.Floor((10 * weightInKg) + (6.25 * heightInCm) - (5 * float64(age)) - 161)
}

func GetWeightInKg(weightInLbs float64) float64 {
	return math.Floor(weightInLbs / 2.20462)
}

func GetHeightInCm(heightInInches float64) float64 {
	return math.Floor(heightInInches * 2.54)
}

func GetMaleTdee(weightInKg, heightInCm float64, age int, activityLevel string) float64 {
	base := GetMaleBaseCalculation(weightInKg, heightInCm, age)
	return base * ActivityLevels[strings.ToUpper(activityLevel)]
}

func GetFemaleTdee(weightInKg, heightInCm float64, age int, activityLevel string) float64 {
	base := GetFemaleBaseCalculation(weightInKg, heightInCm, age)
	return base * ActivityLevels[strings.ToUpper(activityLevel)]
}
