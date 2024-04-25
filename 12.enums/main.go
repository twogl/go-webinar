package main

import "fmt"

type Weekday string

const (
	Monday    Weekday = "Monday"
	Tuesday   Weekday = "Tuesday"
	Wednesday Weekday = "Wednesday"
	Thursday  Weekday = "Thursday"
	Friday    Weekday = "Friday"
	Saturday  Weekday = "Saturday"
	Sunday    Weekday = "Sunday"
)

func main() {

	fmt.Println(IsWorkDay(Wednesday))
	fmt.Println(IsWorkDay(Sunday))

	fmt.Println(IsWorkDay("Monday"))
	fmt.Println(IsWorkDay("Saturday"))
}

func IsWorkDay(day Weekday) bool {
	if day != Saturday && day != Sunday {
		return false
	}
	return true
}

func IsWorkDayStr(day string) bool {
	weekDay := Weekday(day)
	if weekDay != Saturday && weekDay != Sunday {
		return false
	}
	return true
}
