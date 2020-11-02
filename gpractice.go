package gpractice

import (
	"fmt"
	"github.com/vk23/gpractice/model"
	"github.com/vk23/gpractice/repo"
	"log"
)

type ReportTotal struct {
	days    int
	hours   int
	minutes int
}

type Report struct {
	days  int
	since string
	total ReportTotal
}

type GPractice struct {
	Repo repo.Repository
}

func (gp *GPractice) Add(item model.Item) model.Item {
	log.Println(fmt.Sprintf("adding item %v", item))

	return item
}

func (gp *GPractice) Get(date string) model.Item {
	log.Println(fmt.Sprintf("Getting item by %s", date))

	return model.Item{}
}

func (gp *GPractice) GetAll() []model.Item {
	log.Println("Getting all items")

	return nil
}

func (gp *GPractice) Delete(item model.Item) bool {
	log.Println(fmt.Sprintf("Deleting item %v", item))

	return false
}

func (gp *GPractice) DeleteByDate(date string) bool {
	log.Println(fmt.Sprintf("Deleting item by %s", date))

	return false
}

func (gp *GPractice) GetReport() Report {
	log.Println("Getting report")

	return Report{}
}
