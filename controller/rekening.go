package controller

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