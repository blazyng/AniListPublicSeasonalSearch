package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// Import the generated package
	"github.com/blazyng/AniListPublicSeasonalSearch/generated"
)

func main() {
	// Create a standard HTTP client
	httpClient := http.DefaultClient

	// The generated code provides a NewClient function
	graphqlClient := generated.NewClient(httpClient)

	// Call the generated function, which is named after your query
	response, err := generated.GetAnimeBySeason(
		context.Background(),
		graphqlClient,
		"WINTER", // The value for the $season variable
		2024,     // The value for the $seasonYear variable
	)
	if err != nil {
		log.Fatal(err)
	}

	// Now you can access the response data with type safety!
	// No more manual JSON parsing.
	for _, anime := range response.Page.Media {
		fmt.Printf("- %s (%s)\n", anime.Title.Romaji, anime.Title.English)
	}
}
