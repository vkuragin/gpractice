package gpractice

import (
	"github.com/vk23/gpractice/repo"
	"reflect"
	"testing"
)

const (
	DATE     = "2020-10-31"
	DATE2    = "1999-01-01"
	DURATION = 15000
)

var gPractice GPractice
var items map[uint64]repo.Item

func setupTestCase(t *testing.T) func(t *testing.T) {
	items = map[uint64]repo.Item{
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
		id uint64
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

func TestGetAll(t *testing.T) {
	var testItem = repo.Item{Id: 1, Date: DATE, Duration: DURATION}

	tests := []struct {
		name string
		want []repo.Item
	}{
		{"get all", []repo.Item{testItem}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			t.Logf("After testSetup: %v", items)

			if got, _ := gPractice.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		id uint64
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
	tests := []struct {
		name string
		want repo.Report
	}{
		{"report", repo.Report{Days: 1, Since: DATE, Total: repo.ReportTotal{Days: 0, Hours: 0, Minutes: 0, Seconds: 15}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got, _ := gPractice.GetReport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
