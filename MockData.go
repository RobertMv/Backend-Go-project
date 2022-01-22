package main

import "time"

var employees = []employee{mockEmployeeOne, mockEmployeeTwo}
var positions = []position{director, intern}
var restaurants = []restaurant{restOne, restTwo}

var director = position{1, 111, "Director", 200000.00}
var intern = position{2, 023, "Intern", 30000.00}

var restOne = restaurant{3, "Lancelot", "Kazan, Iskra, 7", 80009990.00, 7800000.00}
var restTwo = restaurant{4, "Pobeda", "Kazan, Pobeda Avenue, 139", 90009990.00, 9800000.00}

var mockEmployeeOne = employee{
	5, "Robert", "Mulyukov", "Rustamovich",
	time.Date(2000, 10, 18,
		0, 0, 0, 0, time.UTC),
	'M', "9214756693", time.Now(),
	"+79196952500", intern, restOne}

var mockEmployeeTwo = employee{
	6, "Aleksey", "Chepakhin", "Valeryievich",
	time.Date(2000, 02, 07,
		0, 0, 0, 0, time.UTC),
	'M', "9214756693", time.Now(),
	"+79196952500", director, restTwo}
