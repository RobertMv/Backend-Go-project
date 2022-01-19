package main

import "time"

type restaurant struct {
	name        string
	address     string
	yearProfit  float32
	monthProfit float32
}

type employee struct {
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
	code   uint8
	name   string
	salary float32
}
