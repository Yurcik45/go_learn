package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/sheets/v4"
)

type RequestData struct {
	Timestamp int64       `json:"timestamp"`
	UserID    float64     `json:"user_id"`
	EventName string      `json:"event_name"`
	EventData interface{} `json:"event_data"`
}

func extractTableName(spreadsheetID string) string {
	parts := strings.Split(spreadsheetID, ".")
	if len(parts) >= 3 {
		return parts[2]
	}
	return ""
}

func getNextAvailableRow(client *sheets.Service, ctx context.Context, spreadsheetID string) int {
	resp, err := client.Spreadsheets.Values.Get(spreadsheetID, "Sheet1").Context(ctx).Do()
	if err != nil {
		return 1
	}
	lastRow := len(resp.Values)
	return lastRow + 1
}

func write_to_table(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Initialize Google Sheets API client (credentials are provided automatically)
	client, err := sheets.NewService(ctx)
	if err != nil {
		fmt.Println("Error writing data to Google Sheets:", err.Error())
		http.Error(w, "Error initializing Sheets client", http.StatusInternalServerError)
		return
	}

	// Specify the spreadsheet ID where you want to write data
	spreadsheetID := "yuras-event-server-side.shopify_pixel.hanitqftest_test"

	nextRow := getNextAvailableRow(client, ctx, spreadsheetID)

	tableName := extractTableName(spreadsheetID)

	// Specify the range where to last Row)
	writeRange := fmt.Sprintf("%s!A%d", tableName, nextRow)

	// Parse the request body into a RequestData struct
	var requestData RequestData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	} else {
		fmt.Println(w, "Data received successfully")
	}

	fmt.Println("Data EventData: ", requestData.EventData)
	fmt.Println("Data UserID: ", requestData.UserID)
	fmt.Println("Data EventName: ", requestData.EventName)
	fmt.Println("Data Timestamp: ", requestData.Timestamp)

	// Create a value range to write
	values := []interface{}{
		requestData.Timestamp,
		requestData.UserID,
		requestData.EventName,
		requestData.EventData,
	}

	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{values},
	}

	fmt.Println("Data values: ", values)
	fmt.Println("======================")
	fmt.Println("spreadsheetID: ", spreadsheetID)
	fmt.Println("nextRow: ", nextRow)
	fmt.Println("tableName: ", tableName)
	fmt.Println("writeRange: ", writeRange)
	fmt.Println("======================")

	fmt.Println("+++++++++++++++++++")
	// Write the data to Google Sheets
	_, err = client.Spreadsheets.Values.Update(spreadsheetID, writeRange, valueRange).
		ValueInputOption("RAW").
		Context(ctx).
		Do()
	if err != nil {
		fmt.Println("Error writing data to Google Sheets:", err.Error())
		http.Error(w, "Error writing data to Google Sheets", http.StatusInternalServerError)
		return
	}
	fmt.Println("+++++++++++++++++++")

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Data written to Google Sheets successfully")
}

func shopify_data_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Error reading request body", http.StatusBadRequest)
	// }

	// defer r.Body.Close()

	// var data RequestData

	// if err := json.Unmarshal(body, &data); err != nil {
	// 	http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
	// 	return
	// }

	write_to_table(w, r)

}

func init() {
	functions.HTTP("HelloHTTP", shopify_data_handler)
}
