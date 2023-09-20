package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/bigquery"
	// "github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/option"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

type RequestData struct {
	Timestamp int64       `bigquery:"timestamp" json:"timestamp"`
	UserID    float64     `bigquery:"user_id" json:"user_id"`
	EventName string      `bigquery:"event_name" json:"event_name"`
	EventData interface{} `bigquery:"event_data" json:"event_data"`
}

func insertToBigQuery(data RequestData) error {
	fmt.Println("======= inserg BIG QUERY LOGS =======")

	projectID := "yuras-event-server-side"
	datasetID := "shopify_pixel"
	tableID := "hanitqftest_test"

	// serviceAccountKeyJSON := os.Getenv("SERVICE_ACCOUNT_KEY_JSON")
	fmt.Println("Data to be inserted:", data)

	// fmt.Println("Service Account Key JSON:", serviceAccountKeyJSON)

	// opt := option.WithCredentialsJSON([]byte(serviceAccountKeyJSON))
	opt := option.WithCredentialsFile("yuras-event-server-side-d47244471b44.json")
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID, opt)
	if err != nil {
		return err
	}
	defer client.Close()

	// Define the BigQuery schema to match your structure.
	schema := bigquery.Schema{
		{Name: "timestamp", Type: bigquery.TimestampFieldType},
		{Name: "user_id", Type: bigquery.StringFieldType},
		{Name: "event_name", Type: bigquery.StringFieldType},
		{Name: "event_data", Type: bigquery.JSONFieldType},
	}

	// Create a BigQuery dataset reference.
	datasetRef := client.Dataset(datasetID)

	// Create a BigQuery table reference.
	tableRef := datasetRef.Table(tableID)

	fmt.Println("Table Reference:", tableRef)
	fmt.Println("Schema:", schema)

	eventDataJSON, err := json.Marshal(data.EventData)
	if err != nil {
		return err
	}

	// Create a BigQuery struct to insert data into the table.
	rowsToInsert := []*bigquery.StructSaver{
		{
			Schema:   schema,
			InsertID: fmt.Sprintf("%d", time.Now().UnixNano()), // Unique insert ID, you can use a timestamp.
			Struct:   &data,
		},
	}

	rowsToInsert[0].Struct.(*RequestData).EventData = string(eventDataJSON)

	// Perform the insert operation.
	if err := tableRef.Inserter().Put(ctx, rowsToInsert); err != nil {
		return err
	}

	return nil
}

func shopifyDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	var data RequestData

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("JSON PARSING ERR: ", err.Error())
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}

	// Insert data into BigQuery.
	if err := insertToBigQuery(data); err != nil {
		fmt.Println("Error inserting data into BigQuery:", err.Error())
		fmt.Println("============================================")

		http.Error(w, "Error inserting data into BigQuery", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "Data received successfully")

}

func main() {
	// functions.HTTP("HelloHTTP", shopifyDataHandler)
	run_local()
}

func run_local() {
	// Register the function to handle HTTP requests.
	funcframework.RegisterHTTPFunction("/", shopifyDataHandler)

	// Specify the custom port (e.g., 8080) on which the function will run locally.
	port := "8080"

	// Start the Functions Framework with the specified port.
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("Error starting Functions Framework: %v", err)
	}
}
