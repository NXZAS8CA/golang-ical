package main

import (
	"fmt"
	"log"
	"main/ical"
)

//This file is just an placeholder for the usage of golang-ical
//The caller has to provide the name and the needed events in proper format in order to work

//CALSCALE is per default Gregorian, could be removed from file
//VERSION is mandatory to be 2.0

func main() {

	ical.Cal.SetConfig(ical.CalendarConfig{PRODID: "SIMON", CALSCALE: "GREGORIAN", NAME: "scouttools", VERSION: 2.0})
	ical.Cal.SetHTTPPort(8080)

	log.Println("Starting Server at Port:" + fmt.Sprint(ical.Cal.Port))
	err := ical.Cal.NewCalendarHTTPConnection()
	if err != nil {
		log.Fatal(err)
	}

}
