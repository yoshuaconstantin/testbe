package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/lib/pq" // postgres golang driver

	Aunth "testbe/globalvariable/authenticator"
	"testbe/module"
	"testbe/schemas/request"
	"testbe/schemas/response"
)

// Insert Feedback User func
func InsertFeedbackUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var GetDataRequest request.RequestInsertFeedback
	err := json.NewDecoder(r.Body).Decode(&GetDataRequest)
	if err != nil {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, GetDataRequest.Token)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	_, errGetDat := module.GetUserProfileDataFromDB(GetUserIdAunth)

	if errGetDat != nil {

		http.Error(w, errGetDat.Error(), http.StatusInternalServerError)
		return
	}

	// InsertUserFeedback, errInsertFeedback := module.InsertFeedbackUserToDB(GetUserIdAunth, *GetProfileData[0].Nickname, GetDataRequest.Comment)

	// if errInsertFeedback != nil {

	// 	http.Error(w, errInsertFeedback.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if !InsertUserFeedback {

	// 	http.Error(w, "Failed to insert feedback", http.StatusInternalServerError)
	// 	return
	// }

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Update Comment Feedback User func
func UpdateCommentFeedbackUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var GetDataRequest request.RequestEditFeedback
	err := json.NewDecoder(r.Body).Decode(&GetDataRequest)
	if err != nil {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, GetDataRequest.Token)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	EditFeedback, errEditFeedback := module.EditFeedBackUserFromDB(GetDataRequest.Id, GetDataRequest.Comment, GetUserIdAunth)

	if errEditFeedback != nil {

		http.Error(w, errEditFeedback.Error(), http.StatusInternalServerError)
		return
	}

	if !EditFeedback {

		http.Error(w, "Failed to Edit feedback", http.StatusInternalServerError)
		return
	}

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Get All feedback data
func GetAllFeedbackUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryParams := r.URL.Query()

	tokenParam := queryParams.Get("token")
	indexParam := queryParams.Get("index")

	index, errI := strconv.Atoi(indexParam)

	if errI != nil {

		http.Error(w, errI.Error(), http.StatusBadRequest)
		return
	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, tokenParam)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	var offset = index * 10

	GetAllFeedbackData, errGetFeedbackData := module.GetFeedBackUserDataFromDB(GetUserIdAunth, offset)

	if errGetFeedbackData != nil {

		http.Error(w, errGetFeedbackData.Error(), http.StatusInternalServerError)
		return
	}

	var response response.ResponseGetAllFeedBackUser
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = GetAllFeedbackData

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Delete Users Feedback func
func DeletUserFeedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var GetDataRequest request.RequestDeleteSingleFeedbackData
	err := json.NewDecoder(r.Body).Decode(&GetDataRequest)
	if err != nil {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, GetDataRequest.Token)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	DeleteSingleFeedBackUser, errDeleteSinglFdbck := module.DeleteFeedBackUserFromDB(GetDataRequest.Id, GetUserIdAunth)

	if errDeleteSinglFdbck != nil {

		http.Error(w, errDeleteSinglFdbck.Error(), http.StatusInternalServerError)
		return
	}

	if !DeleteSingleFeedBackUser {

		http.Error(w, "Cannot delete this feedback", http.StatusInternalServerError)
		return
	}

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
