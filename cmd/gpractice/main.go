package main

import (
	"flag"
	"fmt"
	"github.com/vk23/gpractice"
	"github.com/vk23/gpractice/repo"
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
	idFlag := flag.Uint64("id", 0, "id")
	flag.Parse()

	execute(*actionFlag, *idFlag, *dateFlag, *minutesFlag)
}

func execute(action string, id uint64, date string, minutes uint64) {
	log.Println(fmt.Sprintf("Executing action [%s] with values: [%v, %v, %v]", action, id, date, minutes))
	gp := initGPractice()

	item := repo.Item{Id: id, Date: date, Duration: uint64(minutes * 60 * 1000)}
	switch Action(action) {
	case ALL:
		all, err := gp.GetAll()
		log.Printf("result: %v\n, error: %v\n", all, err)
	case ADD:
		item, err := gp.Save(item)
		log.Printf("result: %v\n, error: %v\n", item, err)
	case GET:
		item, err := gp.Get(item.Id)
		log.Printf("result: %v\n, error: %v\n", item, err)
	case DEL:
		res, err := gp.Delete(item.Id)
		log.Printf("result: %v\n, error: %v\n", res, err)
	case REPORT:
		report, err := gp.GetReport()
		log.Printf("result: %v\n, error: %v\n", report, err)
	default:
		log.Fatalf("Unknown action: %v\n", action)
		os.Exit(1)
	}
}

func initGPractice() gpractice.GPractice {
	//m := make(map[uint64]model.Item, 0)
	//gp := gpractice.GPractice{Repo: &repo.StubRepo{m}}
	sqlRepo := &repo.MySQLRepo{}
	sqlRepo.Init()
	gp := gpractice.GPractice{Repo: sqlRepo}
	return gp
}
