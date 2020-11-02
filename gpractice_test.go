package gpractice

import (
	"github.com/vk23/gpractice/model"
	"github.com/vk23/gpractice/repo"
	"reflect"
	"testing"
)

var gPractice GPractice

func setupTestCase(t *testing.T) func(t *testing.T) {
	items := make(map[string]model.Item)
	gPractice = GPractice{&repo.StubRepo{items}}
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		item model.Item
	}

	var testItem = model.Item{"2020-10-31", 15}

	tests := []struct {
		name string
		args args
		want model.Item
	}{
		{"new", args{testItem}, testItem},
		{"update", args{testItem}, model.Item{"2020-10-31", 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	tests := []struct {
		name string
		args args
		want model.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gPractice.Get(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name string
		want []model.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gPractice.DeleteByDate(tt.args.date); got != tt.want {
				t.Errorf("DeleteByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReport(t *testing.T) {
	tests := []struct {
		name string
		want Report
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gPractice.GetReport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
