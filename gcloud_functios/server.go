package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Define your Firestore collection name
const collectionName = "yourCollectionName"

func main() {
	r := gin.Default()

	r.POST("/createTableAndAddRow", createTableAndAddRowHandler)

	http.Handle("/", r)
}

func createTableAndAddRowHandler(c *gin.Context) {
	// Parse the request to get the table fields data
	var requestData YourRequestStruct
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initialize Firebase Admin SDK
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/your/serviceAccountKey.json")

	// logic for parse fields and put it to sheet

	c.JSON(http.StatusOK, gin.H{"message": "Table created and row added"})
}
