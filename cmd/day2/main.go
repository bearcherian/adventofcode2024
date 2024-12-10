package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reports := getReports("input")

	safeReports := getSafeReports(reports)
	fmt.Println("safe reports: ", len(safeReports))
}

func getSafeReports(reports [][]int) [][]int {

	var safeReports [][]int

	for _, report := range reports {

		if isReportSafe(report, false) {
			safeReports = append(safeReports, report)
		}
	}

	return safeReports

}

func isReportSafe(report []int, isAdjusted bool) bool {

	var isIncreasing bool
	var isDecreasing bool
	for i, level := range report {
		// if this is the first one, skip it
		if i == 0 {
			continue
		} 

		prevLevel := report[i-1]
		
		// if the next one is less, we are decreasing
		if level < prevLevel {
			isDecreasing = true
		} 
		if level > prevLevel {
			isIncreasing = true
		} 

		diff := level - prevLevel
		absDiff := math.Abs(float64(diff)) 

		var isInRange bool
		if absDiff >= 1 && absDiff <=3 {
			isInRange = true
		}

		if (isIncreasing && isDecreasing) || !isInRange {
			// we have a bad level, can we remove it here? (part 2)
			if !isAdjusted {
				// If this is item 1, check if removing item 0 makes it sage
				if i == 1 {
					adjustedReportA := report[1:] // could use 'i', but let's be explicit
					if isReportSafe(adjustedReportA, true) { return true }
				}
				adjustedReportB := append(report[:i], report[i+1:]...)
				return isReportSafe(adjustedReportB, true)// || isReportSafe(adjustedReportA, true)  // if the new report is bad, we don't 
			}
			return false
		}	

	}

	return true
}

func getReports(input string) [][]int {
	var reports [][]int

	// open file
    f, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        // do something with a line
		line := scanner.Text()

		levels := strings.Fields(line)

		if len(levels) > 0 {
			
			var report []int
			for _, level := range levels {
				l, _ := strconv.Atoi(level)
				report = append(report, l)
			}

			reports = append(reports, report)
		}
        
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return reports
}