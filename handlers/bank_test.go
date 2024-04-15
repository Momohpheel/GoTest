package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-interview/database"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {

	type payloadStruct struct {
		AccountID uint   `json:"account_id"`
		Reference string `json:"reference"`
		Amount    uint   `json:"amount"`
	}

	type testPayload struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		payload      payloadStruct
		token        string
	}

	test := testPayload{
		description:  "Test: create transaction, get http status 201",
		route:        "/create-transaction",
		expectedCode: 201,
		payload: payloadStruct{
			AccountID: 1,
			Reference: "wie1234",
			Amount:    200000,
		},
	}

	// Define Fiber app.
	app := fiber.New()
	database.Start()

	app.Post("/create-transaction", CreateTransaction)

	testItem := test

	// Create a new http request with the route from the test case
	payload, err := json.Marshal(testItem.payload)
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodPost, testItem.route, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request plain with the app,
	resp, err := app.Test(req, -1)
	if err != nil {
		log.Println(err)
	}

	// Verify, if the status code is as expected
	assert.Equalf(t, testItem.expectedCode, resp.StatusCode, testItem.description)

}
