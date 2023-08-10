package tests

import (
	"cronJob/utils"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendSlackMessage(t *testing.T) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = utils.SendSlackMessage("Test message")

	if err != nil {
		t.Errorf("error sending slack message : %s", err)
	}
}
