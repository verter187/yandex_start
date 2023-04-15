package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func NextDay(day Weekday) Weekday {
	fmt.Println("day: ", day, "day % 7: ", day%7)
	return (day % 7) + 1
}

func main() {
	var today Weekday = Saturday
	tomorrow := NextDay(today)
	fmt.Println("today =", today, "tomorrow =", tomorrow)
}
