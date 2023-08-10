package main

import (
	"cronJob/utils"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/robfig/cron"
)

func main() {

	LoadEnvVariables()
	loc, err := time.LoadLocation(os.Getenv("CRON_TIMEZONE"))
	if err != nil {
		panic(err)
	}

	cronJob := cron.NewWithLocation(loc)
	// sec, min, hr, day, month,  day of week
	err = cronJob.AddFunc(os.Getenv("CRON_SCEDULE"), func() {
		slackMessage := "hello!"
		err = utils.SendSlackMessage(slackMessage)
		if err != nil {
			panic(err)
		}
		log.Println("Successfully sent")
		// fmt.Println("hello world", time.Now())
	})

	if err != nil {
		panic(err)
	}

	cronJob.Run()
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
