package jobs

import (
	"log"

	"github.com/hababisha/otop/repository"
	"github.com/robfig/cron/v3"
)

func StartOtpCleanup(){
	c := cron.New()
	
	//everyday at 12:00 | can do @every 1m tomake it every 1 min
	_, err := c.AddFunc("0 12 * * *", func(){
		log.Println("runnign otp cleanup cron job")

		if err := repository.CleanOldOtps(); err != nil {
			log.Println("otp cleanup failed", err)
		} else {
			log.Println("otp cleanup completed successfully")
		}
	})

	if err != nil {
		log.Fatal("failed to start otp cron", err)
	}

	c.Start()
}