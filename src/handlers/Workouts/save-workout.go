package workouts

import (
	"liftlab/src/config"
	"liftlab/src/helpers"
	"liftlab/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	collection := helpers.Client.Database(config.DB_NAME).Collection("workouts")

	filter := bson.M{"uid": req.UID}
	update := bson.M{
		"$set": bson.M{
			"name":       req.Name,
			"exercises":  req.Exercises,
			"cardio":     req.Cardio,
			"updated_at": time.Now(),
		},
		"$setOnInsert": bson.M{
			"created_at": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(c.Context(), filter, update, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "Failed to save workout"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Workout saved successfully"})
}
