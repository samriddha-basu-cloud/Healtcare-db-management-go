package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI   = "mongodb://localhost:27017" // Update with your MongoDB URI
	dbName     = "myhealthdb"                // Update with your database name
	colName    = "patients"                  // Collection name
	collection *mongo.Collection
)

var client *mongo.Client
var ctx context.Context

type Patient struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	DOB       time.Time          `json:"dob" bson:"dob"`
	Condition string             `json:"condition" bson:"condition"`
}

func init() {
	// Initialize MongoDB client
	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the context for database operations
	ctx = context.TODO()

	// Select the MongoDB collection
	collection = client.Database(dbName).Collection(colName)

	// Check the connection status
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	} else {
		fmt.Println("Connected to MongoDB successfully!")
	}
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Define API routes
	r.GET("/patients", getPatients)
	r.GET("/patients/:id", getPatientByID)
	r.POST("/patients", addPatient)
	r.PUT("/patients/:id", updatePatient)
	r.DELETE("/patients/:id", deletePatientByID)
	r.DELETE("/patients", deleteAllPatients)

	// Run the HTTP server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func getPatients(c *gin.Context) {
	// Retrieve all patients from MongoDB and return them as JSON
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var patients []Patient
	for cursor.Next(ctx) {
		var patient Patient
		if err := cursor.Decode(&patient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		patients = append(patients, patient)
	}

	c.JSON(http.StatusOK, patients)
}

func getPatientByID(c *gin.Context) {
	// Retrieve a patient by ID from MongoDB and return it as JSON
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient Patient
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&patient)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func addPatient(c *gin.Context) {
	// Add a new patient to MongoDB with an ascending ID and DOB as time.Time
	var patient Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a custom ascending ID using a Unix timestamp
	patient.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	insertResult, err := collection.InsertOne(ctx, patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, insertResult.InsertedID)
}

func updatePatient(c *gin.Context) {
	// Update a patient in MongoDB
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": patient}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Patient with ID %s updated", id))
}

func deletePatientByID(c *gin.Context) {
	// Delete a patient by ID from MongoDB
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Patient with ID %s deleted", id))
}

func deleteAllPatients(c *gin.Context) {
	// Delete all patients from MongoDB
	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "All patients deleted")
}
