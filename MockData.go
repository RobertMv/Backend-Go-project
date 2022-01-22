package main

import "time"

var employees = []employee{mockEmployeeOne, mockEmployeeTwo}
var positions = []position{director, intern}
var restaurants = []restaurant{restOne, restTwo}

var director = position{111, "Director", 200000.00}
var intern = position{023, "Intern", 30000.00}

var restOne = restaurant{"Lancelot", "Kazan, Iskra, 7", 80009990.00, 7800000.00}
var restTwo = restaurant{"Pobeda", "Kazan, Pobeda Avenue, 139", 90009990.00, 9800000.00}

var mockEmployeeOne = employee{
	"Robert", "Mulyukov", "Rustamovich",
	time.Date(2000, 10, 18,
		0, 0, 0, 0, time.UTC),
	'M', "9214756693", time.Now(),
	"+79196952500", intern, restOne}

var mockEmployeeTwo = employee{
	"Aleksey", "Chepakhin", "Valeryievich",
	time.Date(2000, 02, 07,
		0, 0, 0, 0, time.UTC),
	'M', "9214756693", time.Now(),
	"+79196952500", director, restTwo}
