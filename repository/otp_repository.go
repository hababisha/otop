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

func GetOTP(phonenumber string) (*OTPRecord, error) {
	query := `
		SELECT value, expires_at, used FROM otps WHERE phonenumber=$1 ORDER BY created_at DESC LIMIT 1
	`

	var record OTPRecord
	err := utils.DB.QueryRow(query, phonenumber).Scan(&record.OTP, &record.ExpiresAt, &record.Used)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func MarkOtpAsUsed(phonenumber, otp string) error{
	query := `
		UPDATE otps SET used = true where phonenumber=$1 AND value=$2
	`

	_, err := utils.DB.Exec(query, phonenumber, otp)
	return err
}