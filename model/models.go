package model

type Item struct {
	Date     string
	Duration uint64
}

type ReportTotal struct {
	days    int
	hours   int
	minutes int
}

type Report struct {
	days  int
	since string
	total ReportTotal
}
