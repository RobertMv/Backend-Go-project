package models

import "time"

type Restaurant struct {
	id          int64
	name        string
	address     string
	yearProfit  float32
	monthProfit float32
}

type Employee struct {
	id             int64
	name           string
	surname        string
	patronymic     string
	birthDate      time.Time
	sex            rune
	passport       string
	employmentDate time.Time
	phone          string
	Position
	Restaurant
}

type Position struct {
	id     int64
	code   uint8
	name   string
	salary float32
}
