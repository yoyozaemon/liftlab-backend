package handlers

import (
	auth "liftlab/src/handlers/Auth"
	users "liftlab/src/handlers/Users"
	workouts "liftlab/src/handlers/Workouts"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}
type UserHandler struct{}
type WorkoutHandler struct{}

var Auth = new(AuthHandler)
var User = new(UserHandler)
var Workout = new(WorkoutHandler)

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	return auth.Login(c)
}

func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	return users.GetUser(c)
}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	return users.UpdateUser(c)
}

func (w *WorkoutHandler) SaveWorkout(c *fiber.Ctx) error {
	return workouts.SaveWorkout(c)
}

func (w *WorkoutHandler) GetAllWorkouts(c *fiber.Ctx) error {
	return workouts.GetAllWorkouts(c)
}

func (w *WorkoutHandler) GetCurrentDayWorkouts(c *fiber.Ctx) error {
	return workouts.GetCurrentDayWorkouts(c)
}
