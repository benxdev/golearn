package main

import (
	"fmt"
)

// define non-altering constants 
const (
	DaysInWeek = 7
	DaysInMonth = 30
	DaysInYear = 365
	TotalHours  = 10000
)

func hours(dedicatedHours int) (int, int, int, int) {
	var totalDaysTaken, year, month, week, day int
	totalDaysTaken = (TotalHours / dedicatedHours)

	// loop through and check for necessary day, week, month and timeframe conditions
	for counter := 0; totalDaysTaken < TotalHours; {
		// calculate number of days
		if totalDaysTaken > 1 && totalDaysTaken <= DaysInWeek  {
			counter = counter + 1
			day = counter
		}
		// calculate number of weeks
		if totalDaysTaken > DaysInWeek && totalDaysTaken <= DaysInMonth {
			counter = counter + DaysInWeek
			week = counter
		}
		// calculate number of months
		if totalDaysTaken > DaysInMonth && totalDaysTaken <= DaysInYear {
			counter = counter + DaysInMonth
			month = counter
		}
		
		// calculate number of years
		if totalDaysTaken > DaysInYear {
			counter = counter + DaysInYear
			year = counter
		} 
	}
	
	return year, month, week, day
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
		fmt.Print("Not Enough Hours Dedicated!")
	}

}