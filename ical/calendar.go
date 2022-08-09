package ical

var calendarConfig = CalendarConfig{}
var Cal = Calendar{Config: calendarConfig}

func (calendar *Calendar) GetConfig() CalendarConfig {
	return calendar.Config
}

func (calendar *Calendar) SetConfig(config CalendarConfig) *Calendar {
	calendar.Config = config
	return calendar
}

func (calendar *Calendar) GetHTTPort() int {
	return calendar.Port
}

func (calendar *Calendar) SetHTTPPort(Port int) {
	calendar.Port = Port
}
