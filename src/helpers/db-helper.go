package helpers

import (
	"context"
	"liftlab/src/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDB() {
	config.LoadConfig()
	clientOptions := options.Client().ApplyURI(config.MONGODB_URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to MongoDB!")
	Client = client
}
