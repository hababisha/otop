package repository

import (
	"time"

	"github.com/hababisha/otop/utils"
)

type OTPRecord struct {
	OTP string
	ExpiresAt time.Time
	Used bool
}

func CreateOTP(phonenumber, otp string, expiresAt time.Time) error {
	query := `
		INSERT INTO otps (phonenumber, value, expires_at, used)
		VALUES ($1, $2, $3, false)
	`

	_, err := utils.DB.Exec(query, phonenumber, otp, expiresAt)
	return err
}