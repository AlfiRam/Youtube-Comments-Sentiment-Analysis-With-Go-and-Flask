# Youtube Comments Sentiment Analysis With Go and Flask
 
This project is a sentiment analysis service that utilizes a machine learning model to analyze and predict the sentiment of given text inputs. The service is built using a Python Flask backend application, which hosts the machine learning model, and a Go (Golang) application that acts as a client for making requests to the Flask service. Additionally, the Go application fetches comments from any YouTube video given its ID and then sends these comments to the Flask service for sentiment analysis.

# Prerequisites
- Youtube API Key: Ensure you have a youtube API Key (YouTube Data API v3)
- Python: The project requires Python, as the Flask application and machine learning model are Python-based.
  - Python Libraries: flask, tensorflow, numpy, pickle, etc.
  - A trained machine learning model file and a tokenizer file are also required. (I have provided mine **'sentiment_analysis_model.h5'**  and **'tokenizer.pickle'**)
- Go (Golang): Used for the client application that sends requests to the Flask backend and fetches comments from YouTube.
  - Go packages:
    - github.com/joho/godotenv
    - google.golang.org/api/option
    - google.golang.org/api/youtube/v3
- Postman (Optional): For testing the API endpoint manually.

# Set API KEY
in the `.env` file, set ```API_KEY``` to your youtube API key.

# Run Flask App

```python return_sentiment.py```

# Test with Postman (Optional)

- Make a POST request to http://localhost:5000/analyze (adjust the URL if your Flask app is hosted differently).
- In the request body, send a JSON object with the text to be analyzed:
```
{
  "text": "Your text here"
}
```

# Run Go client app
Run the Go application and input a YouTube video ID when prompted. In `main.go`, replace `"youtube_video_id"` with the actual id of a youtube video of your choice, which can be found between the `v=_____&`  of the video url. The application will fetch comments from the specified video and send them to the Flask backend for sentiment analysis.

```go run main.go```



