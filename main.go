package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Receipt struct {
	Id           uuid.UUID
	Retailer     string
	PurchaseDate time.Time
	PurchaseTime time.Time
	Items        Item
	Total        string
}

type Item struct {
	ShortDescription string
	Price            string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /receipts/process", handleReceiptProcess)
	mux.HandleFunc("GET /receipts/{id}/points", handleReceiptPoints)

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

func handleReceiptProcess(w http.ResponseWriter, r *http.Request) {
	/*
		* Path: `/receipts/process`
		* Method: `POST`
		* Payload: Receipt JSON
		* Response: JSON containing an id for the receipt.

		Description:

		Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID generated by your code.

		The ID returned is the ID that should be passed into `/receipts/{id}/points` to get the number of points the receipt
		was awarded.

		How many points should be earned are defined by the rules below.

		Reminder: Data does not need to survive an application restart. This is to allow you to use in-memory solutions to track any data generated by this endpoint.

		Example Response:
		```json
		{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
		```
	*/
}

func handleReceiptPoints(w http.ResponseWriter, r *http.Request) {
	/*
		* Path: `/receipts/{id}/points`
		* Method: `GET`
		* Response: A JSON object containing the number of points awarded.

		A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

		Example Response:
		```json
		{ "points": 32 }
		```
	*/
}

func calculatePoints(receipt Receipt) int {
	/*
		## Examples

		```json
		{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
		}
		```
		```text
		Total Points: 28
		Breakdown:
			6 points - retailer name has 6 characters
			10 points - 5 items (2 pairs @ 5 points each)
			3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
						item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
			3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
						item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
			6 points - purchase day is odd
		+ ---------
		= 28 points
		```

		----

		```json
		{
		"retailer": "M&M Corner Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items": [
			{
			"shortDescription": "Gatorade",
			"price": "2.25"
			},{
			"shortDescription": "Gatorade",
			"price": "2.25"
			},{
			"shortDescription": "Gatorade",
			"price": "2.25"
			},{
			"shortDescription": "Gatorade",
			"price": "2.25"
			}
		],
		"total": "9.00"
		}
		```
		```text
		Total Points: 109
		Breakdown:
			50 points - total is a round dollar amount
			25 points - total is a multiple of 0.25
			14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
						note: '&' is not alphanumeric
			10 points - 2:33pm is between 2:00pm and 4:00pm
			10 points - 4 items (2 pairs @ 5 points each)
		+ ---------
		= 109 points
		```
	*/
	points := 0
	return points
}
