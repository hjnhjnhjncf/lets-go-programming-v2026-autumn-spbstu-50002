package main

import "fmt"

const (
	minimalTemperature = 15
	maximalTemperature = 30
)

func computeTemperature(measurement string, minimalLocal, maximalLocal *int) int {
	switch measurement {
	case "<=":
		var temperature int
		_, err := fmt.Scan(&temperature)

		if err != nil {
			fmt.Println("Invalid temperature")
		} else if temperature < *maximalLocal {
			*maximalLocal = temperature
		}
	case ">=":
		var temperature int
		_, err := fmt.Scan(&temperature)

		if err != nil {
			fmt.Println("Invalid temperature")
		} else if temperature > *minimalLocal {
			*minimalLocal = temperature
		}
	}

	if *maximalLocal-*minimalLocal >= 0 {
		return *minimalLocal
	}

	return -1
}

func main() {
	var depNumber int
	_, err := fmt.Scan(&depNumber)

	if err != nil {
		fmt.Println("Invalid depNumber")

		return
	}

	for range depNumber {
		var workerNumber int
		_, err := fmt.Scan(&workerNumber)

		if err != nil {
			fmt.Println("Invalid workerNumber")

			return
		}

		minimalLocal := minimalTemperature
		maximalLocal := maximalTemperature

		for range workerNumber {
			var measurement string
			_, err := fmt.Scan(&measurement)

			if err != nil || measurement != "<=" && measurement != ">=" {
				fmt.Println("Invalid measure")

				return
			}

			temp := computeTemperature(measurement, &minimalLocal, &maximalLocal)
			fmt.Println(temp)
		}
	}
}
