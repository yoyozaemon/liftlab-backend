package auth

import (
	"liftlab/src/config"
	"liftlab/src/helpers"
	"liftlab/src/models"
	"liftlab/src/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoUrl string `json:"photo_url"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email is required"})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	collection := helpers.Client.Database(config.DB_NAME).Collection("users")
	var user models.User

	err := collection.FindOne(c.Context(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			uid, err := utils.GenerateUID()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate UID"})
			}

			newUser := models.User{
				UID:      uid,
				Name:     req.Name,
				Email:    req.Email,
				PhotoUrl: req.PhotoUrl,
				BaseModel: models.BaseModel{
					ID:        primitive.NewObjectID(),
					CreatedAt: time.Now(),
				},
			}

			_, err = collection.InsertOne(c.Context(), newUser)
			if err != nil {
				helpers.Logger.Printf("Failed to insert user: %v", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
			}
			user = newUser
		} else {
			helpers.Logger.Printf("Database error during login: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}
	}

	token, err := utils.GenerateToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
