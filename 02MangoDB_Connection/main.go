package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func ConnectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://umapathi1014:f4tLW7LjP7loDjlX@cluster0.sbelj.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB!")
	return client
}

func GetAllStudents(c *fiber.Ctx) error {
	collection := dbClient.Database("Golang_Fiber").Collection("Students")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching students: " + err.Error(),
		})
	}
	defer cursor.Close(ctx)

	var students []bson.M
	if err = cursor.All(ctx, &students); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error decoding students: " + err.Error(),
		})
	}

	return c.JSON(students)
}

func main() {

	dbClient = ConnectMongoDB()
	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
	}()

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber with MongoDB!")
	})

	app.Get("/students", GetAllStudents)

	log.Fatal(app.Listen(":3000"))
}
