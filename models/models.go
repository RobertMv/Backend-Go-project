package models

import "time"

type Restaurant struct {
	Id          int64
	Name        string
	Address     string
	YearProfit  float32
	MonthProfit float32
}

type Employee struct {
	Id             int64
	Name           string
	Surname        string
	Patronymic     string
	BirthDate      time.Time
	Sex            rune
	Passport       string
	EmploymentDate time.Time
	Phone          string
	Position
	Restaurant
}

type Position struct {
	Id     int64
	Code   uint8
	Name   string
	Salary float32
}
