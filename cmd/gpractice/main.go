package main

import (
	"flag"
	"fmt"
	"github.com/vkuragin/gpractice"
	"github.com/vkuragin/gpractice/repo"
	"log"
	"os"
	"time"
)

type Action string

const (
	ADD    Action = "add"
	GET           = "get"
	DEL           = "del"
	REPORT        = "report"
	IMPORT        = "import"
	EXPORT        = "export"
)

func main() {
	actionFlag := flag.String("action", string(REPORT), "one of possible actions: add, del, get, report, import, export")
	dateFlag := flag.String("date", time.Now().Format(repo.DateFormat), "practice date yyyy-MM-dd")
	minutesFlag := flag.Int("minutes", 0, "practice time in minutes")
	idFlag := flag.Int("id", 0, "id")
	fileFlag := flag.String("f", "data.csv", "file")
	flag.Parse()

	execute(*actionFlag, *idFlag, *dateFlag, *minutesFlag, *fileFlag)
}

func execute(action string, id int, date string, minutes int, file string) {
	log.Println(fmt.Sprintf("Executing action [%s] with values: [%v, %v, %v]", action, id, date, minutes))

	// config
	cfg, err := gpractice.LoadCfg("~/.gpractice/config.yaml")
	if err != nil {
		log.Fatalf("Cannot load config: %s", err)
		os.Exit(1)
	}

	// initialize db
	sqlRepo := &repo.MySQLRepo{DbUser: cfg.Db.UserName, DbPass: cfg.Db.UserPass, DbHost: cfg.Db.Host, DbPort: cfg.Db.Port, DbName: cfg.Db.Name}
	err = sqlRepo.Init()
	if err != nil {
		log.Fatalf("Cannot initialize db: %s", err)
		os.Exit(1)
	}
	defer sqlRepo.Close()

	gp := gpractice.GPractice{Repo: sqlRepo}

	item := repo.Item{Id: int(id), Date: date, Duration: minutes * 60}
	switch Action(action) {
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
		report, err := gp.GetReport(time.Unix(0, 0), time.Now())
		log.Printf("result: %v\n, error: %v\n", report, err)
	case IMPORT:
		err := gp.Import(file)
		log.Printf("result: done\n, error: %v\n", err)
	case EXPORT:
		err := gp.Export(file)
		log.Printf("result: done\n, error: %v\n", err)
	default:
		log.Fatalf("Unknown action: %v\n", action)
		os.Exit(1)
	}
}
