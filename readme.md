# Go Anime Season Fetcher

A small Go project to fetch seasonal anime data from the [AniList GraphQL API](https://docs.anilist.co/guide/graphql/).

## 📖 Overview

This project is a simple Go application that queries the AniList API for a list of anime released in a specific year and season.

The main goal of this project is to learn the Go programming language by building a practical, real-world application that interacts with an external API.

## ✨ Project Status

This section tracks the implemented and planned features.

### ✅ Implemented
-   [x] **Core API Logic:** Successfully sends a GraphQL query to the AniList API using Go's standard `net/http` and `encoding/json` packages.
-   [x] **JSON Parsing:** Correctly parses the JSON response from the API into Go structs.
-   [x] **Basic Output:** Prints a formatted list of anime titles to the console.

### ⬜️ To-Do / Planned Features
-   [ ] **Dynamic Input:** Accept `year` and `season` as command-line arguments instead of being hardcoded.
-   [ ] **JSON Output:** Return the final data as a clean, formatted JSON string.
-   [ ] **Refactor:** Encapsulate the API logic into a reusable function `GetAnimeBySeason(year int, season string)`.
-   [ ] **Improved Error Handling:** Add more robust error handling for API and network issues.

## 🚀 Tech Stack

* **Go** - Including the standard libraries:
    * `net/http` for making API requests.
    * `encoding/json` for handling JSON data.
* **AniList GraphQL API** - As the data source for all anime information.

## 🛠️ Prerequisites

Make sure you have [Go](https://go.dev/doc/install) installed on your system.

## ⚙️ Installation & Usage

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/YOUR-GITHUB-USERNAME/YOUR-REPO-NAME.git](https://github.com/YOUR-GITHUB-USERNAME/YOUR-REPO-NAME.git)
    cd YOUR-REPO-NAME
    ```

2.  **Run the application:**
    Currently, the year and season are hardcoded in `main.go`. You can run the program directly:
    ```bash
    go run .
    ```

3.  **Planned Usage (Future):**
    The goal is to run the tool with command-line flags:
    ```bash
    go run . --year 2025 --season SPRING
    ```

## 📚 API Reference

This project uses the public [AniList API](https://anilist.co). The documentation for their GraphQL endpoint can be found here: [AniList APIv2 Docs](https://anilist.github.io/ApiV2-GraphQL-Docs/).

## 🎯 Learning Goals

* Go fundamentals (variables, structs, functions, error handling).
* Making HTTP POST requests and building a request body.
* Interacting with a GraphQL API without a third-party client.
* Processing JSON data (marshalling & unmarshalling).
* Structuring a simple Go command-line application.