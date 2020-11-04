package repo

import (
	"github.com/vk23/gpractice/model"
)

type Repository interface {
	Save(item model.Item) model.Item
	Delete(item model.Item) bool
	DeleteByDate(date string) bool
	Get(date string) model.Item
	GetAll() []model.Item
}

type StubRepo struct {
	Map map[string]model.Item
}

func (r *StubRepo) Save(item model.Item) model.Item {
	r.Map[item.Date] = item
	return r.Map[item.Date]
}

func (r *StubRepo) Delete(item model.Item) bool {
	date := item.Date
	return r.DeleteByDate(date)
}

func (r *StubRepo) DeleteByDate(date string) bool {
	if _, ok := r.Map[date]; ok {
		delete(r.Map, date)
		return true
	}
	return false
}

func (r *StubRepo) Get(date string) model.Item {
	if v, ok := r.Map[date]; ok {
		return v
	}
	return model.Item{}
}

func (r *StubRepo) GetAll() []model.Item {
	result, i := make([]model.Item, len(r.Map)), 0
	for _, v := range r.Map {
		result[i] = v
		i++
	}
	return result
}
