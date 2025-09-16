package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// The Go structs to hold the JSON response from the API.
// We use `json:"..."` tags to map the JSON keys to our struct fields.
type ResponseData struct {
	Data struct {
		Page struct {
			Media []struct {
				ID    int `json:"id"`
				Title struct {
					Romaji  string `json:"romaji"`
					English string `json:"english"`
				} `json:"title"`
			} `json:"media"`
		} `json:"Page"`
	} `json:"data"`
}

// A struct to hold our GraphQL query and variables, which we will convert to JSON.
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	// Define the GraphQL query as a multi-line string.
	query := `
		query($season: MediaSeason, $seasonYear: Int) {
			Page {
				media(season: $season, seasonYear: $seasonYear, type: ANIME, sort: POPULARITY_DESC) {
					id
					title {
						romaji
						english
					}
				}
			}
		}`

	// Create the request payload object.
	requestPayload := GraphQLRequest{
		Query: query,
		Variables: map[string]interface{}{
			"season":     "WINTER",
			"seasonYear": 2024,
		},
	}

	// Convert the payload object to a JSON byte buffer.
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatal("Error creating JSON payload:", err)
	}
	payloadBuffer := bytes.NewBuffer(payloadBytes)

	// Create the HTTP POST request.
	req, err := http.NewRequestWithContext(context.Background(), "POST", "https://graphql.anilist.co", payloadBuffer)
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Decode the JSON response into our Go structs.
	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	// Print the results.
	fmt.Println("Successfully fetched anime for WINTER 2024:")
	for _, anime := range responseData.Data.Page.Media {
		if anime.Title.Romaji != "" {
			fmt.Printf("- %s (%s)\n", anime.Title.Romaji, anime.Title.English)
		}
	}
}
