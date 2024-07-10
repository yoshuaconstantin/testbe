package controller

import (
	//"bytes"
	"encoding/json"
	"fmt"

	//"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"testbe/globalvariable/authenticator"
	Aunth "testbe/globalvariable/authenticator"
	"testbe/globalvariable/function"
	Lg "testbe/logging"
	jwttoken "testbe/middleware"
	"testbe/module"
	"testbe/schemas/request"
	"testbe/schemas/response"
	//"Comnfo/session"
)

var ipMap = make(map[string]int)

var blockedIPs = make(map[string]time.Time)

func isBlocked(ipAddr string) bool {
	blockedAt, found := blockedIPs[ipAddr]
	if found && time.Since(blockedAt) < 4*time.Hour {
		return true
	}
	return false
}

// Block an IP address
func blockIP(ipAddr string) {
	blockedIPs[ipAddr] = time.Now()
}

// Reset a blocked IP address after 4 hours
func resetIP(ipAddr string) {
	delete(blockedIPs, ipAddr)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Welcome To The Jungle"

	ipAddr := r.RemoteAddr

	Lg.Info("This ip: " + ipAddr + " has been checked out welcome API test")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipAddr := r.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the CreateNewAccount")

	var data request.RequestUserWithProfile

	if err := request.DynamicUnmarshalFromReader(r.Body, &data); err != nil {
		Lg.Error("This ip: ", ipAddr+" is facing this error: "+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if email contains "@"
	if !strings.Contains(data.DataProfile.Email, "@") {
		errorMsg := "Invalid email format: must contain '@'"
		Lg.Error("This ip: ", ipAddr+" is facing this error: "+errorMsg)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	GetTokenAndJwt, errInsert := module.CreateAccountToDB(data)

	if errInsert != nil {

		http.Error(w, errInsert.Error(), http.StatusInternalServerError)
		Lg.Error("Error create account to db, ip: ", ipAddr+" error: "+errInsert.Error())
		return
	}

	var response response.ResponseUserLoginWithJWT
	response.Status = http.StatusOK
	response.Message = "Data user baru telah di tambahkan"
	response.JwtToken = GetTokenAndJwt.JWT

	Lg.Info("Successfully created account to db, ip: " + ipAddr + " with usernames: " + data.DataLogin.Username)

	// kirim response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// AmbilUser mengambil single data dengan parameter id
func GetSnglUsr(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// dapatkan idUser dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// memanggil models GetSingleUser dengan parameter id yg nantinya akan mengambil single data
	user, err := module.GetSingleUser(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data User. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(user)
}

func LoginAccountTest(cr function.CustomWR) {
  
	ipAddr := cr.R.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the LoginAccount")

	IsIpBlocked := isBlocked(ipAddr)

	if IsIpBlocked {
		function.Ehandler(cr, http.StatusUnauthorized, "Your IP has been blocked try again in 4 more hours.", "This ip has been blocked: "+ ipAddr)
		return
	} else {
		resetIP(ipAddr)
	}

	if ipMap[ipAddr] >= 5 {
		// Return an error response indicating that the IP address has been blocked
		blockIP(ipAddr)
		message := "Too many failed login attempts. Your IP has been blocked"
		function.Ehandler(cr, http.StatusUnauthorized, message, message + "ip: "+ipAddr)
		return
	}

	var loginData request.RequestLoginData

	if err := request.DynamicUnmarshalFromReader(cr.R.Body, &loginData); err != nil {
		function.Ehandler(cr, http.StatusBadRequest, err.Error(), "This ip: "+ipAddr+" is facing this error: "+err.Error())

		return
	}

	storedPassword, errStrdPswd := authenticator.ValidateGetStoredPasswordByUsername(loginData.Username)

	if errStrdPswd != nil {
		function.Ehandler(cr, http.StatusInternalServerError, errStrdPswd.Error(), "Password error : "+ errStrdPswd.Error())
		return
	}

	PasswordValidation, errPassvalidate := authenticator.ValidateUserPassword(loginData.Password, storedPassword)

	if errPassvalidate != nil {
		// kirim respon 400 kalau ada error
		ipMap[ipAddr]++

		attemptsLeft := 5 - ipMap[ipAddr]

		errorMsg := fmt.Sprintf("Password not match. You have %d attempts left", attemptsLeft)
		function.Ehandler(cr, http.StatusConflict, errorMsg, "Password not match, ip: "+ipAddr)

		return
	}

	if !PasswordValidation {
		ipMap[ipAddr]++

		attemptsLeft := 5 - ipMap[ipAddr]

		errorMsg := fmt.Sprintf("Password not match. You have %d attempts left", attemptsLeft)
		function.Ehandler(cr, http.StatusConflict, errorMsg, "Password not match, ip: "+ipAddr)

		return
	}

	GetUserID, errGetUuid := authenticator.ValidateUsernameGetUUID(loginData.Username)

	if errGetUuid != nil {
		message := "error got userId, ip : "+ipAddr+" error: "+errGetUuid.Error()
		function.Ehandler(cr, http.StatusInternalServerError, errGetUuid.Error(), message)
		return
	}

	GetUserStatus, errGetUserStatus := Aunth.CheckUserAccountStatusForLogin(GetUserID)

	if errGetUserStatus != nil {
		message := "error got userId, ip : "+ipAddr+" error: "+errGetUserStatus.Error()
		function.Ehandler(cr, http.StatusInternalServerError, errGetUserStatus.Error(), message)
		return
	}

	// Generate JWT Token after Succesfully passed the aunth system
	GenerateJwtToken, errGenerate := jwttoken.GenerateToken(GetUserID)

	if errGenerate != nil {
		message :=  "Error generating JWT : " +ipAddr+" error: "+errGenerate.Error()
		function.Ehandler(cr, http.StatusInternalServerError, errGenerate.Error(), message)
		return
	}

	// jika user status false (belum verifikasi otp), maka akan return json dengan email dan jwt
	if !GetUserStatus {

		GetEmailFromDatabase, errGetEmail := Aunth.ValidateUsnGetEmail(GetUserID)

		if errGetEmail != nil {
			message := "error get email from userId, ip : "+ipAddr+" error: "+errGetEmail.Error()
			function.Ehandler(cr, http.StatusUnauthorized, errGetEmail.Error(), message)
			return
		}
		
		message := "Account hasn't verify yet, but return email and jwt for OTP verification, ip : "+ ipAddr+", email: "+GetEmailFromDatabase
		function.Rsp(cr.W).StatusCode(http.StatusCreated).LogE(message).JSON(response.DataResponseUserLoginWithEmailAndJWT(GetEmailFromDatabase, GenerateJwtToken))
		return
	}

	ipMap[ipAddr] = 0
	message := "Successfully logined with username: " + loginData.Username + " with return of JWT token"

	function.Rsp(cr.W).StatusCode(http.StatusOK).Log(message).JSON(response.DataResponseLogin("success", GenerateJwtToken))
}

func GetAllUsr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call the GetAllUser method from the models package
	users, err := module.GetAllUser()

	if err != nil {
		log.Fatalf("Unable to retrieve data. %v", err)

		var response response.GeneralResponseNoData
		response.Status = http.StatusInternalServerError
		response.Message = "Error retrieving data"

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)

		return
	}

	var response response.ResponseUserLogin
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = users

	// Send the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Update User Password func
func UpdateAccountPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var updatePswdModel request.RequestChangePasswordData

	err := json.NewDecoder(r.Body).Decode(&updatePswdModel)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, updatePswdModel.Token)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	UpdatePswd, errUpdatePswd := module.UpdatePasswordAccountFromDB(GetUserIdAunth, updatePswdModel.Password)

	if errUpdatePswd != nil {

		http.Error(w, errUpdatePswd.Error(), http.StatusInternalServerError)
		return
	}

	if !UpdatePswd {

		http.Error(w, "Cannot change password", http.StatusInternalServerError)
		return
	}

	var response response.GeneralResponseNoData
	response.Message = "Succes"
	response.Status = http.StatusOK

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Delete User func
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryParams := r.URL.Query()

	tokenParam := queryParams.Get("token")

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, tokenParam)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	_, errDeleteUser := module.RemoveAccountFromDB(GetUserIdAunth)

	if errDeleteUser != nil {

		http.Error(w, errDeleteUser.Error(), http.StatusInternalServerError)
		return

	}

	var response response.GeneralResponseNoData
	response.Message = "Delete user operation success"
	response.Status = http.StatusOK

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Logout User func
func LogoutAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryParams := r.URL.Query()

	tokenParam := queryParams.Get("token")

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, tokenParam)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	logoutUser, errLogout := module.LogoutAccountFromDB(GetUserIdAunth)

	if errLogout != nil {

		http.Error(w, errLogout.Error(), http.StatusInternalServerError)
		return
	}

	if !logoutUser {

		http.Error(w, "Error when trying to logout", http.StatusInternalServerError)
		return
	}

	var response response.GeneralResponseNoData
	response.Message = "Logout Succes"
	response.Status = http.StatusOK

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Refresh JWT, place this on init
func RefreshJWT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipAddr := r.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the RefrshToken")

	UserId, err := authenticator.CheckJWTnGetUserId(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error validating JWT : ", ipAddr+" error: "+err.Error())
		return
	}

	GenerateJwt, errGen := jwttoken.GenerateToken(UserId)

	if errGen != nil {
		http.Error(w, errGen.Error(), http.StatusConflict)
		Lg.Error("This ip: ", ipAddr+" is facing error when generating JWT, error: "+err.Error())
		return
	}

	var response response.ResponseUserLoginWithJWT
	response.Message = "Token Refreshed"
	response.Status = http.StatusOK
	response.JwtToken = GenerateJwt

	Lg.Info("Successfully refresh JWT with return of new generated JWT")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Testing Generate JWT
func TestGenerateJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var loginData request.RequestLoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	GenerateJwt, errGen := jwttoken.GenerateToken(loginData.Username)

	if errGen != nil {
		http.Error(w, errGen.Error(), http.StatusConflict)
		return
	}

	var response response.ResponseUserLoginWithJWT
	response.Status = http.StatusOK
	response.Message = "Testing Generate JWT"
	response.JwtToken = GenerateJwt

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Testing Verify JWT
func TestVerifyJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		// If the authorization header is empty, return an error
		var response response.GeneralResponseNoData
		response.Status = http.StatusBadRequest
		response.Message = "Missing authorization header"

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	CheckJwtTokenValidation, userId, erroCheckJWt := jwttoken.VerifyToken(tokenString)

	if erroCheckJWt != nil {
		var response response.GeneralResponseNoData
		response.Status = http.StatusUnauthorized
		response.Message = erroCheckJWt.Error()

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	if !CheckJwtTokenValidation {
		var response response.GeneralResponseNoData
		response.Status = http.StatusUnauthorized
		response.Message = "Unauthorized user"

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = userId

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
