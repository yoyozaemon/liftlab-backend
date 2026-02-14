package routers

import (
	"liftlab/src/handlers"
	"liftlab/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	SetupHealthRoutes(api)
	SetupAuthRoutes(api)
	SetupUserRoutes(api)
	SetupWorkoutRoutes(api)
}

func SetupHealthRoutes(router fiber.Router) {
	health := router.Group("/health")
	health.Get("/", handlers.HealthCheck)
}

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/login", handlers.Auth.Login)
}

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Get("/get-user/:email", middlewares.VerifyToken, handlers.User.GetUser)
	user.Put("/update-user", middlewares.VerifyToken, handlers.User.UpdateUser)
}

func SetupWorkoutRoutes(router fiber.Router) {
	workout := router.Group("/workout")
	workout.Post("/save-workout", middlewares.VerifyToken, handlers.Workout.SaveWorkout)
	workout.Get("/get-all-workouts/:uid", middlewares.VerifyToken, handlers.Workout.GetAllWorkouts)
	workout.Get("/get-current-day-workouts", middlewares.VerifyToken, handlers.Workout.GetCurrentDayWorkouts)
}
