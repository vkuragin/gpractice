package gpractice

import (
	"github.com/vkuragin/gpractice/repo"
	"reflect"
	"testing"
	"time"
)

const (
	DATE     = "2020-10-31"
	DATE2    = "1999-01-01"
	DURATION = 15
	FORMAT   = "2006-01-02"
)

var gPractice GPractice
var items map[int]repo.Item

func setupTestCase(t *testing.T) func(t *testing.T) {
	items = map[int]repo.Item{
		1: {Id: 1, Date: DATE, Duration: DURATION},
	}
	gPractice = GPractice{&repo.StubRepo{Map: items}}
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func TestSave(t *testing.T) {
	type args struct {
		item repo.Item
	}

	var notExisting = repo.Item{Date: DATE2, Duration: DURATION}
	var existing = repo.Item{Id: 1, Date: DATE, Duration: DURATION}

	tests := []struct {
		name string
		args args
		want repo.Item
	}{
		{"update", args{existing}, existing},
		{"new", args{notExisting}, repo.Item{Id: 2, Date: DATE2, Duration: DURATION}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got, _ := gPractice.Save(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		id int
	}

	var testItem = repo.Item{Id: 1, Date: DATE, Duration: DURATION}

	tests := []struct {
		name string
		args args
		want repo.Item
	}{
		{"get", args{1}, testItem},
		//{"get", args{DATE2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got, _ := gPractice.Get(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		id int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"delete", args{1}, true},
		{"delete", args{0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got, _ := gPractice.Delete(tt.args.id); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReport(t *testing.T) {
	var testItem = repo.Item{Id: 1, Date: DATE, Duration: DURATION}

	tests := []struct {
		name string
		want repo.Report
	}{
		{"report", repo.Report{Items: []repo.Item{testItem}, Days: 1, DateStart: DATE, DateEnd: DATE, Total: 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			date, e := time.Parse(FORMAT, DATE)
			if e != nil {
				t.Errorf("Error: %v", e)
				return
			}
			if got, _ := gPractice.GetReport(date, date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
