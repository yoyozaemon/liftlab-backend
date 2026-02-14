package workouts

import (
	"encoding/json"
	"fmt"
	"liftlab/src/config"
	"liftlab/src/helpers"
	"liftlab/src/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type SaveWorkoutRequest struct {
	UID       string                   `json:"uid" bson:"uid"`
	Name      string                   `json:"name" bson:"name"`
	Exercises []models.Weight_Exercise `json:"exercises" bson:"exercises"`
	Cardio    []models.Cardio_Exercise `json:"cardio" bson:"cardio"`
}

func SaveWorkout(c *fiber.Ctx) error {
	var req SaveWorkoutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "Invalid request body"})
	}
	fmt.Print(json.Unmarshal(c.Body(), &req))

	collection := helpers.Client.Database(config.DB_NAME).Collection("workouts")
	var workout models.Workout

	err := collection.FindOne(c.Context(), bson.M{"uid": req.UID}).Decode(&workout)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": "Workout not found"})
	}

	workout.Name = req.Name
	workout.Exercises = req.Exercises
	workout.Cardio = req.Cardio

	_, err = collection.UpdateOne(c.Context(), bson.M{"uid": req.UID}, bson.M{"$set": workout})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "Failed to update workout"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Workout updated successfully"})
}
