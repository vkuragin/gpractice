package repo

import (
	"log"
	"time"
)

const (
	DateFormat = "2006-01-02"
)

// Practice item
type Item struct {
	Id       int
	Date     string
	Duration int
}

// Report contains items and a summary: total days, since first date, total duration
type Report struct {
	Items     []Item
	Days      int
	DateStart string
	DateEnd   string
	Total     int
}

// Practice item DTO
type ItemDto struct {
	Id          int    `json:"id,string"`
	Date        string `json:"date"`
	Duration    int    `json:"duration,string"`
	DurationStr string `json:"duration-str"`
}

// Report DTO contains summary: total days, since first date, total duration
type ReportDto struct {
	Items     []ItemDto
	Days      int
	DateStart string
	DateEnd   string
	Total     int
	TotalStr  string
}

// PageData - data to be rendered
type PageData struct {
	Item   ItemDto
	Items  []ItemDto
	Report ReportDto
}

func ItemToDto(i Item) ItemDto {
	return ItemDto{
		Id:          i.Id,
		Date:        i.Date,
		Duration:    i.Duration,
		DurationStr: SecToDurationStr(i.Duration),
	}
}

func DtoToItem(dto ItemDto) Item {
	duration := dto.Duration
	if duration == 0 {
		d, err := time.ParseDuration(dto.DurationStr)
		if err != nil {
			log.Printf("Failed to parse duration: %s\n", dto.DurationStr)
		}
		duration = int(d.Seconds())
	}

	return Item{
		Id:       dto.Id,
		Date:     dto.Date,
		Duration: duration,
	}
}

func ReportToDto(r Report) ReportDto {
	return ReportDto{
		Days:      r.Days,
		DateStart: r.DateStart,
		DateEnd:   r.DateEnd,
		Total:     r.Total,
		TotalStr:  SecToDurationStr(r.Total),
	}
}

func SecToDurationStr(s int) string {
	return time.Duration(s * 1e9).String()
}
