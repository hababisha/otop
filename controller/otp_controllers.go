package controller

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/hababisha/otop/models"
	"github.com/hababisha/otop/repository"
)

// var store = make(map[string]models.Otp)

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

    // store[req.PhoneNumber] = models.Otp{
    //     Value: otp,
    //     ExpiresAt: expiresAt,
    //     Used: false,
    // }

    if err := repository.CreateOTP(req.PhoneNumber, otp, expiresAt); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create otp"})
        log.Fatal(err)
        return
    }

    log.Println("OTP for ", req.PhoneNumber, ":", otp)


    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "expires_at": expiresAt,
    })
}

func VerifyOTP(c *gin.Context){
    var req struct{
        PhoneNumber string `json:"phonenumber" binding:"required"`
        Otp string `json:"otp" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "bad request",
        })
        return
    }

    // storedOtp, exists := store[req.PhoneNumber]     
    record, err := repository.GetOTP(req.PhoneNumber)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"valid": false})
        return
    }

    if record.Used || time.Now().After(record.ExpiresAt) || record.OTP != req.Otp {
        c.JSON(http.StatusOK, gin.H{"valid": false})
        return 
    }

    _ = repository.MarkOtpAsUsed(req.PhoneNumber, req.Otp)

    c.JSON(http.StatusOK, gin.H {
        "valid" : true,
    })

}

