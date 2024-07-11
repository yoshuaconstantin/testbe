package function

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)


// Send Json response for errors
func JsonResponseWithError(w http.ResponseWriter, errorMsg string, statusCode int) {
    // Encode error message as JSON
    jsonError := map[string]string{"error": errorMsg}
    jsonResponse, _ := json.Marshal(jsonError)

    // Set appropriate headers and return JSON response
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(statusCode)
    w.Write(jsonResponse)
}

// Dynamic check query parameters status
func CheckQueryParameters(param string) error {
    	// Check if queryParam is null or empty
	if len(param) == 0 {
		return fmt.Errorf("queryParam is required")
	}

	// Check if email queryParam is null or empty
	if param == "" {
		return fmt.Errorf("Old password queryParam is required")
	}

    return nil
}

// Function to generate random string
func GenerateRandomString() string {
    // Get current date string
    //dateString := time.Now().Format("20060102")
  
    // Define character set for random string
    charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
  
    // Generate random string
    randomString := make([]byte, 10)
    for i := range randomString {
      randomString[i] = charset[rand.Intn(len(charset))]
    }
  
  
    return string(randomString)
}