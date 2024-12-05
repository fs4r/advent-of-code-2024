package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	var inputPath string
	flag.StringVar(&inputPath, "input-path", "input.txt", "path to input file")
	flag.Parse()

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list1, list2 := parseFile(scanner)

	if len(list1) != len(list2) {
		log.Fatal("Lists are not the same length")
	}

	sort.Ints(list1)
	sort.Ints(list2)
	sum := 0
	similarity := 0
	for i := 0; i < len(list1); i++ {
		sum += absInt(list1[i] - list2[i])
		occurrences := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				occurrences += 1
			} else if list1[i] < list2[j] {
				break
			}
		}
		similarity += occurrences * list1[i]
	}

	log.Println("Result is: ", sum)
	log.Println("Similarity is: ", similarity)
}

func parseFile(scanner *bufio.Scanner) ([]int, []int) {
	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			log.Fatalf("Invalid input line (expected two parts): %q", line)
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to parse x as integer: %q in line %q", parts[0], line)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Failed to parse y as integer: %q in line %q", parts[1], line)
		}

		list1 = append(list1, x)
		list2 = append(list2, y)
	}

	return list1, list2
}
