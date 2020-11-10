package repo

type Item struct {
	Id       int    `json:"id,string"`
	Date     string `json:"date"`
	Duration int    `json:"duration,string"`
}

type ReportTotal struct {
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

type Report struct {
	Days  int
	Since string
	Total ReportTotal
}

type PageData struct {
	Item   Item
	Items  []Item
	Report Report
}

func MsToReportTotal(secs int) ReportTotal {
	mins := secs / 60.0
	hours := mins / 60.0
	days := hours / 24.0
	return ReportTotal{days, hours, mins, secs}
}
