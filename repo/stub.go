package repo

import (
	"errors"
	"time"
)

const (
	NewId  = 0
	FORMAT = "2006-01-02"
)

type Repository interface {
	Save(item Item) (Item, error)
	Delete(id int) (bool, error)
	Get(id int) (Item, error)
	GetAll(from time.Time, to time.Time) ([]Item, error)
	Close()
}

// Stub implementation of Repository interface (backed by map[int]Item)
type StubRepo struct {
	Map map[int]Item
}

func (r *StubRepo) Save(item Item) (Item, error) {
	if item.Id == NewId {
		item.Id = int(r.nextId())
	}
	r.Map[item.Id] = item
	return r.Map[item.Id], nil
}

func (r *StubRepo) Delete(id int) (bool, error) {
	if _, ok := r.Map[id]; ok {
		delete(r.Map, id)
		return true, nil
	}
	return false, nil
}

func (r *StubRepo) Get(id int) (Item, error) {
	if v, ok := r.Map[id]; ok {
		return v, nil
	}
	return Item{}, errors.New("not found")
}

func (r *StubRepo) GetAll(from time.Time, to time.Time) ([]Item, error) {
	result, i := make([]Item, len(r.Map)), 0
	for _, v := range r.Map {
		date, e := time.Parse(FORMAT, v.Date)
		if e != nil {
			return nil, e
		}
		if date.Before(from) || date.After(to) {
			continue
		}
		result[i] = v
		i++
	}
	return result, nil
}

func (r *StubRepo) Close() {
	// no-op
}

func (r *StubRepo) nextId() int {
	max := findMaxId(r.Map)
	max++
	return max
}

func findMaxId(items map[int]Item) int {
	max := 1
	for k := range items {
		if k > max {
			max = k
		}
	}
	return max
}
