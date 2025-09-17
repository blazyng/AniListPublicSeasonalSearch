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
				StartDate struct { // NEW
					Year  int `json:"year"`
					Month int `json:"month"`
					Day   int `json:"day"`
				} `json:"startDate"`
				EndDate struct { // NEW
					Year  int `json:"year"`
					Month int `json:"month"`
					Day   int `json:"day"`
				} `json:"endDate"`
				Episodes   int      `json:"episodes"` // NEW
				CoverImage struct { // NEW
					Large string `json:"large"`
				} `json:"coverImage"`
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
					startDate { # NEW
						year
						month
						day
					}
					endDate { # NEW
						year
						month
						day
					}
					episodes     # NEW
					coverImage { # NEW
						large
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
		// Print title and episode count
		fmt.Printf("\n---\nTitle: %s (%s)\n", anime.Title.Romaji, anime.Title.English)
		fmt.Printf("Episodes: %d\n", anime.Episodes)

		// Format and print start date
		startDate := fmt.Sprintf("%d-%02d-%02d", anime.StartDate.Year, anime.StartDate.Month, anime.StartDate.Day)
		fmt.Printf("Start Date: %s\n", startDate)

		// Only print end date if it exists (year is not 0)
		if anime.EndDate.Year != 0 {
			endDate := fmt.Sprintf("%d-%02d-%02d", anime.EndDate.Year, anime.EndDate.Month, anime.EndDate.Day)
			fmt.Printf("End Date: %s\n", endDate)
		}

		// Print cover image link
		fmt.Printf("Cover Image: %s\n", anime.CoverImage.Large)
	}
}
