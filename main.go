package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func main() {
	// Replace with your own email and password
	email := "rootnj"
	password := "Marune@123!"

	// Launch a new browser instance with custom user agent
	url := launcher.New().
		Headless(false). // Headless mode set to false to see the browser window
		NoSandbox(true). // Disable sandbox for some environments
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36").
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// Create a new page with stealth mode enabled
	page := stealth.MustPage(browser)

	// Navigate to the login page
	page.MustNavigate("https://zorogaming.admin.enes.tech/authentication/login")

	// Wait for the username field to be present
	page.MustElement(`body > cap-root > cap-authentication > main > cap-login > div > div.login__body > cap-login-form > div > form > div.eui-fieldset > div:nth-child(1) > rmn-form-control > div > input[type=text]`)

	// Fill in the login form using JS paths
	page.MustElement(`body > cap-root > cap-authentication > main > cap-login > div > div.login__body > cap-login-form > div > form > div.eui-fieldset > div:nth-child(1) > rmn-form-control > div > input[type=text]`).MustInput(email)
	page.MustElement(`body > cap-root > cap-authentication > main > cap-login > div > div.login__body > cap-login-form > div > form > div.eui-fieldset > div:nth-child(2) > rmn-password-input-control > div > input[type=password]`).MustInput(password)
	page.MustElement(`body > cap-root > cap-authentication > main > cap-login > div > div.login__body > cap-login-form > div > form > div.dialog-footer > button`).MustClick()

	// Wait for a specific element that appears after login
	// Replace with an actual selector that appears after a successful login
	page.MustElement(`body > cap-root > cap-main-layout > main > cap-home > div > div.link-card-container > div:nth-child(1)`)

	fmt.Println("Login completed")

	// Take a screenshot after login
	screenshotData := page.MustScreenshot()
	if err := os.WriteFile("screenshot.png", screenshotData, 0644); err != nil {
		log.Fatalf("Failed to save screenshot: %v", err)
	} else {
		fmt.Println("Screenshot saved as screenshot.png")
	}
}
