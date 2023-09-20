#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{
  "timestamp": 432423432,
  "user_id": 4342432.434234,
  "event_name": "test",
  "event_data": {"test_key": "test_value"}
}' localhost:8080
