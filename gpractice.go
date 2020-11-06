package gpractice

import (
	"fmt"
	"github.com/vk23/gpractice/model"
	"github.com/vk23/gpractice/repo"
	"log"
	"sort"
	"time"
)

type GPractice struct {
	Repo repo.Repository
}

func (gp *GPractice) Save(item model.Item) model.Item {
	log.Println(fmt.Sprintf("saving item %v", item))
	return gp.Repo.Save(item)
}

func (gp *GPractice) Get(id uint64) model.Item {
	log.Println(fmt.Sprintf("Getting item by id = %d", id))
	return gp.Repo.Get(id)
}

func (gp *GPractice) GetAll() []model.Item {
	log.Println("Getting all items")
	return gp.Repo.GetAll()
}

func (gp *GPractice) Delete(id uint64) bool {
	log.Println(fmt.Sprintf("Deleting item by id %v", id))
	return gp.Repo.Delete(id)
}

func (gp *GPractice) GetReport() model.Report {
	log.Println("Getting report")
	items := gp.Repo.GetAll()
	sortByDate(items)

	earliest, days, total := time.Now(), 0, uint64(0)
	prev := time.Now()
	for _, v := range items {
		d, e := time.Parse("2006-01-02", v.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v", v.Date)
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
	return model.Report{Days: days, Since: earliest.Format("2006-01-02"), Total: model.MsToReportTotal(total)}
}

func sortByDate(items []model.Item) {
	sort.Slice(items, func(i, j int) bool {
		one, two := items[i], items[j]
		d1, e := time.Parse("2006-01-02", one.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v", one.Date)
		}
		d2, e := time.Parse("2006-01-02", two.Date)
		if e != nil {
			log.Printf("Failed to parse date: %v", two.Date)
		}
		return d1.Before(d2)
	})
}
