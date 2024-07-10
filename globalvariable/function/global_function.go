package function

import (
	"testbe/config"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//This function is used to insert a dynamic column and values, for example
//tableName := "users_profile"
//columnNames := []string{"user_id", "total_post", "total_experience", "total_comments", "avg_user_stars"}
//values := []interface{}{123, 0, 0, 0, 0}
//insertData, err := dynamicInsertDB(tableName, columnNames, values), it will return boolean and error
func DynamicInsertDB(tableName string, columnNames []string, values []interface{}) (bool, error) {
    db := config.CreateConnection()
    defer db.Close()

    // Prepare the SQL statement dynamically
    var placeholders []string
    for range columnNames {
        placeholders = append(placeholders, "?")
    }

    sqlStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columnNames, ", "), strings.Join(placeholders, ", "))

    // Prepare the SQL statement and the data to be inserted
    stmt, err := db.Prepare(sqlStatement)
    if err != nil {
        return false, fmt.Errorf("error preparing SQL statement: %v", err)
    }
    defer stmt.Close()

    // Execute the SQL statement and handle any errors
    _, err = stmt.Exec(values...)
    if err != nil {
        return false, fmt.Errorf("error executing SQL statement: %v", err)
    }

    return true, nil
}

func GenerateRandomPassword() string {
	rand.NewSource(time.Now().UnixNano())

	const (
		letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		length  = 8
	)

	password := make([]byte, length)
	for i := range password {
		password[i] = letters[rand.Intn(len(letters))]
	}

	return string(password)
}

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

// Function to generate ticket number
func GenerateTicketString() string {
    // Get current date string
    dateString := time.Now().Format("20060102")
  
    // Define character set for random string
    charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
  
    // Generate random string
    randomString := make([]byte, 10)
    for i := range randomString {
      randomString[i] = charset[rand.Intn(len(charset))]
    }
  
    // Combine date string and random string
    ticketString := "TX-" + dateString + "-" + string(randomString)
  
    return ticketString
}