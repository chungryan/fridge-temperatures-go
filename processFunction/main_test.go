package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("Example from specs", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Body: "[{\"id\":\"a\",\"timestamp\":1509493641,\"temperature\":3.53},{\"id\":\"b\",\"timestamp\":1509493642,\"temperature\":4.13},{\"id\":\"c\",\"timestamp\":1509493643,\"temperature\":3.96},{\"id\":\"a\",\"timestamp\":1509493644,\"temperature\":3.63},{\"id\":\"c\",\"timestamp\":1509493645,\"temperature\":3.96},{\"id\":\"a\",\"timestamp\":1509493645,\"temperature\":4.63},{\"id\":\"a\",\"timestamp\":1509493646,\"temperature\":3.53},{\"id\":\"b\",\"timestamp\":1509493647,\"temperature\":4.15},{\"id\":\"c\",\"timestamp\":1509493655,\"temperature\":3.95},{\"id\":\"a\",\"timestamp\":1509493677,\"temperature\":3.66},{\"id\":\"b\",\"timestamp\":1510113646,\"temperature\":4.15},{\"id\":\"c\",\"timestamp\":1510127886,\"temperature\":3.36},{\"id\":\"c\",\"timestamp\":1510127892,\"temperature\":3.36},{\"id\":\"a\",\"timestamp\":1510128112,\"temperature\":3.67},{\"id\":\"b\",\"timestamp\":1510128115,\"temperature\":3.88}]",
		}

		response, err := handler(request)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if response.StatusCode != 200 {
			t.Fatalf("Incorrect status code: %d", response.StatusCode)
		}

		decodedResponse := []Stats{}
		if err := json.Unmarshal([]byte(response.Body), &decodedResponse); err != nil {
			t.Fatalf("Response does not seem to be well JSON formatted: %s", err.Error())
		}

		for _, s := range decodedResponse {
			switch s.Sensor {
			case "a":
				if !reflect.DeepEqual(s, Stats{
					Sensor:  "a",
					Average: 3.78,
					Median:  3.65,
					Mode:    []float64{3.53},
				}) {
					t.Fatalf("Incorrect stats %v", s)
				}
			case "b":
				if !reflect.DeepEqual(s, Stats{
					Sensor:  "b",
					Average: 4.08,
					Median:  4.14,
					Mode:    []float64{4.15},
				}) {
					t.Fatalf("Incorrect stats %v", s)
				}
			case "c":
				if !reflect.DeepEqual(s, Stats{
					Sensor:  "c",
					Average: 3.72,
					Median:  3.95,
					Mode:    []float64{3.36, 3.96},
				}) {
					t.Fatalf("Incorrect stats %v", s)
				}
			default:
				t.Fatalf("Incorrect stats %v", s)
			}
		}
	})

	t.Run("Invalid request JSON body", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Body: "}",
		}

		response, err := handler(request)

		if response.StatusCode != 400 {
			t.Fatalf("Incorrect status code: %d", response.StatusCode)
		}

		if err == nil {
			t.Fatal("Expected error but none received")
		}

		if err.Error() != "invalid character '}' looking for beginning of value" {
			t.Fatalf("Incorrect error message: %s", err.Error())
		}
	})

	t.Run("Invalid request JSON body", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Body: "}",
		}

		response, err := handler(request)

		if response.StatusCode != 400 {
			t.Fatalf("Incorrect status code: %d", response.StatusCode)
		}

		if err == nil {
			t.Fatal("Expected error but none received")
		}

		if err.Error() != "invalid character '}' looking for beginning of value" {
			t.Fatalf("Incorrect error message: %s", err.Error())
		}
	})

	t.Run("No temperature", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Body: "[{\"id\":\"a\",\"timestamp\":1509493641}]",
		}

		response, err := handler(request)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if response.StatusCode != 200 {
			t.Fatalf("Incorrect status code: %d", response.StatusCode)
		}

		decodedResponse := []Stats{}
		if err := json.Unmarshal([]byte(response.Body), &decodedResponse); err != nil {
			t.Fatalf("Response does not seem to be well JSON formatted: %s", err.Error())
		}

		for _, s := range decodedResponse {
			switch s.Sensor {
			case "a":
				if !reflect.DeepEqual(s, Stats{
					Sensor:  "a",
					Average: 0,
					Median:  0,
					Mode:    []float64{0},
				}) {
					t.Fatalf("Incorrect stats %v", s)
				}
			default:
				t.Fatalf("Incorrect stats %v", s)
			}
		}
	})

	t.Run("No sensor ID", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Body: "[{\"timestamp\":1509493641,\"temperature\":5.5}]",
		}

		response, err := handler(request)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if response.StatusCode != 200 {
			t.Fatalf("Incorrect status code: %d", response.StatusCode)
		}

		decodedResponse := []Stats{}
		if err := json.Unmarshal([]byte(response.Body), &decodedResponse); err != nil {
			t.Fatalf("Response does not seem to be well JSON formatted: %s", err.Error())
		}

		for _, s := range decodedResponse {
			switch s.Sensor {
			case "":
				if !reflect.DeepEqual(s, Stats{
					Sensor:  "",
					Average: 5.5,
					Median:  5.5,
					Mode:    []float64{5.5},
				}) {
					t.Fatalf("Incorrect stats %v", s)
				}
			default:
				t.Fatalf("Incorrect stats %v", s)
			}
		}
	})

	t.Run("No readings", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Body: "[]",
		}

		response, err := handler(request)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}

		if response.StatusCode != 200 {
			t.Fatalf("Incorrect status code: %d", response.StatusCode)
		}

		if response.Body != "[]" {
			t.Fatalf("Incorrect response: %s", response.Body)
		}
	})
}
