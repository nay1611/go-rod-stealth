package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
	log "github.com/sirupsen/logrus"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

// func init() {
//		launcher.NewBrowser().MustGet()
// }

const weeweebrowser = "chromium"

func main() {
	browserLauncher := launcher.New().Bin(weeweebrowser).MustLaunch()
	browser := rod.New().ControlURL(browserLauncher).Timeout(time.Minute).MustConnect()
	defer browser.MustClose()

	log.Infoln("Launched Browser")

	stealthPage := stealth.MustPage(browser)
	normalPage := browser.MustPage("https://bot.sannysoft.com")

	byPassBrowsing(stealthPage, browser)
	normalBrowsing(normalPage, browser)
}

func byPassBrowsing(page *rod.Page, browser *rod.Browser) {
	bypassPage := stealth.MustPage(browser)
	log.Infof("StealthJS hash: %x", md5.Sum([]byte(stealth.JS)))

	bypassPage.MustNavigate("https://bot.sannysoft.com/")
	log.Infoln("visiting testing page with stealth")

	printReport(bypassPage, "stealth")
}

func normalBrowsing(page *rod.Page, browser *rod.Browser) {
	normalPage := browser.MustPage("https://bot.sannysoft.com/")
	log.Infoln("visiting testing page without stealth")

	printReport(normalPage, "normal")
}

func printReport(page *rod.Page, reportMode string) {
	log.Infof("Fetching page report for %s mode", reportMode)

	el := page.MustElement("#broken-image-dimensions.passed")
	
	log.Infoln("Checking for passed elements")
	for _, row := range el.MustParents("table").First().MustElements("tr:nth-child(n+2)") {
		log.Infof("Fetching Elements\t\t")
		cells := row.MustElements("td")
		key := cells[0].MustProperty("textContent")

		if strings.HasPrefix(key.String(), "User Agent") {
			fmt.Printf("\t\t%s: %t\n\n", key, !strings.Contains(cells[1].MustProperty("textContent").String(), "HeadlessChrome/"))
		} else if strings.HasPrefix(key.String(), "Hairline Feature") {
			continue
		} else {
			fmt.Printf("\t\t%s: %s\n\n", key, cells[1].MustProperty("textContent"))
		}
	}

	page.MustScreenshot("")
}

