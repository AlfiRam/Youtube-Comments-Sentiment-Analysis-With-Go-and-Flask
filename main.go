package main

import (
    "context"
    "fmt"
    "log"
    "os"

	"github.com/joho/godotenv"
    "google.golang.org/api/option"
    "google.golang.org/api/youtube/v3"
	"bytes"
    "encoding/json"
    "net/http"
)

type SentimentResponse struct {
    Sentiment string `json:"sentiment"`
}

func analyzeSentiment(comment string) (string, error) {
    requestBody, err := json.Marshal(map[string]string{
        "text": comment,
    })
    if err != nil {
        return "", err
    }

    resp, err := http.Post("http://localhost:5000/analyze", "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var result SentimentResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    return result.Sentiment, nil
}

func getComments(service *youtube.Service, videoID string) ([]string, error) {
    var comments []string
    call := service.CommentThreads.List([]string{"snippet"}).VideoId(videoID).TextFormat("plainText")
    response, err := call.Do()
    if err != nil {
        return nil, err
    }

    for _, item := range response.Items {
        comment := item.Snippet.TopLevelComment.Snippet.TextDisplay
        comments = append(comments, comment)
    }

    return comments, nil
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	ctx := context.Background()

    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        log.Fatal("API_KEY environment variable not set")
    }

    service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        log.Fatalf("Error creating YouTube client: %v", err)
    }

    videoID := "youtube_video_id" // Set video ID here
    comments, err := getComments(service, videoID)
    if err != nil {
    log.Fatalf("Error fetching comments: %v", err)
    }

	//fetch comments
    for _, comment := range comments {
        fmt.Println(comment)
    }

	var positiveComments, negativeComments []string

    for _, comment := range comments {
        sentiment, err := analyzeSentiment(comment)
        if err != nil {
            log.Printf("Error analyzing sentiment for comment '%s': %v", comment, err)
            continue
        }

        if sentiment == "Positive" {
            positiveComments = append(positiveComments, comment)
        } else if sentiment == "Negative" {
            negativeComments = append(negativeComments, comment)
        }
    }

    fmt.Printf("Total Positive Comments: %d\n", len(positiveComments))
    fmt.Printf("Total Negative Comments: %d\n", len(negativeComments))

    fmt.Println("\nPositive Comments:")
    for _, comment := range positiveComments {
        fmt.Println(comment)
    }

    fmt.Println("\nNegative Comments:")
    for _, comment := range negativeComments {
        fmt.Println(comment)
    }
}
