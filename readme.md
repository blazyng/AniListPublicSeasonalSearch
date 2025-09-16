# Go Anime Season Fetcher

A small Go project to fetch seasonal anime data from the [AniList GraphQL API](https://docs.anilist.co/guide/graphql/).

## 📖 Overview

This project is a simple Go application, designed as a function or a small command-line tool. You provide a **year** and a **season** (WINTER, SPRING, SUMMER, FALL), and the program returns a list of anime from that season in JSON format.

The main goal of this project is to learn the Go programming language and gain practical experience with external APIs.

## ✨ Features

* Takes a year and a season as input.
* Sends a request to the AniList GraphQL API.
* Parses the JSON response into Go structs.
* Outputs the filtered anime data as clean JSON.

## 🚀 Tech Stack

* **Go** - The programming language of choice
* **AniList GraphQL API** - The data source for all anime information

## 🛠️ Prerequisites

Make sure you have [Go](https://go.dev/doc/install) installed on your system.

## ⚙️ Installation & Usage

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/YOUR-GITHUB-USERNAME/YOUR-REPO-NAME.git](https://github.com/YOUR-GITHUB-USERNAME/YOUR-REPO-NAME.git)
    cd YOUR-REPO-NAME
    ```

2.  **Install dependencies:**
    *(This step will become relevant once you add external packages like a GraphQL client.)*
    ```bash
    go mod tidy
    ```

3.  **Run the application:**
    *(An example of what the final command might look like)*
    ```bash
    go run main.go --year 2024 --season WINTER
    ```

## 📚 API Reference

This project uses the public [AniList API](https://anilist.co). You can find the documentation for their GraphQL endpoint here: [AniList APIv2 Docs](https://anilist.github.io/ApiV2-GraphQL-Docs/).

## 🎯 Learning Goals

* Go fundamentals (variables, structs, functions, error handling).
* Making HTTP requests and interacting with an external API in Go.
* Using a GraphQL client.
* Processing JSON data (marshalling & unmarshalling).
* Structuring a simple Go project.