package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

const (
	//Set constants separately chromedriver.exe Address and local call port of
	seleniumPath = "/usr/local/bin/chromedriver"
	port         = 9515
)

func main() {
	//1. Enable selenium service
	//Set the option of the selium service to null. Set as needed.
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//Delay service shutdown
	defer service.Stop()

	//2. Call browser
	//Set browser compatibility. We set the browser name to chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//Call browser urlPrefix: test reference: defaulturlprefix =“ http://127.0.0.1 :4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	//Delay exiting chrome
	defer wd.Quit()

	//3. Operate on page elements
	//Get Baidu page
	if err := wd.Get("https://www.yatra.com/"); err != nil {
		panic(err)
	}
	//Find Baidu input box id
	we, err := wd.FindElement(selenium.ByID, "booking_engine_hotels")
	if err != nil {
		panic(err)
	}
	//Send '' to input box
	err = we.Click()
	if err != nil {
		panic(err)
	}

	wq, err := wd.FindElement(selenium.ByID, "BE_hotel_htsearch_btn")
	if err != nil {
		panic(err)
	}
	//Send '' to input box
	err = wq.Click()
	if err != nil {
		panic(err)
	}

	//Find Baidu submit button id
	//we, err = wd.FindElement(selenium.ByID, "su")
	//if err != nil {
	//	panic(err)
	//}
	//Click Submit
	//err = we.Click()
	//if err != nil {
	//	panic(err)
	//}

	//Quit after 20 seconds of sleep
	time.Sleep(20 * time.Second)
}
