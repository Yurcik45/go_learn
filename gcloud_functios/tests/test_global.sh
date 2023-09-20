#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{
  "timestamp": 432423432,
  "user_id": 4342432.434234,
  "event_name": "test",
  "event_data": {"test_key": "test_blobal_value"}
 }' https://europe-west1-yuras-event-server-side.cloudfunctions.net/shopify_data
