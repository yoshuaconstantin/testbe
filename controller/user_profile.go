package controller

import (
	"testbe/globalvariable/authenticator"
	Aunth "testbe/globalvariable/authenticator"
	"testbe/globalvariable/function"
	Lg "testbe/logging"
	"testbe/module"
	"testbe/schemas/models"
	"testbe/schemas/request"
	"testbe/schemas/response"
	"encoding/json"
	"net/http"
)

// Update Display Name from profile page
func UpdateDisplayNameProfilePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipAddr := r.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the UpdateDisplayNameProfilePage")

	UserId, err := authenticator.CheckJWTnGetUserId(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error validating JWT : ", ipAddr+" error: "+err.Error())
		return
	}

	queryParams := r.URL.Query()

	DisplayNameParam := queryParams.Get("displayName")

	errCheckQueryParam := function.CheckQueryParameters(DisplayNameParam)

	if errCheckQueryParam != nil {
		http.Error(w, errCheckQueryParam.Error(), http.StatusBadRequest)
		Lg.Error("This ip: ", ipAddr+" has not fill the query parameters with error :"+errCheckQueryParam.Error())
		return
	}

	error := module.UpdateProfileDisplayNameToDB(UserId, DisplayNameParam)

	if error != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error updating new display name, ip: ", ipAddr+" error: "+error.Error())
		return
	}

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success updating display name"

	Lg.Info("Successfully updated new display name, ip: " + ipAddr)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Update Short Bio from profile page
func UpdateShortBioProfilePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipAddr := r.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the UpdateShortBioProfilePage")

	UserId, err := authenticator.CheckJWTnGetUserId(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error validating JWT : ", ipAddr+" error: "+err.Error())
		return
	}

	queryParams := r.URL.Query()

	ShortBioParam := queryParams.Get("shortBio")

	errCheckQueryParam := function.CheckQueryParameters(ShortBioParam)

	if errCheckQueryParam != nil {
		http.Error(w, errCheckQueryParam.Error(), http.StatusBadRequest)
		Lg.Error("This ip: ", ipAddr+" has not fill the query parameters with error :"+errCheckQueryParam.Error())
		return
	}

	error := module.UpdateProfileShortBioToDB(UserId, ShortBioParam)

	if error != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error updating new short bio, ip: ", ipAddr+" error: "+error.Error())
		return
	}

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success updating short bio"

	Lg.Info("Successfully updated new short bio, ip: " + ipAddr)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Change passowr
func UpdatePasswordProfilePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipAddr := r.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the UpdatePasswordProfilePage")

	UserId, err := authenticator.CheckJWTnGetUserId(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error validating JWT : ", ipAddr+" error: "+err.Error())
		return
	}

	queryParams := r.URL.Query()

	NewPasswordParam := queryParams.Get("newPassword")

	errCheckQueryParam := function.CheckQueryParameters(NewPasswordParam)

	if errCheckQueryParam != nil {
		http.Error(w, errCheckQueryParam.Error(), http.StatusBadRequest)
		Lg.Error("This ip: ", ipAddr+" has not fill the query parameters with error :"+errCheckQueryParam.Error())
		return
	}

	hashedPassword, errhashed := authenticator.ValidatePasswordToEncrypt(NewPasswordParam)

	if errhashed != nil {
		http.Error(w, errhashed.Error(), http.StatusConflict)
		Lg.Error("This ip: ", ipAddr+" is facing error: "+errhashed.Error())
		return
	}

	error := module.UpdateProfileNewPasswordToDB(UserId, hashedPassword)

	if error != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error updating new password, ip: ", ipAddr+" error: "+error.Error())
		return
	}

	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success updating account password"

	Lg.Info("Successfully updated new account password, ip: " + ipAddr)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

/*
<Documentation>
	Step-by-step how to use this (currently untested yet)
	- Upload an img byte from apps to UploadImage endpoint -> if success then server will return response with ImageUrl to user
	- Hit InsertDataProfile endpoint and store the remaining data with stored ImageUrl from Upload
	- Leave / Refresh the apps then hit GetDataProfile to get the whole data
** Alt.Step : Hit only UploadImage to get image only to database (untested, unchecked flow)
</Documentation>
*/

// Upload image to local storage outside of source code with return to user ImageUrl string
func UploadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read the binary data from the request body
	var imageData request.RequestUploadImageData
	if err := json.NewDecoder(r.Body).Decode(&imageData); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, imageData.Token)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	// Func Convert from byte to image and return string.
	UploadImageToDB, errUploadImageToDB := module.UploadUserProfilePhotoToDB(GetUserIdAunth, imageData.Data)

	if errUploadImageToDB != nil {

		http.Error(w, errUploadImageToDB.Error(), http.StatusNotAcceptable)
		return
	}

	if !UploadImageToDB {

		http.Error(w, "Failed to upload Image Url to database", http.StatusInternalServerError)
		return
	}

	// Send back the response to user with file name
	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success Upload Image"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Update image profile, replace the string
func UpdateImageProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read the binary data from the request body
	var imageData request.RequestUpdateProfileImageData
	if err := json.NewDecoder(r.Body).Decode(&imageData); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, imageData.Token)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	// Func Convert from byte to image and return string.
	GetNewImageUrl, errGetImageUrl := module.ConvertByteToImgString(imageData.Data, GetUserIdAunth)

	if errGetImageUrl != nil {

		http.Error(w, errGetImageUrl.Error(), http.StatusNotAcceptable)
		return
	}

	UpdateUsersProfileImage, errUpdateProfileImage := module.UpdateUserProfileImageFromDB(GetUserIdAunth, GetNewImageUrl, imageData.OldImageUrl)

	if errUpdateProfileImage != nil {

		http.Error(w, errUpdateProfileImage.Error(), http.StatusInternalServerError)
		return
	}

	if !UpdateUsersProfileImage {

		http.Error(w, "An error occured when updating users profile image", http.StatusInternalServerError)
		return
	}

	// Send back the response to user with file name
	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success Update Image"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Delete image profile, replace with empty string
func DeleteImageProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Read the binary data from the request body
	var imageData request.RequestDeleteProfileImageData
	if err := json.NewDecoder(r.Body).Decode(&imageData); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, imageData.Token)

	if errAunth != nil {

		http.Error(w, errAunth.Error(), AunthStatus)
		return
	}

	DeleteUserProfileImage, errDeleteProfileImage := module.DeleteUserImageProfileFromDB(GetUserIdAunth, imageData.OldImageUrl)

	if errDeleteProfileImage != nil {

		http.Error(w, errDeleteProfileImage.Error(), http.StatusInternalServerError)
		return
	}

	if !DeleteUserProfileImage {

		http.Error(w, "An error occured when deleting users profile image", http.StatusInternalServerError)
		return
	}

	// Send back the response to user with file name
	var response response.GeneralResponseNoData
	response.Status = http.StatusOK
	response.Message = "Success Update Image"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Insert profile data into database
// func InsertDataProfile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	var userData request.RequestInsertProfileData

// 	err := json.NewDecoder(r.Body).Decode(&userData)
// 	if err != nil {

// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	GetUserIdAunth, AunthStatus, errAunth := Aunth.SecureAuthenticator(w, r, userData.Token)

// 	if errAunth != nil {

// 		http.Error(w, errAunth.Error(), AunthStatus)
// 		return
// 	}

// 	data := userData.Data[0]

// 	// Create a UserProfileData instance using the data from the InsertProfileData struct
// 	userProfileData := models.UserProfileDataModel{
// 		Nickname: &data.Nickname,
// 		Age:      &data.Age,
// 		Gender:   &data.Gender,
// 		ImageUrl: &data.ImageUrl,
// 	}

// 	insertDataProfile, errInsertUserProfile := module.UpdateUserProfileToDatabase(userProfileData, GetUserIdAunth)

// 	if errInsertUserProfile != nil {

// 		logging.InsertLog(r, constant.HomeUserProfile, errInsertUserProfile.Error(), userData.Token, http.StatusInternalServerError, 2, 3)

// 		http.Error(w, errInsertUserProfile.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var response response.GeneralResponseNoData
// 	response.Status = http.StatusOK
// 	response.Message = insertDataProfile

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

// Get the profile data with path to local dir
func GetDataProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ipAddr := r.RemoteAddr

	Lg.Info("this ip: " + ipAddr + " has accessed to the GetDataProfile")

	UserId, err := authenticator.CheckJWTnGetUserId(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		Lg.Error("Error validating JWT, ip: ", ipAddr+" error: "+err.Error())
		return
	}

	var dataModel *models.GetUserProfileDataModel

	dataModel, errGetDataProfile := module.GetUserProfileDataFromDB(UserId)

	if errGetDataProfile != nil {
		Lg.Error("This user facing error when getting data Profile, ip: ", ipAddr+" error: "+errGetDataProfile.Error())
		http.Error(w, errGetDataProfile.Error(), http.StatusInternalServerError)
		return
	}

	Lg.Info("User with nicname: "+ dataModel.DisplayName +" successfully get data profile")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataModel)
}
