package ical

import (
	"fmt"
	"log"
	"os"
)

//TODO: add proper logging
//TODO: add global config struct

type Event struct {
	SUMMARY string //Title of event
	UID     int    //ID of event
	DTSTAMP string
	DTSTART string
	DTEND   string
}

func writeEvent(file *os.File, e Event) {

	//TODO: make content configurable
	content := "BEGIN:VEVENT\r\n" +
		"DTSTAMP:" + e.DTSTAMP + "\r\n" +
		"DTSTART:" + e.DTSTART + "\r\n" +
		"DTEND:" + e.DTEND + "\r\n" +
		"SUMMARY:" + e.SUMMARY + "\r\n" +
		"UID:" + fmt.Sprint(e.UID) + "\r\n" +
		"END:VEVENT\r\n"

	_, err := file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

}

func initFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	//TODO: make config configurable
	config := "BEGIN:VCALENDAR\r\n" +
		"PRODID:SIMON\r\n" +
		"CALSCALE:GREGORIAN\r\n" +
		"NAME:scouttools\r\n" +
		"VERSION:2.0\r\n"

	_, err = file.WriteString(config)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func DeleteFile(name string) {
	if err := os.Remove(name); err != nil {
		log.Fatal(err)
	}
}

func MakeFile(name string, events []Event) {
	file := initFile(name)
	defer file.Close()

	for _, element := range events {
		writeEvent(file, element)
	}

	file.WriteString("END:VCALENDAR")
}
