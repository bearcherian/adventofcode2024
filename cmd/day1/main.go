package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	listA, listB := makeListsFromInput("input")

	sortInt32(listA)
	sortInt32(listB)
	
	// fmt.Println(listA)
	// fmt.Println(listB)

	distances := getDistancesFromLists(listA, listB)

	var totalDistance int32
	for _, d := range distances {
		totalDistance += d
	}

	fmt.Println("distances: ", totalDistance)

	scores := similarityScores(listA, listB)

	var totalScores int32
	for _, d := range scores {
		totalScores += d
	}

	fmt.Println("scores: ", totalScores)
}

func similarityScores(a []int32, b []int32) []int32 {
	var scores []int32
	for _, id := range a {
		timesInB := findElementCount(b, id)

		score := id * timesInB
		scores = append(scores, score)
	}

	return scores
}

func findElementCount(list []int32, id int32) int32 {
	var count int32
	for _, id2 := range list {
		if id2 == id {
			count++
		}
	}

	return count
}

func getDistancesFromLists(a []int32, b []int32) []int32 {
	var distances []int32
	for i, id := range a {
		distance := id - b[i]
		distances = append(distances, int32(math.Abs(float64(distance))))
	}

	return distances
}

func sortInt32(list []int32) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

func makeListsFromInput(input string)([]int32, []int32) {
	var listA []int32
	var listB []int32

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

		ids := strings.Fields(line)

		if len(ids) > 0 {
			idA, _ := strconv.Atoi(ids[0])
			idB, _ := strconv.Atoi(ids[1])
			listA = append(listA, int32(idA))
			listB = append(listB, int32(idB))
		}
        
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return listA, listB
}