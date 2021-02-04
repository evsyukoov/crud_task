package database

import (
	"Avito/Application/io_data"
	"sort"
	"time"
)

type 	Data struct {
	Date	string		`json:"date"`
	Views   int        `json:"views"`
	Clicks  int        `json:"clicks"`
	Cost	float32      `json:"cost"`
	Cpc    float32       `json:"cpc"`
	Cpm    float32        `json:"cpm"`
}

func SortFromTypes(req io_data.Request, arr []*Data)  {
	if req.Sort_field == "" || req.Sort_field == "date" {
		sort.SliceStable(arr, func (i int, j int) bool {
			t1, _ := time.Parse("2006-01-02", arr[i].Date);
			t2, _ := time.Parse("2006-01-02", arr[i].Date);
			return t1.Before(t2);
		})
	}
	if req.Sort_field == "clicks" {
		sort.SliceStable(arr, func (i int, j int) bool {
			return (arr)[i].Clicks < (arr)[j].Clicks
		})
	}
	if req.Sort_field == "cost" {
		sort.SliceStable(arr, func (i int, j int) bool {
			return (arr)[i].Cost < (arr)[j].Cost
		})
	}
	if req.Sort_field == "views" {
		sort.SliceStable(arr, func (i int, j int) bool {
			return (arr)[i].Views < (arr)[j].Views
		})
	}
}
