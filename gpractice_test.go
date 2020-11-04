package gpractice

import (
	"github.com/vk23/gpractice/model"
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
var items map[string]model.Item

func setupTestCase(t *testing.T) func(t *testing.T) {
	items = map[string]model.Item{
		DATE: {Date: DATE, Duration: DURATION},
	}
	gPractice = GPractice{&repo.StubRepo{Map: items}}
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		item model.Item
	}

	var existing = model.Item{Date: DATE, Duration: DURATION}
	var notExisting = model.Item{Date: DATE2, Duration: DURATION}

	tests := []struct {
		name string
		args args
		want model.Item
	}{
		{"new", args{notExisting}, notExisting},
		{"update", args{existing}, model.Item{DATE, DURATION * 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got := gPractice.Add(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		date string
	}

	var testItem = model.Item{Date: DATE, Duration: DURATION}

	tests := []struct {
		name string
		args args
		want model.Item
	}{
		{"get", args{DATE}, testItem},
		//{"get", args{DATE2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got := gPractice.Get(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	var testItem = model.Item{Date: DATE, Duration: DURATION}

	tests := []struct {
		name string
		want []model.Item
	}{
		{"get all", []model.Item{testItem}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			t.Logf("After testSetup: %v", items)

			if got := gPractice.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		item model.Item
	}

	var testItem = model.Item{Date: DATE, Duration: DURATION}
	var nonExisting = model.Item{Date: DATE2}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"delete", args{testItem}, true},
		{"delete", args{nonExisting}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got := gPractice.Delete(tt.args.item); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteByDate(t *testing.T) {
	type args struct {
		date string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"delete", args{DATE}, true},
		{"delete", args{DATE2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got := gPractice.DeleteByDate(tt.args.date); got != tt.want {
				t.Errorf("DeleteByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReport(t *testing.T) {
	tests := []struct {
		name string
		want model.Report
	}{
		{"report", model.Report{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTestCase(t)
			if got := gPractice.GetReport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
