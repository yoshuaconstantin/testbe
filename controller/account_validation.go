package controller

import (
	"testbe/globalvariable/authenticator"
	Auth "testbe/globalvariable/authenticator"
	"testbe/globalvariable/function"
	Lg "testbe/logging"
	"net/http"
)

func ValidatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipAddr := r.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the ValidatePassword")

	UserId, err := authenticator.CheckJWTnGetUserId(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error validating JWT : ", ipAddr+" error: "+err.Error())
		return
	}

	queryParams := r.URL.Query()

	OldPasswordParam := queryParams.Get("oldPassword")

	// Check the query parameters status
	errCheckQueryParam := function.CheckQueryParameters(OldPasswordParam)

	if errCheckQueryParam != nil {
		http.Error(w, errCheckQueryParam.Error(), http.StatusBadRequest)
		Lg.Error("This ip: ", ipAddr+" has not fill the query parameters with error :"+errCheckQueryParam.Error())
		return
	}

	storedPassword, errStrPswd := Auth.ValidateGetStoredPasswordByUserId(UserId)

	if errStrPswd != nil {
		http.Error(w, errStrPswd.Error(), http.StatusUnauthorized)
		Lg.Error("Error validating get stored password by userId, ip : ", ipAddr+" error: "+errStrPswd.Error())
		return
	}

	PasswordValidation, errPassvalidate := Auth.ValidateUserPassword(OldPasswordParam, storedPassword)

	if errPassvalidate != nil {
		http.Error(w, errPassvalidate.Error(), http.StatusBadRequest)
		Lg.Error("Error validating passowrd, ip : ", ipAddr+" error: "+errPassvalidate.Error())
		return
	}

	if !PasswordValidation {
		http.Error(w, string("Password not match!"), http.StatusBadRequest)
		Lg.Error("Error validating passowrd, ip : ", ipAddr+" error: password not match!")
		return
	}

	w.WriteHeader(http.StatusOK)
}
