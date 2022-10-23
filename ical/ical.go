package ical

import (
	"fmt"
	"log"
	"os"
)

type Event struct {
	SUMMARY string
	UID     int
	DTSTAMP string
	DTSTART string
	DTEND   string
}

func WriteEvents(filename string, e []Event) {
	for _, element := range e {
		writeEvent(filename, element)
	}
}

func writeEvent(name string, e Event) {

	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content := "BEGIN:VEVENT\r\n" +
		"DTSTAMP:" + e.DTSTAMP + "\r\n" +
		"DTSTART:" + e.DTSTART + "\r\n" +
		"DTEND:" + e.DTEND + "\r\n" +
		"SUMMARY:" + e.SUMMARY + "\r\n" +
		"UID:" + fmt.Sprint(e.UID) + "\r\n" +
		"END:VEVENT\r\n"

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

}

func InitFile(name string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	config := "BEGIN:VCALENDAR\r\n" +
		"PRODID:SIMON\r\n" +
		"CALSCALE:GREGORIAN\r\n" +
		"NAME:scouttools\r\n" +
		"VERSION:2.0\r\n"

	_, err = f.WriteString(config)
	if err != nil {
		log.Fatal(err)
	}

}

func FinishFile(name string) {
	f, err := os.OpenFile(name, os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("END:VCALENDAR")
}

func DeleteFile(name string) {
	err := os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
}
