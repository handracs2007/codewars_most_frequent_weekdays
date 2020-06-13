// https://www.codewars.com/kata/56eb16655250549e4b0013f4/train/go

package main

import (
	"fmt"
	"time"
)

func MostFrequentDays(year int) []string {
	dayMap := make(map[time.Weekday]int)

	days := make([]time.Weekday, 7)
	days[0] = time.Monday
	days[1] = time.Tuesday
	days[2] = time.Wednesday
	days[3] = time.Thursday
	days[4] = time.Friday
	days[5] = time.Saturday
	days[6] = time.Sunday

	d := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	for d.Year() == year {
		day := d.Weekday()
		dayMap[day] += 1

		d = d.Add(time.Hour * 24)
	}

	// Get the maximum number of days
	maxCount := 0
	for _, value :=  range dayMap {
		if value > maxCount {
			maxCount = value
		}
	}

	// List the days with maximum count
	result := make([]string, 0)
	for i := 0; i < 7; i ++ {
		value := dayMap[days[i]]

		if value == maxCount {
			result = append(result, days[i].String())
		}
	}

	return result
}

func main() {
	fmt.Println(MostFrequentDays(1984))
	fmt.Println(MostFrequentDays(2427))
	fmt.Println(MostFrequentDays(2185))
	fmt.Println(MostFrequentDays(2860))
	fmt.Println(MostFrequentDays(3076))
}
