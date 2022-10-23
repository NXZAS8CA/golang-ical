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

func WriteEvents(f *os.File, e []Event) {
	for _, element := range e {

		content := "BEGIN:VEVENT\r\n" +
			"DTSTAMP:" + element.DTSTAMP + "\r\n" +
			"DTSTART:" + element.DTSTART + "\r\n" +
			"DTEND:" + element.DTEND + "\r\n" +
			"SUMMARY:" + element.SUMMARY + "\r\n" +
			"UID:" + fmt.Sprint(element.UID) + "\r\n" +
			"END:VEVENT\r\n"

		_, err := f.WriteString(content)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func writeEvent(f *os.File, e Event) {

	content := "BEGIN:VEVENT\r\n" +
		"DTSTAMP:" + e.DTSTAMP + "\r\n" +
		"DTSTART:" + e.DTSTART + "\r\n" +
		"DTEND:" + e.DTEND + "\r\n" +
		"SUMMARY:" + e.SUMMARY + "\r\n" +
		"UID:" + fmt.Sprint(e.UID) + "\r\n" +
		"END:VEVENT\r\n"

	_, err := f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

}

func InitFile(name string) (f *os.File) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	config := "BEGIN:VCALENDAR\r\n" +
		"PRODID:SIMON\r\n" +
		"CALSCALE:GREGORIAN\r\n" +
		"NAME:scouttools\r\n" +
		"VERSION:2.0\r\n"

	_, err = f.WriteString(config)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func FinishFile(f *os.File) {

	f.WriteString("END:VCALENDAR")
}

func DeleteFile(name string) {
	err := os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
}
