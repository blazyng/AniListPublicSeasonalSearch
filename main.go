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

type AnimeOutput struct {
	Title        string `json:"title"`
	EnglishTitle string `json:"english_title,omitempty"`
	Episodes     int    `json:"episodes"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date,omitempty"` // omitempty:
	CoverImage   string `json:"cover_image"`
}

func main() {
	// Define command-line flags.
	// We use the current year as the default value for the 'year' flag.
	currentYear := time.Now().Year()
	yearPtr := flag.Int("year", currentYear, "The year of the anime season (e.g., 2025)")
	seasonPtr := flag.String("season", "WINTER", "The anime season (WINTER, SPRING, SUMMER, FALL)")

	// Parse the flags from the command line.
	flag.Parse()
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

	// Create the request payload object using the values from the flags.
	requestPayload := GraphQLRequest{
		Query: query,
		Variables: map[string]interface{}{
			"season":     *seasonPtr, //
			"seasonYear": *yearPtr,   //
		},
	}

	// Convert the payload object to a JSON byte buffer.
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatal("Error creating JSON payload:", err)
	}
	payloadBuffer := bytes.NewBuffer(payloadBytes)

	req, err := http.NewRequestWithContext(context.Background(), "POST", "https://graphql.anilist.co", payloadBuffer)
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	// --- ÄNDERUNG: Die formatierte Textausgabe wird ersetzt ---

	// 1. Erstelle eine Liste (slice) für unsere sauberen Anime-Daten.
	var results []AnimeOutput

	// 2. Gehe durch die API-Antwort und fülle unsere saubere Liste.
	for _, anime := range responseData.Data.Page.Media {
		output := AnimeOutput{
			Title:        anime.Title.Romaji,
			EnglishTitle: anime.Title.English,
			Episodes:     anime.Episodes,
			StartDate:    fmt.Sprintf("%d-%02d-%02d", anime.StartDate.Year, anime.StartDate.Month, anime.StartDate.Day),
			CoverImage:   anime.CoverImage.Large,
		}
		// Füge das Enddatum nur hinzu, wenn es existiert.
		if anime.EndDate.Year != 0 {
			output.EndDate = fmt.Sprintf("%d-%02d-%02d", anime.EndDate.Year, anime.EndDate.Month, anime.EndDate.Day)
		}
		results = append(results, output)
	}

	// 3. Konvertiere unsere saubere Liste in schön formatiertes JSON.
	jsonOutput, err := json.MarshalIndent(results, "", "  ") // "" für kein Prefix, "  " für 2 Leerzeichen Einrückung
	if err != nil {
		log.Fatal("Error creating JSON output:", err)
	}

	// 4. Gib das finale JSON auf der Konsole aus.
	fmt.Println(string(jsonOutput))
}
