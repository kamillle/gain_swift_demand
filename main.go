package main

import (
	"log"
	"os"

	"github.com/sclevine/agouti"
)

const loginUrl = "https://www.swiftdemand.com/users/sign_in"

func main() {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			// "--headless",
			"--disable-gpu",
			"--no-sandbox",
		},
		),
		agouti.Debug,
	)

	err := driver.Start()
	if err != nil {
		log.Printf("Failed to start driver: %v", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	err = page.Navigate(loginUrl)
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	identity := page.FindByID("user_email")
	password := page.FindByID("user_password")
	identity.Fill(os.Getenv("SWIFT_DEMAND_ID"))
	password.Fill(os.Getenv("SWIFT_DEMAND_PASSWORD"))

	err = page.FindByClass("submit-button").Submit()
	if err != nil {
		log.Printf("Failed to login: %v", err)
	}

	err = page.FindByButton("Claim").Click()
	if err != nil {
		log.Printf("Failed to Submit: %v", err)
	}

	driver.Stop()
}
