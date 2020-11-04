package main

import (
	"flag"
	"fmt"
	"github.com/vk23/gpractice"
	"github.com/vk23/gpractice/model"
	"log"
	"os"
	"time"
)

type Action string

const (
	ALL    Action = "all"
	ADD           = "add"
	GET           = "get"
	DEL           = "del"
	REPORT        = "report"
)

func main() {
	actionFlag := flag.String("action", string(ALL), "one of possible actions: all, add, del, get, report")
	dateFlag := flag.String("date", time.Now().Format("2006-01-02"), "practice date yyyy-MM-dd")
	minutesFlag := flag.Uint64("minutes", 0, "practice time in minutes")
	flag.Parse()

	execute(*actionFlag, *dateFlag, *minutesFlag)
}

func execute(action string, date string, minutes uint64) {
	log.Println(fmt.Sprintf("Executing action [%s] with values: [%v, %v]", action, date, minutes))
	gp := gpractice.GPractice{}

	item := model.Item{Date: date, Duration: uint64(minutes * 60 * 1000)}
	switch Action(action) {
	case ALL:
		all := gp.GetAll()
		log.Println(fmt.Sprintf("result: %v", all))
	case ADD:
		item := gp.Add(item)
		log.Println(fmt.Sprintf("result: %v", item))
	case GET:
		item := gp.Get(date)
		log.Println(fmt.Sprintf("result: %v", item))
	case DEL:
		res := gp.Delete(item)
		log.Println(fmt.Sprintf("result: %v", res))
	case REPORT:
		report := gp.GetReport()
		log.Println(fmt.Sprintf("result: %v", report))
	default:
		log.Fatalf("Unknown action: %v", action)
		os.Exit(1)
	}
}
