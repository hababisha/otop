package models

import "time"


type Otp struct {
	Value string
	ExpiresAt time.Time
	Used bool
}