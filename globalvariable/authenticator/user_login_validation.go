package authenticator

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	Lg"testbe/logging"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"testbe/config"
	Ct "testbe/globalvariable/constant"
	Gv "testbe/globalvariable/variable"
)

func Validate(username, password string) (string, error) {
	// Connect to the database.
	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `SELECT token FROM users_login WHERE username = $1 AND password = $2`

	// Execute the SQL statement.
	var token sql.NullString
	err := db.QueryRow(sqlStatement, username, password).Scan(&token)

	// If the user is not found, return an error.
	if err == sql.ErrNoRows {
		Lg.Error("this uername: ", username+" is not found in database")
		return "", fmt.Errorf("user not found")
	}

	// If there's an error in executing the SQL statement, return the error.
	if err != nil {
		log.Fatalf("Error executing the SQL statement: %v", err)
		Lg.Error("this uername: ", username+" is error while executing the SQL statement with error: "+err.Error())
		return "", err
	}

	if token.Valid {
		return token.String, nil
	} else {
		return "", nil
	}
}

func ValidateGenerateNewToken(username, password string) (string, error) {
	// Connect to the database.
	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `UPDATE users_login SET token = $1 WHERE username = $2`

	sum := md5.Sum([]byte(password + username + Gv.CurrentTime.String()))
	tokenGenerated := hex.EncodeToString(sum[:])

	_, errExec := db.Exec(sqlStatement, tokenGenerated, username)

	if errExec != nil {
		return "", errExec
	}

	return tokenGenerated, nil
}

// Validate email to get UUID, but its critical and not really necessary use other than needed ones.
func ValidateEmailGetUUID(email string) (string, error) {
		// Connect to the database.
		db := config.CreateConnection()

		// Close the connection at the end of the process.
		defer db.Close()
	
		// Create a SQL query to retrieve the token based on the username and password.
		sqlStatement := `SELECT user_id FROM users_profile WHERE email = $1`
	
		// Execute the SQL statement.
		var uuid sql.NullString
		err := db.QueryRow(sqlStatement, email).Scan(&uuid)
	
		// If the user is not found, return an error.
		if err == sql.ErrNoRows {
			Lg.Error("This Email: ",email+" does not exist")
			return "", fmt.Errorf("%s", "User Validation - user not found")
		}
	
		// If there's an error in executing the SQL statement, return the error.
		if err != nil {
			log.Fatalf("Error executing the SQL statement: %v", err)
			Lg.Error("This email: ",email+" facing error executing the SQL statement with error: "+err.Error())
			return "", err
		}
	
		if uuid.Valid {
			Lg.Info("this email: "+email+" successfully getting the userId from Database")
			return uuid.String, nil
		} else {
			Lg.Error("This email: ",email+" has faileed the username validation")
			return "", fmt.Errorf("%s", "Username Validation - Invalid Username")
		}
}

// Validate users token to get user id
func ValidateUsernameGetUUID(userName string) (string, error) {
	// Connect to the database.
	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `SELECT user_id FROM users_login WHERE username = $1`

	// Execute the SQL statement.
	var uuid sql.NullString
	err := db.QueryRow(sqlStatement, userName).Scan(&uuid)

	// If the user is not found, return an error.
	if err == sql.ErrNoRows {
		Lg.Error("This username: ",userName+" does not exist")
		return "", fmt.Errorf("%s", "User Validation - user not found")
	}

	// If there's an error in executing the SQL statement, return the error.
	if err != nil {
		log.Fatalf("Error executing the SQL statement: %v", err)
		Lg.Error("This username: ",userName+" facing error executing the SQL statement with error: "+err.Error())
		return "", err
	}

	if uuid.Valid {
		Lg.Info("this username: "+userName+" successfully getting the userId from Database")
		return uuid.String, nil
	} else {
		Lg.Error("This username: ",userName+" has faileed the username validation")
		return "", fmt.Errorf("%s", "Username Validation - Invalid Username")
	}
}

// Validate users username to get stored password
func ValidateGetStoredPasswordByUsername(username string) (string, error) {

	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `SELECT password FROM users_login WHERE username = $1`

	// Execute the SQL statement.
	var password sql.NullString

	err := db.QueryRow(sqlStatement, username).Scan(&password)

	if err == sql.ErrNoRows {
		Lg.Error("This username: ",  username+" face error with error: Sql No Rows returned")
		return "", fmt.Errorf("%s", "Password not found")
	}

	if err != nil {
		log.Fatalf("Error executing the SQL statement: %v", err)
		Lg.Error("This username: ",  username+" face error with error: "+err.Error())
		return "", err
	}

	Lg.Info("This username: "+username+" Succesfully get password using username for validating")
	return password.String, nil
}

func ValidateGetStoredPasswordByUserId(userId string) (string, error) {

	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `SELECT password FROM users_login WHERE user_id = $1`

	// Execute the SQL statement.
	var password sql.NullString

	err := db.QueryRow(sqlStatement, userId).Scan(&password)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("%s", "Password not found")
	}

	if err != nil {
		log.Fatalf("Error executing the SQL statement: %v", err)
		return "", err
	}

	return password.String, nil
}

// Validate username to get email
func ValidateUsnGetEmail(userId string) (string, error) {
	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `SELECT email FROM users_profile WHERE user_id = $1`

	// Execute the SQL statement.

	var result string
	err := db.QueryRow(sqlStatement, userId).Scan(&result)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("VALIDATE USERNAME - No rows")
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

// Validate username when user create an account
func ValidateCreateNewUsername(username string) (bool, error) {
	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `SELECT username FROM users_login WHERE username = $1`

	// Execute the SQL statement.

	var result string
	err := db.QueryRow(sqlStatement, username).Scan(&result)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	if result == username {
		return false, nil
	} else {
		return true, nil
	}
}

// Validate email when user create an account
func ValidateCreateNewEmail(email string) (bool, error) {
	db := config.CreateConnection()

	// Close the connection at the end of the process.
	defer db.Close()

	// Create a SQL query to retrieve the token based on the username and password.
	sqlStatement := `SELECT email FROM users_profile WHERE email = $1`

	// Execute the SQL statement.

	var result string
	err := db.QueryRow(sqlStatement, email).Scan(&result)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	if result == email {
		return false, nil
	} else {
		return true, nil
	}
}

// Validate users password before login into account
func ValidateUserPassword(enteredPassword string, storedPassword string) (bool, error) {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), append(Ct.Salt, []byte(enteredPassword)...)) == nil, nil
}

// Validate user password to get encrypted password
func ValidatePasswordToEncrypt(password string) (string, error) {
	hashedPassword, errhashed := bcrypt.GenerateFromPassword(append(Ct.Salt, []byte(password)...), Ct.BcryptCost)

	if errhashed != nil {
		return "", errhashed
	}
	return string(hashedPassword), nil
}
