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

		if isReportSafe(report) {
			safeReports = append(safeReports, report)
		}
	}

	return safeReports

}

func isReportSafe(report []int) bool {

	var isIncreasing bool
	var isDecreasing bool
	for i, level := range report {
		if i == 0 {
			continue
		} 

		
		if level < report[i-1] {
			isDecreasing = true
		}
		if level > report[i-1] {
			isIncreasing = true
		}

		if isIncreasing && isDecreasing {
			return false
		}	

		diff := level - report[i-1]

		absDiff := math.Abs(float64(diff)) 
		if absDiff > 0 && absDiff < 4 {
			continue
		} else {
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