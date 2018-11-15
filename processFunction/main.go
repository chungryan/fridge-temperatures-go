package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chungryan/fridge-temperatures-go/math"
)

// Temperature value
type Temperature = float64

// Reading of a temperature sensor
type Reading struct {
	Sensor      string      `json:"id"`
	Timestamp   uint        `json:"timestamp"`
	Temperature Temperature `json:"temperature"`
}

// Stats of a sensor
type Stats struct {
	Sensor  string        `json:"id"`
	Average Temperature   `json:"average"`
	Median  Temperature   `json:"median"`
	Mode    []Temperature `json:"mode"`
}

// Get a simple analysis of the sensor temperatures
func analyse(temperatures map[string][]Temperature) ([]Stats, error) {
	stats := []Stats{}

	for s, t := range temperatures {
		avg, err := math.Average(t)
		if err != nil {
			return nil, err
		}

		med, err := math.Median(t)
		if err != nil {
			return nil, err
		}

		mode, err := math.Mode(t)
		if err != nil {
			return nil, err
		}

		stats = append(stats, Stats{
			Sensor:  s,
			Average: avg,
			Median:  med,
			Mode:    mode,
		})
	}

	return stats, nil
}

// Group the reading temperatures by sensor
func groupTemperaturesBySensor(readings []Reading) map[string][]Temperature {
	temperatures := map[string][]Temperature{}

	for _, r := range readings {
		temperatures[r.Sensor] = append(temperatures[r.Sensor], r.Temperature)
	}

	return temperatures
}

// Decode the JSON body into temperature object list
func decodeReadings(requestBody string) ([]Reading, error) {
	var readings []Reading

	if err := json.Unmarshal([]byte(requestBody), &readings); err != nil {
		return nil, err
	}

	return readings, nil
}

// Process the readings an get the statistics on the senrors
func processReadings(encodedReadings string) (*string, error) {
	readings, err := decodeReadings(encodedReadings)
	if err != nil {
		return nil, err
	}

	temperatures := groupTemperaturesBySensor(readings)
	stats, err := analyse(temperatures)
	if err != nil {
		return nil, err
	}

	encodedStats, err := json.Marshal(stats)
	if err != nil {
		return nil, err
	}

	finalStats := string(encodedStats)
	return &finalStats, nil
}

// Lambda handler function
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, err := processReadings(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       *response,
	}, nil
}

func main() {
	lambda.Start(handler)
}
