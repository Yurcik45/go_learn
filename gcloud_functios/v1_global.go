package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/option"
)

type RequestData struct {
	Timestamp int64       `bigquery:"timestamp" json:"timestamp"`
	UserID    float64     `bigquery:"user_id" json:"user_id"`
	EventName string      `bigquery:"event_name" json:"event_name"`
	EventData interface{} `bigquery:"event_data" json:"event_data"`
}

func insertToBigQuery(data RequestData) error {
	projectID := "yuras-event-server-side"
	datasetID := "shopify_pixel"
	tableID := "hanitqftest_test"
	serviceAccountKeyJSON := os.Getenv("SERVICE_ACCOUNT_KEY_JSON")
	opt := option.WithCredentialsJSON([]byte(serviceAccountKeyJSON))
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID, opt)
	if err != nil {
		return err
	}

	defer client.Close()

	schema := bigquery.Schema{
		{Name: "timestamp", Type: bigquery.TimestampFieldType},
		{Name: "user_id", Type: bigquery.StringFieldType},
		{Name: "event_name", Type: bigquery.StringFieldType},
		{Name: "event_data", Type: bigquery.StringFieldType},
	}

	datasetRef := client.Dataset(datasetID)
	tableRef := datasetRef.Table(tableID)

	eventDataJSON, err := json.Marshal(data.EventData)
	if err != nil {
		return err
	}

	rowsToInsert := []*bigquery.StructSaver{
		{
			Schema:   schema,
			InsertID: fmt.Sprintf("%d", time.Now().UnixNano()),
			Struct:   &data,
		},
	}

	rowsToInsert[0].Struct.(*RequestData).EventData = string(eventDataJSON)

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
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}

	if err := insertToBigQuery(data); err != nil {
		http.Error(w, "Error inserting data into BigQuery", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "Data received successfully")
}

func init() {
	functions.HTTP("HelloHTTP", shopifyDataHandler)
}
