package domain

import "time"

type InitRegistration struct {
	InitRegId      int
	PhoneNum       string
	OTP            string
	OTPExpiredDate time.Time
	OTPSendVia     string
	IsOTPMatched   int
	OTPMatchedDate time.Time
	SecretQuestion string
	SecretAnswer   string
}
