package repo

// Practice item
type Item struct {
	Id       int    `json:"id,string"`
	Date     string `json:"date"`
	Duration int    `json:"duration,string"`
}

// Report contains summary: total days, since first date, total duration
type Report struct {
	Days  int
	Since string
	Total string
}

// PageData - data to be rendered
type PageData struct {
	Item   Item
	Items  []Item
	Report Report
}
