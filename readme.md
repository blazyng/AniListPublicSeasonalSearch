# Go Anime Season Fetcher

A small Go command-line tool to fetch seasonal anime data from the [AniList GraphQL API](https://docs.anilist.co/guide/graphql/).

## 📖 Overview

This project is a simple Go application that queries the AniList API for a list of anime released in a specific year and season. It takes the year and season as command-line arguments and returns a clean JSON array of the results.

The main goal of this project was to learn the Go programming language by building a practical application that interacts with an external API.

## ✨ Project Status

This section tracks the implemented and planned features.

### ✅ Implemented
-   [x] **Core API Logic:** Successfully sends a GraphQL query to the AniList API using Go's standard `net/http` and `encoding/json` packages.
-   [x] **JSON Parsing:** Correctly parses the JSON response from the API into Go structs.
-   [x] **Dynamic Input:** Accepts `year` and `season` as command-line flags.
-   [x] **JSON Output:** Returns the final data as a clean, formatted JSON string.
-   [x] **Refactor:** The core API logic is encapsulated in a reusable function `GetAnimeBySeason(year int, season string)`.

### ⬜️ Future Improvements
-   [ ] **Improved Error Handling:** Add more user-friendly error messages for API or network issues.
-   [ ] **Add More Fields:** The query and structs can be expanded to include more data per anime (e.g., studio, genre, score).
-   [ ] **Unit Tests:** Add tests for the `GetAnimeBySeason` function.

## 🚀 Tech Stack

* **Go** - Including the standard libraries:
    * `net/http` for making API requests.
    * `encoding/json` for handling JSON data.
    * `flag` for command-line argument parsing.
* **AniList GraphQL API** - As the data source for all anime information.

## 🛠️ Prerequisites

Make sure you have [Go](https://go.dev/doc/install) installed on your system.

## ⚙️ Installation & Usage

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/blazyng/AniListPublicSeasonalSearch
    cd AniListPublicSeasonalSearch
    ```

2.  **Run the application:**
    Use the `-year` and `-season` flags to specify the desired query.
    ```bash
    # Get anime for Summer 2024
    go run . -year 2024 -season SUMMER

    # Get anime for Fall 2025
    go run . -year 2025 -season FALL

    # Get help
    go run . -h
    ```

## 📚 API Reference

This project uses the public [AniList API](https://anilist.co). The documentation for their GraphQL endpoint can be found here: [AniList APIv2 Docs](https://docs.anilist.co/reference/).

## 🎯 Learning Goals Achieved

* Go fundamentals (variables, structs, functions, error handling).
* Making HTTP POST requests and building a request body.
* Interacting with a GraphQL API without a third-party client.
* Processing JSON data (marshalling & unmarshalling).
* Structuring a simple Go command-line application and separating logic from execution.
