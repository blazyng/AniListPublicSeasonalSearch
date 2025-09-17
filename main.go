package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

// The Go structs to hold the JSON response from the API.
type ResponseData struct {
	Data struct {
		Page struct {
			Media []struct {
				ID    int `json:"id"`
				Title struct {
					Romaji  string `json:"romaji"`
					English string `json:"english"`
				} `json:"title"`
				StartDate struct {
					Year  int `json:"year"`
					Month int `json:"month"`
					Day   int `json:"day"`
				} `json:"startDate"`
				EndDate struct {
					Year  int `json:"year"`
					Month int `json:"month"`
					Day   int `json:"day"`
				} `json:"endDate"`
				Episodes   int `json:"episodes"`
				CoverImage struct {
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

// A clean struct for our final JSON output.
type AnimeOutput struct {
	Title        string `json:"title"`
	EnglishTitle string `json:"english_title,omitempty"`
	Episodes     int    `json:"episodes"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date,omitempty"`
	CoverImage   string `json:"cover_image"`
}

// GetAnimeBySeason handles the logic of fetching and parsing anime data from the API.
func GetAnimeBySeason(year int, season string) ([]AnimeOutput, error) {
	query := `
		query($season: MediaSeason, $seasonYear: Int) {
			Page {
				media(season: $season, seasonYear: $seasonYear, type: ANIME, sort: POPULARITY_DESC) {
					id
					title { romaji english }
					startDate { year month day }
					endDate { year month day }
					episodes
					coverImage { large }
				}
			}
		}`

	requestPayload := GraphQLRequest{
		Query: query,
		Variables: map[string]interface{}{
			"season":     season,
			"seasonYear": year,
		},
	}

	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, fmt.Errorf("error creating JSON payload: %w", err)
	}
	payloadBuffer := bytes.NewBuffer(payloadBytes)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "https://graphql.anilist.co", payloadBuffer)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	var results []AnimeOutput
	for _, anime := range responseData.Data.Page.Media {
		output := AnimeOutput{
			Title:        anime.Title.Romaji,
			EnglishTitle: anime.Title.English,
			Episodes:     anime.Episodes,
			StartDate:    fmt.Sprintf("%d-%02d-%02d", anime.StartDate.Year, anime.StartDate.Month, anime.StartDate.Day),
			CoverImage:   anime.CoverImage.Large,
		}
		if anime.EndDate.Year != 0 {
			output.EndDate = fmt.Sprintf("%d-%02d-%02d", anime.EndDate.Year, anime.EndDate.Month, anime.EndDate.Day)
		}
		results = append(results, output)
	}

	return results, nil
}

func main() {
	// --- Main function---

	// 1. Handle command-line input
	currentYear := time.Now().Year()
	yearPtr := flag.Int("year", currentYear, "The year of the anime season (e.g., 2025)")
	seasonPtr := flag.String("season", "WINTER", "The anime season (WINTER, SPRING, SUMMER, FALL)")
	flag.Parse()

	// 2. Call the core logic
	results, err := GetAnimeBySeason(*yearPtr, *seasonPtr)
	if err != nil {
		log.Fatal(err) // If an error occurs, exit the program
	}

	// 3. Handle the final output
	jsonOutput, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatal("Error creating JSON output:", err)
	}

	fmt.Println(string(jsonOutput))
}
