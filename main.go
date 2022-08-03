package main

import (
	"fmt"
	"log"
	"main/ical"
	"net/http"
)

func makeFile(name string) {
	ical.InitFile(name)
	//ical.WriteEvent(name, ical.Event{SUMMARY: "Joel", UID: 22, DTSTAMP: "20220728T205217Z", DTSTART: "20220916T203000Z", DTEND: "20220917T060000Z"})
	//ical.WriteEvent(name, ical.Event{SUMMARY: "Simon", UID: 20, DTSTAMP: "20220728T205217Z", DTSTART: "20220916T203000Z", DTEND: "20220917T060000Z"})
	ical.WriteEvents(name, []ical.Event{{SUMMARY: "Joel", UID: 22, DTSTAMP: "20220728T205217Z", DTSTART: "20220916T203000Z", DTEND: "20220917T060000Z"},
		{SUMMARY: "Simon", UID: 23, DTSTAMP: "20220728T205217Z", DTSTART: "20220916T203000Z", DTEND: "20220917T060000Z"}})
	ical.FinishFile(name)

}

func calHandler(w http.ResponseWriter, r *http.Request) {
	//Get name and set Header for response
	name := r.URL.Query()["user"][0] + ".ical"
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	w.Header().Set("Content-Type", "text/calendar;charset=UTF-8")

	//Send File
	makeFile(name)
	http.ServeFile(w, r, name)

	//Delete File
	ical.DeleteFile(name)

}

func main() {

	log.Println("Starting Server at :8080")
	http.HandleFunc("/cal", calHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
