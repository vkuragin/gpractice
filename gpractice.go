package gpractice

import (
	"fmt"
	"github.com/vk23/gpractice/model"
	"github.com/vk23/gpractice/repo"
	"log"
)

type GPractice struct {
	Repo repo.Repository
}

func (gp *GPractice) Add(item model.Item) model.Item {
	log.Println(fmt.Sprintf("adding item %v", item))
	stored := gp.Repo.Get(item.Date)
	item.Duration += stored.Duration
	return gp.Repo.Save(item)
}

func (gp *GPractice) Get(date string) model.Item {
	log.Println(fmt.Sprintf("Getting item by %s", date))
	return gp.Repo.Get(date)
}

func (gp *GPractice) GetAll() []model.Item {
	log.Println("Getting all items")
	return gp.Repo.GetAll()
}

func (gp *GPractice) Delete(item model.Item) bool {
	log.Println(fmt.Sprintf("Deleting item %v", item))
	return gp.Repo.Delete(item)
}

func (gp *GPractice) DeleteByDate(date string) bool {
	log.Println(fmt.Sprintf("Deleting item by %s", date))
	return gp.Repo.DeleteByDate(date)
}

func (gp *GPractice) GetReport() model.Report {
	log.Println("Getting report")
	_ = gp.Repo.GetAll()
	return model.Report{}
}
