package gpractice

import (
	"encoding/csv"
	"fmt"
	"github.com/vkuragin/gpractice/repo"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// Main service
type GPractice struct {
	// Repository, implementations: StubRepo, MySQLRepo
	Repo repo.Repository
}

func (gp *GPractice) Save(item repo.Item) (repo.Item, error) {
	log.Printf("Saving item: %v\n", item)
	saved, err := gp.Repo.Save(item)
	log.Printf("Saving item result: %v\n", item)
	return saved, err
}

func (gp *GPractice) Get(id int) (repo.Item, error) {
	log.Printf("Getting item by id: %d\n", id)
	item, err := gp.Repo.Get(id)
	log.Printf("Getting item by id: %d, result: %v\n", id, item)
	return item, err
}

func (gp *GPractice) Delete(id int) (bool, error) {
	log.Println(fmt.Sprintf("Deleting item by id %v\n", id))
	return gp.Repo.Delete(id)
}

func (gp *GPractice) GetReport(from *time.Time, to *time.Time) (repo.Report, error) {
	// date range (defaults to current year)
	now := time.Now()
	if from == nil {
		t := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
		from = &t
	}
	if to == nil {
		to = &now
	}

	log.Printf("Getting report: %v - %v\n", from, to)
	items, err := gp.Repo.GetAll(*from, *to)
	if err != nil {
		return repo.Report{}, err
	}

	sortByDate(items, false)

	earliest, days, total := time.Now(), 0, 0
	prev := time.Now()
	for _, v := range items {
		d, e := time.Parse(repo.DateFormat, v.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v\n", v.Date)
			continue
		}
		if d.Before(earliest) {
			earliest = d
		}
		if prev != d {
			prev = d
			days++
		}
		total += v.Duration
	}

	sortByDate(items, true)

	report := repo.Report{
		Items:     items,
		Days:      days,
		DateStart: from.Format(repo.FORMAT),
		DateEnd:   to.Format(repo.FORMAT),
		Total:     total,
	}
	//log.Printf("Getting report result: %v\n", report)
	return report, nil
}

func (gp *GPractice) Import(filePath string) error {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer closeFile(file)

	// read all records from file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// save all records (assuming that the first row is a header)
	for i := 1; i < len(records); i++ {
		rec := records[i]
		item, err := fromCsvRecord(rec)
		if err != nil {
			log.Printf("Failed to process record: %v. Error: %v\n", rec, err)
			continue
		}
		_, err = gp.Repo.Save(item)
		if err != nil {
			log.Printf("Failed to save record: %v. Error: %v\n", item, err)
			continue
		}
	}
	log.Printf("records: %v\n", records)
	return nil
}

func (gp *GPractice) Export(filePath string) error {
	// open or create file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer closeFile(file)

	// get all records
	items, err := gp.Repo.GetAll(time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local), time.Now())
	if err != nil {
		return err
	}

	// convert and write to file
	records := make([][]string, len(items)+1)
	records[0] = []string{"id", "date", "duration(seconds)"}
	for i, v := range items {
		records[i+1] = toCsvRecord(v)
	}
	w := csv.NewWriter(file)
	err = w.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("Failed to close file: %v. Error: %v\n", file, err)
	}
}

func sortByDate(items []repo.Item, reverse bool) {
	sort.Slice(items, func(i, j int) bool {
		one, two := items[i], items[j]
		d1, e := time.Parse(repo.DateFormat, one.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v\n", one.Date)
		}
		d2, e := time.Parse(repo.DateFormat, two.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v\n", two.Date)
		}
		if reverse {
			return d2.Before(d1)
		}
		return d1.Before(d2)
	})
}

func toCsvRecord(i repo.Item) []string {
	return []string{fmt.Sprintf("%d", i.Id), i.Date, fmt.Sprintf("%d", i.Duration)}
}

func fromCsvRecord(r []string) (repo.Item, error) {
	if len(r) != 3 {
		return repo.Item{}, fmt.Errorf("expected 3 fields, got: %d", len(r))
	}

	// OK, new record
	id := 0

	// validate date
	date := r[1]
	var validDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !validDate.MatchString(date) {
		return repo.Item{}, fmt.Errorf("invalid date %s", date)
	}

	duration, err := strconv.Atoi(r[2])
	if err != nil {
		return repo.Item{}, fmt.Errorf("invalid duration %s", r[2])
	}

	return repo.Item{id, date, duration}, nil
}
