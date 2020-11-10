package gpractice

import (
	"fmt"
	"github.com/vk23/gpractice/repo"
	"log"
	"sort"
	"time"
)

type GPractice struct {
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

func (gp *GPractice) GetAll() ([]repo.Item, error) {
	log.Printf("Getting all items\n")
	items, err := gp.Repo.GetAll()
	log.Printf("Getting all items result: %v\n", items)
	return items, err
}

func (gp *GPractice) Delete(id int) (bool, error) {
	log.Println(fmt.Sprintf("Deleting item by id %v\n", id))
	return gp.Repo.Delete(id)
}

func (gp *GPractice) GetReport() (repo.Report, error) {
	log.Printf("Getting report\n")
	items, err := gp.Repo.GetAll()
	if err != nil {
		return repo.Report{}, err
	}

	sortByDate(items)

	earliest, days, total := time.Now(), 0, 0
	prev := time.Now()
	for _, v := range items {
		d, e := time.Parse("2006-01-02", v.Date)
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

	report := repo.Report{Days: days, Since: earliest.Format("2006-01-02"), Total: repo.SecondsToReportTotal(total)}
	log.Printf("Getting report result: %v\n", report)
	return report, nil
}

func sortByDate(items []repo.Item) {
	sort.Slice(items, func(i, j int) bool {
		one, two := items[i], items[j]
		d1, e := time.Parse("2006-01-02", one.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v\n", one.Date)
		}
		d2, e := time.Parse("2006-01-02", two.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v\n", two.Date)
		}
		return d1.Before(d2)
	})
}
