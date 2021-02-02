package io_data


//входные данные на POST запрос
type Json struct {
	IsContainsDate bool;
	Date string;
	Views int;
	Clicks int;
	Cost  float32;
}

func NewJson() *Json {
	return &Json{ IsContainsDate: false,
		Clicks: 0,
		Views: 0,
		Cost: 0}
}

// Входные данные на GET
type 	Request struct {
	From string;
	To 	string;
	Sort_field string;
}
