package ical

type CalendarConfig struct {
	PRODID   string
	CALSCALE string
	NAME     string
	VERSION  float32
}

//returns the config of the current used calendar
/*func GetCalendarConfig() string {
	out := "BEGIN:VCALENDAR\r\nPRODID:" + CalendarConfig.PRODID + "\r\n" +
		"CALSCALE:" + CalendarConfig.CALSCALE + "\r\n" +
		"NAME:" + CalendarConfig.NAME + "\r\n" +
		"VERSION:" + CalendarConfig.VERSION + "\r\n"
	return out
}
*/
func SetCalendarConfig(prodid string, calscale string, name string, version float32) *CalendarConfig {
	calendarConfig := CalendarConfig{PRODID: prodid, CALSCALE: calscale, NAME: name, VERSION: version}
	return &calendarConfig
}
