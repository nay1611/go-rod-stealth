package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Replace with your own email and password
	email := "gittest"
	password := "Gittest123"

	// Launch a new browser instance with custom user agent
	url := launcher.New().
		Headless(false). // Headless mode set to false to see the browser window
		NoSandbox(true). // Disable sandbox for some environments
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36").
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://zorogaming.admin.enes.tech/authentication/login")

	// Wait for the page to load
	time.Sleep(15 * time.Second) // Adjust as needed

	// Fill in the login form
	page.MustElement(`body > cap-root > cap-authentication > main > cap-login > div > div.login__body > cap-login-form > div > form > div.eui-fieldset > div:nth-child(1) > rmn-form-control > div > input[type=text]`).MustInput(email)
	page.MustElement(`body > cap-root > cap-authentication > main > cap-login > div > div.login__body > cap-login-form > div > form > div.eui-fieldset > div:nth-child(2) > rmn-password-input-control > div > input[type=password]`).MustInput(password)
	page.MustElement(`body > cap-root > cap-authentication > main > cap-login > div > div.login__body > cap-login-form > div > form > div.dialog-footer > button`).MustClick()

	// Wait for login to complete
	page.MustWaitLoad()
	fmt.Println("Login completed")

	// Take a screenshot after login
	screenshotData := page.MustScreenshot()
	if err := os.WriteFile("screenshot.png", screenshotData, 0644); err != nil {
		log.Fatalf("Failed to save screenshot: %v", err)
	} else {
		fmt.Println("Screenshot saved as screenshot.png")
	}

	// Keep the browser open for a while
	time.Sleep(10 * time.Second)
}
