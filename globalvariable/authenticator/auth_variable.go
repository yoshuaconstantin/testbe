package authenticator

import (
	"testbe/config"
	Gv "testbe/globalvariable/variable"
	Lg "testbe/logging"
	jwttoken "testbe/middleware"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

// Authenticatior main function to securely authenticate request also shorten usage code and it will be added more to the function
func SecureAuthenticator(w http.ResponseWriter, r *http.Request, userName string) (string, int, error) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", http.StatusBadRequest, fmt.Errorf("%s", "Missing authorization header")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	CheckJwtTokenValidation, userId, erroCheckJWt := jwttoken.VerifyToken(tokenString)

	if erroCheckJWt != nil || !CheckJwtTokenValidation {
		return "", http.StatusUnauthorized, erroCheckJWt
	}

	return userId, 0, nil
}

// Function to validate the token and return the userId and JWT token to be used
func CheckJWTnGetUserId(w http.ResponseWriter, r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("%s", "Missing authorization header")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	CheckJwtTokenValidation, userId, erroCheckJWt := jwttoken.VerifyToken(tokenString)

	if erroCheckJWt != nil {
		return "", erroCheckJWt
	}

	if !CheckJwtTokenValidation {
		return "", fmt.Errorf("Token validation failed")
	}

	return userId, nil
}

func OTPVerificationProcess(email string, code string, moduleName string) (bool, error) {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `SELECT verification_code FROM one_time_verification WHERE email=$1 AND verification_code = $2`

	var verificationCode string
	err := db.QueryRow(sqlStatement, email, code).Scan(&verificationCode)

	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("\nVERIFY OTP - Error: %v\n", err)
		return false, err
	}

	if verificationCode == "" {
		return false, nil
	}

	// Update the verification code to an empty string after successful verification
	updateStatement := `UPDATE one_time_verification SET verification_code = '', modules = $1, update_at = $2 WHERE email=$3 AND verification_code=$4`
	_, err = db.Exec(updateStatement, moduleName, Gv.FormattedTimeNowYYYYMMDDHHMM, email, code)

	if err != nil {
		log.Fatalf("\nVERIFIFY OTP - Error updating verification code : %v\n", err)
		return false, err
	}

	return true, nil
}

func SendOTPVerificationCodeToDB(email string, modules string) (string, error) {
	db := config.CreateConnection()
	defer db.Close()

	// Check if a record with the given email exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM one_time_verification WHERE email = $1", email).Scan(&count)
	if err != nil {
		log.Fatalf("\nCHECK EMAIL EXISTENCE - Cannot execute command : %v\n", err)
		return "", err
	}

	var verificationCode string
	verificationCode = fmt.Sprintf("%04d", rand.Intn(9000)+1000) // Generate a new verification code

	if count > 0 {
		// Update the existing record
		_, err = db.Exec("UPDATE one_time_verification SET verification_code = $1, modules = $2, update_at = $3 WHERE email = $4",
			verificationCode, modules, Gv.FormattedTimeNowYYYYMMDDHHMM, email)
		if err != nil {
			log.Fatalf("\nUPDATE EMAIL IN DB - Cannot execute command : %v\n", err)
			return "", err
		}
	} else {
		// Insert a new record
		_, err = db.Exec("INSERT INTO one_time_verification (email, verification_code, modules, update_at) VALUES ($1, $2, $3, $4)",
			email, verificationCode, modules, Gv.FormattedTimeNowYYYYMMDDHHMM)
		if err != nil {
			log.Fatalf("\nINSERT EMAIL TO DB - Cannot execute command : %v\n", err)
			return "", err
		}
	}

	return verificationCode, nil
}


// Set Account Status
func SetUserAccountStatus(userId string, status bool) error {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE users_login SET user_status = $1 WHERE user_id = $2`

	_, err := db.Exec(sqlStatement, status, userId)

	if err != nil {
		log.Fatalf("\nSET USER ACCOUNT STATUS  - Cannot execute command : %v\n", err)
		return err
	}

	return nil
}

func CheckUserAccountStatusForLogin(userId string) (bool, error) {
	db := config.CreateConnection()
	defer db.Close()

	// Query the user_status from the database
	var userStatus bool
	err := db.QueryRow("SELECT user_status FROM users_login WHERE user_id = $1", userId).Scan(&userStatus)
	if err != nil {
		log.Fatalf("\nCHECK USER ACCOUNT STATUS - Cannot execute query: %v\n", err)
		Lg.Error(" this userId: ", userId+" facing QueryRow Error:"+err.Error())
		return false, err
	}

	// Check the user_status and return accordingly
	if !userStatus {
		Lg.Info("This userId: " + userId + " user status is false")
		return false, nil
	}

	Lg.Info("This userId: " + userId + " user status is true")
	return true, nil
}
