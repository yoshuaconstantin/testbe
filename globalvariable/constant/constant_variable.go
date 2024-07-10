package constant

const BcryptCost = 14

// Secret key should be hidden!!
var Salt = []byte("AzureKey")
var Key = []byte("0123456789ABCDEF")

const CreateAccountPost = "/api/create-account"
const UserDeleteAccount = "/api/user/delete-account"
const UserUpdatePassword = "/api/user/update-password"
const UserLogin = "/api/login"
const UserLogout = "/api/logout"
const HomeDashboards = "/api/home/dashboards"
const HomeUpdateDashboards = "/api/home/update/dashboards/data"
const HomeUserProfileImage = "/api/home/user/profile/image"
const HomeUserProfile = "/api/home/user/profile"
const HomeUserFeedback = "/api/home/user/feedback"
const CommunityPost = "/api/community/post"
const CommunityPostLike = "/api/community/post/like"
const CommunityPostComment = "/api/community/post/comment"
const CommunityPostCommentGet = "/api/community/post/comment/get-comment"
const RefreshToken = "/api/token-refresh"

var EmailCategories = []string{
	"Passwords",
	"OTP",
	"Tickets",
}

var EmailSubjects = []string{
	"Comnfo - Reset Password",
	"Comnfo - One-Time Password (OTP) Code",
	"Comnfo - Bugs Report Ticket",
}

var EmailBodyTexts = []string{

	`Dear Comnfo User,` +
		`\n\nYou have requested to reset your password. We generate a new password for your account, use the following password to login:` +
		`\n\nNew Generated Password: %s` +
		`\n\nIf you did not initiate this request or have any concerns, please contact our support team immediately.` +
		`\n\n\nBest Regards,` +
		`\nThe Application Team`,

	`Dear Comnfo User,` +
		`\n\nAs part of our commitment to ensuring the utmost security for your account, we have generated a unique One-Time Password (OTP) code for you. Please find the details below:` +
		`\n\nOTP Code: %s` +
		`\n\nTo complete the verification process, kindly enter this code into Community Forum apps. This code is valid for a single use and should be entered promptly.` +
		`\n\nIf you did not initiate this request or have any concerns, please contact our support team immediately.` +
		`\n\nBest Regards,` +
		`\nThe Application Security Team`,

	`Dear Comnfo User,` +
		`\n\nThank you for reporting the bugs in Comnfo. Your contribution helps us improve the platform and enhance the user experience.` +
		`\n\nWe have created a ticket for tracking purposes. Here is your ticket ID: %s.` +
		`\n\nPlease use this ID when inquiring about the status or providing additional information related to the reported issues.` +
		`\nWe appreciate your support in making Comnfo better for everyone.` +
		`\n\nBest regards,` +
		`\nThe Application Security Team`,
}

var EmailBodyHtml = []string{

	"<!DOCTYPE html>" +
		"<html lang='en'>" +
		"<head>" +
		"<meta charset='UTF-8'>" +
		"<meta name='viewport' content='width=device-width, initial-scale=1.0'>" +
		"<title>OTP Code</title>" +
		"<link rel='stylesheet' href='https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css'>" +
		"<style>" +
		"body {" +
		"font-family: 'Arial', sans-serif;" +
		"margin: 0;" +
		"padding: 0;" +
		"background: linear-gradient(45deg, #8e9eab, #eef2f3);" +
		"display: flex;" +
		"align-items: center;" +
		"justify-content: center;" +
		"height: 100vh;" +
		"}" +
		".container {" +
		"width: 800;" +
		"max-width: 500px;" +
		"margin: 0 auto;" +
		"padding: 20px;" +
		"background: linear-gradient(180deg, #ffffff, #f5f5f5);" +
		"border-radius: 25px;" +
		"box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);" +
		"text-align: center;" +
		"}" +
		".header {" +
		"margin-bottom: 20px;" +
		"}" +
		".header h1 {" +
		"font-size: 32px;" +
		"font-weight: bold;" +
		"color: #333;" +
		"margin-bottom: 10px;" +
		"}" +
		".new-password {" +
		"font-size: 60px;" +
		"font-weight: bold;" +
		"color: #333;" +
		"margin-bottom: 20px;" +
		"background: linear-gradient(45deg, #84aadf, #64a5ff);" +
		"color: #fff;" +
		"padding: 10px;" +
		"border-radius: 25px;" +
		"box-shadow: 0 0 5px rgba(0, 0, 0, 0.2);" +
		"}" +
		".message {" +
		"font-size: 17px;" +
		"font-weight: bold;" +
		"color: #555;" +
		"line-height: 1.5;" +
		"margin-bottom: 20px;" +
		"}" +
		".message-subtitle {" +
		"font-size: 16px;" +
		"font-weight: bold;" +
		"color: #555;" +
		"line-height: 1.5;" +
		"margin-bottom: 20px;" +
		"}" +
		".footer-subtitle {" +
		"margin-top: 20px;" +
		"font-size: 13px;" +
		"color: #777;" +
		"}" +
		".footer {" +
		"margin-top: 20px;" +
		"font-size: 12px;" +
		"color: #777;" +
		"}" +
		".footer p {" +
		"margin-bottom: 0;" +
		"}" +
		"</style>" +
		"</head>" +
		"<body>" +
		"<div class='container'>" +
		" <a href=''><img src='https://i.ibb.co/Bn8f67Z/password-verif.png' alt='password-verif' alt='Password' width='300' height='300'></a>" +
		"<div class='header'>" +
		"<h1>Comnfo Reset Password</h1>" +
		"</div>" +
		"<div class ='message-subtitle'>" +
		"<p>You have requested to reset your password. We generate a new password for your account.</p>" +
		"<p>Use the following password to login</p>" +
		"</div>" +
		"<div class='new-password'>" +
		"%s" +
		"</div>" +
		"<div class='message'>" +
		"Remember to change your password immediately." +
		"</div>" +
		"<div class='footer-subtitle'>" +
		"<p>If you did not initiate this request or have any concerns, please contact our support team immediately</p>" +
		"<p>Do not share this password with anyone.</p>" +
		"</div>" +
		"<div class='footer'>" +
		"<p>Copyright &copy; 2023 Team Azure.</p>" +
		"</div>" +
		"</div>" +
		"</body>" +
		"</html>",

	"<!DOCTYPE html>" +
		"<html lang='en'>" +
		"<head>" +
		"<meta charset='UTF-8'>" +
		"<meta name='viewport' content='width=device-width, initial-scale=1.0'>" +
		"<title>OTP Code</title>" +
		"<link rel='stylesheet' href='https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css'>" +
		"<style>" +
		"body {" +
		"font-family: 'Arial', sans-serif;" +
		"margin: 0;" +
		"padding: 0;" +
		"background: linear-gradient(45deg, #8e9eab, #eef2f3);" +
		"display: flex;" +
		"align-items: center;" +
		"justify-content: center;" +
		"height: 100vh;" +
		"}" +
		".container {" +
		"width: 800;" +
		"max-width: 500px;" +
		"margin: 0 auto;" +
		"padding: 20px;" +
		"background: linear-gradient(180deg, #ffffff, #f5f5f5);" +
		"border-radius: 25px;" +
		"box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);" +
		"text-align: center;" +
		"}" +
		".header {" +
		"margin-bottom: 20px;" +
		"}" +
		".header h1 {" +
		"font-size: 32px;" +
		"font-weight: bold;" +
		"color: #333;" +
		"margin-bottom: 10px;" +
		"}" +
		".otp-code {" +
		"font-size: 60px;" +
		"font-weight: bold;" +
		"color: #333;" +
		"margin-bottom: 20px;" +
		"background: linear-gradient(45deg, #df7077, #ff4f5a);" +
		"color: #fff;" +
		"padding: 10px;" +
		"border-radius: 25px;" +
		"box-shadow: 0 0 5px rgba(0, 0, 0, 0.2);" +
		"}" +
		".message {" +
		"font-size: 17px;" +
		"font-weight: bold;" +
		"color: #555;" +
		"line-height: 1.5;" +
		"margin-bottom: 20px;" +
		"}" +
		".message-subtitle {" +
		"font-size: 16px;" +
		"font-weight: bold;" +
		"color: #555;" +
		"line-height: 1.5;" +
		"margin-bottom: 20px;" +
		"}" +
		".footer-subtitle {" +
		"margin-top: 20px;" +
		"font-size: 13px;" +
		"color: #777;" +
		"}" +
		".footer {" +
		"margin-top: 20px;" +
		"font-size: 12px;" +
		"color: #777;" +
		"}" +
		".footer p {" +
		"margin-bottom: 0;" +
		"}" +
		"</style>" +
		"</head>" +
		"<body>" +
		"<div class='container'>" +
		"<a href=''><img src='https://i.ibb.co/8K2C35S/Mail-sent-removebg-preview.png' alt='Mail-sent' width='300' height='300'></a>" +
		"<div class='header'>" +
		"<h1>Comnfo OTP Code</h1>" +
		"</div>" +
		"<div class='otp-code'>" +
		"%s" +
		"</div>" +
		"<div class='message'>" +
		"This code is valid for a single use." +
		"</div>" +
		"<div class='footer'>" +
		"<p>If you did not initiate this request or have any concerns, please contact our support team immediately</p>" +
		"<p>Do not share this OTP with anyone.</p>" +
		"<p>Copyright &copy; 2023 Team Azure.</p>" +
		"</div>" +
		"</div>" +
		"</body>" +
		"</html>",

	"<!DOCTYPE html>" +
		"<html lang='en'>" +
		"<head>" +
		"<meta charset='UTF-8'>" +
		"<meta name='viewport' content='width=device-width, initial-scale=1.0'>" +
		"<title>Report Ticket</title>" +
		"<link rel='stylesheet' href='https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css'>" +
		"<style>" +
		"body {" +
		"font-family: 'Arial', sans-serif;" +
		"margin: 0;" +
		"padding: 0;" +
		"background: linear-gradient(45deg, #8e9eab, #eef2f3);" +
		"display: flex;" +
		"align-items: center;" +
		"justify-content: center;" +
		"height: 100vh;" +
		"}" +
		".container {" +
		"width: 800;" +
		"max-width: 500px;" +
		"margin: 0 auto;" +
		"padding: 20px;" +
		"background: linear-gradient(180deg, #ffffff, #f5f5f5);" +
		"border-radius: 25px;" +
		"box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);" +
		"text-align: center;" +
		"}" +
		".header {" +
		"margin-bottom: 20px;" +
		"}" +
		".header h1 {" +
		"font-size: 32px;" +
		"font-weight: bold;" +
		"color: #333;" +
		"margin-bottom: 10px;" +
		"}" +
		".ticket_id {" +
		"font-size: 33px;" +
		"font-weight: bold;" +
		"color: #333;" +
		"margin-bottom: 20px;" +
		"background: linear-gradient(45deg, #9496af, #6a6c7e);" +
		"color: #fff;" +
		"padding: 10px;" +
		"border-radius: 25px;" +
		"box-shadow: 0 0 5px rgba(0, 0, 0, 0.2);" +
		"}" +
		".message {" +
		"font-size: 17px;" +
		"font-weight: bold;" +
		"color: #555;" +
		"line-height: 1.5;" +
		"margin-bottom: 20px;" +
		"}" +
		".message-subtitle {" +
		"font-size: 16px;" +
		"font-weight: bold;" +
		"color: #555;" +
		"line-height: 1.5;" +
		"margin-bottom: 20px;" +
		"}" +
		".footer-subtitle {" +
		"margin-top: 20px;" +
		"font-size: 13px;" +
		"color: #777;" +
		"}" +
		".footer {" +
		"margin-top: 20px;" +
		"font-size: 12px;" +
		"color: #777;" +
		"}" +
		".footer p {" +
		"margin-bottom: 0;" +
		"}" +
		"</style>" +
		"</head>" +
		"<body>" +
		"<div class='container'>" +
		" <a href=''><img src='https://cdni.iconscout.com/illustration/premium/thumb/developer-working-on-bug-fixing-5359921-4493613.png' alt='password-verif' alt='Password' width='500' height='300'></a>" +
		"<div class='header'>" +
		"<h1>Comnfo Report Ticket</h1>" +
		"</div>" +
		"<div class ='message-subtitle'>" +
		"<p>You have Submitted a report bug regarding the apps behaviour.</p>" +
		"<p>We give you Ticket ID to keep track of your submitted bugs.</p>" +
		"<p>Here is your Ticket ID</p>" +
		"</div>" +
		"<div class='ticket_id'>" +
		"%s" +
		"</div>" +
		"<div class='message'>" +
		"You can check more information using this Ticket ID later." +
		"</div>" +
		"<div class='footer-subtitle'>" +
		"<p>By reporting bugs, you're helping us build a better app for all users. Thanks!</p>"+
		"<br>" +
		"<p>If you did not initiate this request or have any concerns, please contact our support team immediately</p>" +
		"<p>Do not share this password with anyone.</p>" +
		"</div>" +
		"<div class='footer'>" +
		"<p>Copyright &copy; 2023 Team Azure.</p>" +
		"</div>" +
		"</div>" +
		"</body>" +
		"</html>",
}
