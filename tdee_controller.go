package main

import (
	"github.com/gofiber/fiber/v2"
	"math"
	"strings"
)

func HandleTdeeDaily(ctx *fiber.Ctx) error {
	// Params
	// "weight" - float64
	// "height" - float64
	// "age" - int
	// "sex" - string
	// "activity_level" - string

	weightParam := ctx.QueryFloat("weight")
	heightParam := ctx.QueryFloat("height")
	ageParam := ctx.QueryInt("age")
	sexParam := ctx.Query("sex")
	activityLevelParam := ctx.Query("activity_level")

	if ctx.Queries() == nil {
		return ctx.Status(400).SendString("Missing query parameters")
	}
	if _, ok := ActivityLevels[strings.ToUpper(activityLevelParam)]; !ok {
		return ctx.Status(400).SendString("Invalid activity level")
	}
	if sexParam != "male" && sexParam != "female" {
		return ctx.Status(400).SendString("Invalid sex")
	}

	var tdee float64
	switch sexParam {
	case "male":
		tdee = GetMaleTdee(GetWeightInKg(weightParam), GetHeightInCm(heightParam), ageParam, activityLevelParam)

	case "female":
		tdee = GetFemaleTdee(GetWeightInKg(weightParam), GetHeightInCm(heightParam), ageParam, activityLevelParam)
	}

	return ctx.JSON(fiber.Map{
		"maintenance_calories": math.Floor(tdee),
		"deficits": fiber.Map{
			"deficit_250":  math.Floor(tdee - 250),
			"deficit_500":  math.Floor(tdee - 500),
			"deficit_750":  math.Floor(tdee - 750),
			"deficit_1000": math.Floor(tdee - 1000),
		},
	})
}
