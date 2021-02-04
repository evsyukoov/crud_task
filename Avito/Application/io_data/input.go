package io_data

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

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

// Входные данные на БД

type Config struct {
	Host   string
	User   string
	Pass   string
	Dbname string
}

func loadConfig() (error, *Config) {
	file,err := os.Open("./config/config.json")
	if err != nil {
		log.Println("No such config file")
		return errors.New("config error"), nil
	}
	decoder := json.NewDecoder(file)
	config := new(Config);
	err = decoder.Decode(&config);
	if err != nil {
		log.Println("Bad structure of JSON config")
		return errors.New("config error"), nil;
	}
	return nil, config;
}

func InitDbParametrs() (error, *Config) {
	conf, err := loadConfig();
	if err != nil {
		return nil, err
	} else {
		return conf, err
	}
}
