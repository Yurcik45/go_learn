#!/bin/sh

is_global=$1

url="localhost:8080"

if [ -n "$is_global" ]; then
  url="https://europe-west1-yuras-event-server-side.cloudfunctions.net/shopify_data"
fi

json_file="body.json"

if [ -f "$json_file" ]; then
    json_data=$(cat "$json_file")
else
    json_data='{
      "timestamp": 1234567890,
      "user_id": 987654321.123456,
      "event_name": "default_event",
      "event_data": {"default_key": "default_value"}
    }'
    echo $json_data > "$json_file"
fi

curl -X POST -H "Content-Type: application/json" -d "$json_data" $url
