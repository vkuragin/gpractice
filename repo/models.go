package repo

// Practice item
type Item struct {
	Id       int    `json:"id,string"`
	Date     string `json:"date"`
	Duration int    `json:"duration,string"`
}

// Total duration, part of Report
type ReportTotal struct {
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

// Report contains summary: total days, since first date, total duration
type Report struct {
	Days  int
	Since string
	Total ReportTotal
}

// PageData - data to be rendered
type PageData struct {
	Item   Item
	Items  []Item
	Report Report
}

// Converts total duration in seconds to ReportTotal view
func SecondsToReportTotal(secs int) ReportTotal {
	x, mins, hours, days := secs, 0, 0, 0

	x /= 60.0
	if x > 0 {
		mins = x % 60.0
	}

	x /= 60.0
	if x > 0 {
		hours = x % (60.0 * 60.0)
	}

	x /= 24.0
	if x > 0 {
		days = x % (60.0 * 60.0 * 24.0)
	}
	return ReportTotal{days, hours, mins, secs}
}
