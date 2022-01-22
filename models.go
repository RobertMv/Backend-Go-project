package main

import "time"

type restaurant struct {
	id          int64
	name        string
	address     string
	yearProfit  float32
	monthProfit float32
}

type employee struct {
	id             int64
	name           string
	surname        string
	patronymic     string
	birthDate      time.Time
	sex            rune
	passport       string
	employmentDate time.Time
	phone          string
	position
	restaurant
}

type position struct {
	id     int64
	code   uint8
	name   string
	salary float32
}
