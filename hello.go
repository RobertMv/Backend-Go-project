package main

import (
	"fmt"
	"time"
)

func main() {

	e := employee{
		name:           "Robert",
		surname:        "Mulyukov",
		patronymic:     "Rustamovich",
		birthDate:      time.Now(),
		sex:            'M',
		passport:       "9227756693",
		employmentDate: time.Now(),
		phone:          "89196952500",
		position: position{
			code:   123,
			name:   "Director",
			salary: 110000.00,
		},
		restaurant: restaurant{
			name:        "Central Park Cafe",
			address:     "City, street, number",
			yearProfit:  89496450.00,
			monthProfit: 7845945.00,
		},
	}
	fmt.Println(e)
}
