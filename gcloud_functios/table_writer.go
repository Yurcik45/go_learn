package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// Load your service account key JSON file
	keyFile := "path/to/your/serviceAccountKey.json"

	// Initialize Google Sheets API client
	ctx := context.Background()
	client, err := sheets.NewService(ctx, option.WithCredentialsFile(keyFile))
	if err != nil {
		log.Fatalf("Error initializing Sheets client: %v", err)
	}

	// Specify the spreadsheet ID (or title) where you want to add data
	spreadsheetID := "your-spreadsheet-id"

	// Define the data you want to add
	data := []interface{}{"Value1", "Value2", "Value3"}

	// Specify the range where you want to add the data (e.g., "Sheet1!A1:C1")
	rangeToWrite := "Sheet1!A1:C1"

	// Create the ValueRange with your data
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{data},
	}

	// Write data to the Google Sheet
	_, err = client.Spreadsheets.Values.Update(spreadsheetID, rangeToWrite, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Error writing data to Google Sheet: %v", err)
	}

	fmt.Println("Data added to Google Sheet successfully!")
}
