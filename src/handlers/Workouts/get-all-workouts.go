package workouts

import (
	"liftlab/src/config"
	"liftlab/src/helpers"
	"liftlab/src/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllWorkouts(c *fiber.Ctx) error {
	uid := c.Params("uid")

	collection := helpers.Client.Database(config.DB_NAME).Collection("workouts")
	var workout []models.Workout

	cursor, err := collection.Find(c.Context(), bson.M{"uid": uid})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": "Workout not found"})
	}

	cursor.All(c.Context(), &workout)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Workout fetched successfully", "data": workout})
}
