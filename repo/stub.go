package repo

import (
	"errors"
)

const (
	NewId = 0
)

type Repository interface {
	Save(item Item) (Item, error)
	Delete(id uint64) (bool, error)
	Get(id uint64) (Item, error)
	GetAll() ([]Item, error)
}

type StubRepo struct {
	Map map[uint64]Item
}

func (r *StubRepo) Save(item Item) (Item, error) {
	if item.Id == NewId {
		item.Id = r.nextId()
	}
	r.Map[item.Id] = item
	return r.Map[item.Id], nil
}

func (r *StubRepo) Delete(id uint64) (bool, error) {
	if _, ok := r.Map[id]; ok {
		delete(r.Map, id)
		return true, nil
	}
	return false, nil
}

func (r *StubRepo) Get(id uint64) (Item, error) {
	if v, ok := r.Map[id]; ok {
		return v, nil
	}
	return Item{}, errors.New("not found")
}

func (r *StubRepo) GetAll() ([]Item, error) {
	result, i := make([]Item, len(r.Map)), 0
	for _, v := range r.Map {
		result[i] = v
		i++
	}
	return result, nil
}

func (r *StubRepo) nextId() uint64 {
	max := findMaxId(r.Map)
	max++
	return max
}

func findMaxId(items map[uint64]Item) uint64 {
	max := uint64(1)
	for k := range items {
		if k > max {
			max = k
		}
	}
	return max
}
