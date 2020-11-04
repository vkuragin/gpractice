package repo

import (
	"github.com/vk23/gpractice/model"
)

type Repository interface {
	Save(item model.Item) model.Item
	Delete(item model.Item) bool
	DeleteById(id uint64) bool
	Get(id uint64) model.Item
	GetAll() []model.Item
}

type StubRepo struct {
	Map map[uint64]model.Item
}

func (r *StubRepo) Save(item model.Item) model.Item {
	r.Map[item.Id] = item
	return r.Map[item.Id]
}

func (r *StubRepo) Delete(item model.Item) bool {
	return r.DeleteById(item.Id)
}

func (r *StubRepo) DeleteById(id uint64) bool {
	if _, ok := r.Map[id]; ok {
		delete(r.Map, id)
		return true
	}
	return false
}

func (r *StubRepo) Get(id uint64) model.Item {
	if v, ok := r.Map[id]; ok {
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
