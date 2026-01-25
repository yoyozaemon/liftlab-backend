package users

import (
	"context"
	"liftlab/src/helpers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(c *fiber.Ctx) error {
	email := c.Params("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email is required"})
	}

	collection := helpers.Client.Database("liftlab").Collection("users")
	var user fiber.Map
	projection := bson.M{"_id": 0}
	err := collection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
}
