package users

import (
	"liftlab/src/config"
	"liftlab/src/helpers"
	"liftlab/src/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateUserBody struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
}
type UpdateUserRequest struct {
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Body  UpdateUserBody `json:"body"`
}

func UpdateUser(c *fiber.Ctx) error {
	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "Invalid request body"})
	}

	collection := helpers.Client.Database(config.DB_NAME).Collection("users")
	var user models.User

	err := collection.FindOne(c.Context(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": "User not found"})
	}

	user.Body.Height = req.Body.Height
	user.Body.Weight = req.Body.Weight

	_, err = collection.UpdateOne(c.Context(), bson.M{"email": req.Email}, bson.M{"$set": user})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "User updated successfully"})

}
