package controller

import (
	"math/rand"
    "fmt"
	"github.com/hababisha/otop/models"
)

var store = make(map[string]models.Otp)

func GenerateSixDigitOTP() string{
    return fmt.Sprintf("%06d", rand.Intn(1000000))
}
	
func GenerateFourDigitOTP() string{
    return fmt.Sprintf("%04d", rand.Intn(10000))
}

func GenerateOTP(){} //add timestamp to the otp generated and send with theoption provided

func VerifyOTP(){}

