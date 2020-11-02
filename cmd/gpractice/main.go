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
	minutesFlag := flag.Int("minutes", 0, "practice time in minutes")
	flag.Parse()

	execute(*actionFlag, *dateFlag, *minutesFlag)
}

func execute(action string, date string, minutes int) {
	log.Println(fmt.Sprintf("Executing action [%s] with values: [%v, %v]", action, date, minutes))
	gp := gpractice.GPractice{}

	switch Action(action) {
	case ALL:
		all := gp.GetAll()
		log.Println(fmt.Sprintf("result: %v", all))
	case ADD:
		item := gp.Add(model.Item{date, minutes})
		log.Println(fmt.Sprintf("result: %v", item))
	case GET:
		item := gp.Get(date)
		log.Println(fmt.Sprintf("result: %v", item))
	case DEL:
		res := gp.Delete(model.Item{date, minutes})
		log.Println(fmt.Sprintf("result: %v", res))
	case REPORT:
		report := gp.GetReport()
		log.Println(fmt.Sprintf("result: %v", report))
	default:
		log.Fatalf("Unknown action: %v", action)
		os.Exit(1)
	}
}
