package models

import "time"

type UserPayload struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type SetOtpPayload struct {
	PhoneNumber   string    `json:"phone_number" binding:"required"`
	Otp           string    `json:"otp"`
	OtpExpiryTime time.Time `json:"otp_expiry_time"`
}

type GetOtpPayload struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Otp         string `json:"otp" binding:"required"`
}
