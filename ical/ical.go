package ical

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//TODO: add proper logging

type Event struct {
	SUMMARY string //Title of event
	UID     int    //ID of event
	DTSTAMP string
	DTSTART string
	DTEND   string
}

type Calendar struct {
	Port   int
	Config CalendarConfig
}

type CalendarConfig struct {
	PRODID   string
	CALSCALE string
	NAME     string
	VERSION  float32
}

func writeEvent(file *os.File, e Event) {

	//TODO: maybe redo Event
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

func (calendar *Calendar) initFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	globalconfig := "BEGIN:VCALENDAR\r\n" +
		"PRODID:" + calendar.Config.PRODID + "\r\n" +
		"CALSCALE:" + calendar.Config.CALSCALE + "\r\n" +
		"NAME:" + calendar.Config.NAME + "\r\n" +
		"VERSION:" + fmt.Sprint(calendar.Config.VERSION) + "\r\n"

	_, err = file.WriteString(globalconfig)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (calendar *Calendar) MakeFile(name string, events []Event) *Calendar {
	file := calendar.initFile(name)
	defer file.Close()

	for _, element := range events {
		writeEvent(file, element)
	}

	file.WriteString("END:VCALENDAR")

	return calendar
}

func (calendar *Calendar) DeleteFile(name string) {
	if err := os.Remove(name); err != nil {
		log.Fatal(err)
	}
}

func calHandler(w http.ResponseWriter, r *http.Request) {
	//Get name and set Header for response
	user := r.URL.Query()["user"][0]
	name := fmt.Sprint(user + ".ical")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	w.Header().Set("Content-Type", "text/calendar;charset=UTF-8")

	//Generate File
	//TODO: needs own function and has to be feed from DB
	Cal.MakeFile(name, []Event{{SUMMARY: "Joel", UID: 22, DTSTAMP: "20220728T205217Z", DTSTART: "20220916T203000Z", DTEND: "20220917T060000Z"},
		{SUMMARY: "Simon", UID: 23, DTSTAMP: "20220728T205217Z", DTSTART: "20220916T203000Z", DTEND: "20220917T060000Z"}})

	//Send File
	http.ServeFile(w, r, name)

	//Delete File
	Cal.DeleteFile(name)
	//maybe have one buffer file which is always overwritten

}

func (calendar *Calendar) NewCalendarHTTPConnection() error {
	http.HandleFunc("/cal", calHandler)

	err := http.ListenAndServe((":" + fmt.Sprint(calendar.Port)), nil)
	if err != nil {
		return err
	}

	return nil
}
