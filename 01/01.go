package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func ParseInput(fileContent string) ([]int, []int) {
	lines := strings.Split(fileContent, "\n")
	arr1 := make([]int, len(lines))
	arr2 := make([]int, len(lines))
	for index, line := range lines {
		lineValues := strings.Fields(line)
		arr1[index], _ = strconv.Atoi(lineValues[0])
		arr2[index], _ = strconv.Atoi(lineValues[1])
	}
	return arr1, arr2
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func puzzleOne(arr1 []int, arr2 []int) int {
	slices.Sort(arr1)
	slices.Sort(arr2)

	totalDistance := 0
	for i := 0; i < len(arr1); i++ {
		totalDistance += abs(arr1[i] - arr2[i])
	}
	return totalDistance
}

func findNumFreq(nums []int) map[int]int {
	freqMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]]++
	}
	return freqMap
}

func puzzleTwo(arr1 []int, arr2 []int) int {
	arr2Freq := findNumFreq(arr2)
	similarityScore := 0
	for _, num := range arr1 {
		similarityScore += num * arr2Freq[num]
	}

	return similarityScore
}

func main() {
	fileContent := ReadInput("01_input.txt")
	arr1, arr2 := ParseInput(fileContent)

	ans1 := puzzleOne(arr1, arr2)
	fmt.Println(ans1)
	ans2 := puzzleTwo(arr1, arr2)
	fmt.Println(ans2)
}
