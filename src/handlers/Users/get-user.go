package users

import (
	"liftlab/src/config"
	"liftlab/src/helpers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(c *fiber.Ctx) error {
	email := c.Params("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "Email is required"})
	}

	collection := helpers.Client.Database(config.DB_NAME).Collection("users")
	var user fiber.Map
	projection := bson.M{"_id": 0, "created_at": 0}
	err := collection.FindOne(c.Context(), bson.D{{Key: "email", Value: email}}, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": "User not found"})
		}
		helpers.Logger.Printf("Database error in GetUser: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "Internal server error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "user": user})
}
