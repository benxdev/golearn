package main

import (
	"fmt"
)

// define non-altering constants
const (
	DaysInWeek  = 7
	DaysInMonth = 30
	DaysInYear  = 365
	TotalHours  = 10000
)

func hours(dedicatedHours int) (int, int, int, int) {
	var totalDaysTaken, years, months, weeks, days, daysRemaining int

	if dedicatedHours <= 0 {
		return 0, 0, 0, 0
	} else {
		// calculate number of days if input dedicated hours > 0
		totalDaysTaken = (TotalHours / dedicatedHours)

		// calculate number of years
		years = totalDaysTaken / DaysInYear
		daysRemaining = totalDaysTaken % DaysInYear

		// calculate number of months
		months = daysRemaining / DaysInMonth
		daysRemaining = daysRemaining % DaysInMonth

		// calculate number of weeks
		weeks = daysRemaining / DaysInWeek
		daysRemaining = daysRemaining % DaysInWeek

		// calculate number of days
		days = daysRemaining
	}

	return years, months, weeks, days
}

func main() {
	var dedicatedHours int
	fmt.Println("How many hours are you willing to dedicate daily?")
	fmt.Scanf("%d", &dedicatedHours)

	// call hours function to calculate for total number of days if hours dedicated is up to 1
	if dedicatedHours >= 1 {
		year, month, week, day := hours(dedicatedHours)
		fmt.Print("The total number of days, months or years it will take is exactly: \n")
		fmt.Printf("%v years, %v months, %v weeks and %v days\n", year, month, week, day)
	} else {
		fmt.Print("Not Enough Hours Dedicated!\n")
	}

}
