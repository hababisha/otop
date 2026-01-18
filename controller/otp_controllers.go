package controller

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hababisha/otop/models"
)

var store = make(map[string]models.Otp)

func GenerateSixDigitOTP() string{
    return fmt.Sprintf("%06d", rand.Intn(1000000))
}
	
func GenerateFourDigitOTP() string{
    return fmt.Sprintf("%04d", rand.Intn(10000))
}

func GenerateOTP(c *gin.Context){
    var req struct {
        PhoneNumber string `json:"phonenumber" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "phone number is required",
        })
        return
    }

    otp := GenerateFourDigitOTP() //TODO:change based on pref..can be preference setting
    expiresAt := time.Now().Add(5 * time.Minute)//TODO: expiry time also based on pref

    store[req.PhoneNumber] = models.Otp{
        Value: otp,
        ExpiresAt: expiresAt,
        Used: false,
    }

    log.Println("OTP for ", req.PhoneNumber, ":", otp)


    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "expires_at": expiresAt,
    })
} //add timestamp to the otp generated and send with theoption provided

func VerifyOTP(c *gin.Context){
    var req struct{
        PhoneNumber string `json:"phonenumber" binding:"required"`
        Otp string `json:"otp" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "bad request",
        })
    }

    //check if the otp exists
    storedOtp, exists := store[req.PhoneNumber]     
    if !exists {
        c.JSON(http.StatusOK, gin.H{"valid": false})
        return
    }
    //check if it's not expired
    if time.Now().After(storedOtp.ExpiresAt){
        c.JSON(http.StatusOK, gin.H{"valid": false})
        return
    }
    //check if it's nto used 
    if storedOtp.Used {
        c.JSON(http.StatusOK, gin.H{"valid": false})
        return
    }
    //check if the values match

    if storedOtp.Value != req.Otp{
        c.JSON(http.StatusOK, gin.H{"valid": false})
        return
    }

    storedOtp.Used = true
    store[req.PhoneNumber] = storedOtp

    c.JSON(http.StatusOK, gin.H {
        "valid" : true,
    })

}

