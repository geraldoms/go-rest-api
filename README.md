# Go Rest API

This is a simple RESTful API using the mux package for ToDo entities using.

## Requirements
* Go 1.10 or later

## Installation

`$ go get github.com/gorilla/mux`

## Usage 

`$ go run main.go`

## Request samples 

Request:
```bash
curl http://localhost:3000/api/todos
```
Response:
```json
[
    {
        "id":1,
        "todo":"Make some exercises",
        "description":"Walk and run through the park.",
        "isdone":false,
        "createat":"2018-09-01T19:03:38.214599-04:00"
    },{
        "id":2,
        "todo":"Buy fruits and vegetables",
        "description":"They are on sale this week.",
        "isdone":false,
        "createat":"2018-09-01T19:03:38.2146-04:00"
    },{
        "id":3,
        "todo":"Go to Market",
        "description":"Go to market to get some food",
        "isdone":false,
        "createat":"2018-09-01T19:06:55.327237-04:00"
    }
]
```

Request:
```bash
curl -X POST http://localhost:3000/api/todos \
  -d '{
    "todo": "Go to Market",
    "description": "Go to market to get some food",
    "isdone": false
}'
```
Response:
```json
{
    "id":3,
    "todo":"Go to Market",
    "description":"Go to market to get some samples",
    "isdone":false,"createat":"2018-09-01T19:06:55.327237-04:00"
}
```