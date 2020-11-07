package repo

type Item struct {
	Id       uint64
	Date     string
	Duration uint64
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
	Items  []Item
	Report Report
}

func MsToReportTotal(ms uint64) ReportTotal {
	secs := ms / 1000.0
	mins := secs / 60.0
	hours := mins / 60.0
	days := hours / 24.0
	return ReportTotal{int(days), int(hours), int(mins), int(secs)}
}
