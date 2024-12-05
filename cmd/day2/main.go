package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	values []int
}

func (report *Report) addValue(value int) {
	report.values = append(report.values, value)
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (report *Report) isSafe(modified bool) bool {

	direction := report.values[1] - report.values[0]

	for i := 0; i < len(report.values)-1; i++ {
		diff := report.values[i+1] - report.values[i]

		// Check if the difference is out of allowed range
		if absInt(diff) < 1 || absInt(diff) > 3 {
			if modified {
				return false
			}
			// Try skipping this level
			return report.trySkipping(i-1) || report.trySkipping(i) || report.trySkipping(i+1)
		}

		if direction*diff < 0 {
			if modified {
				return false
			}
			// Try skipping this level
			return report.trySkipping(i-1) || report.trySkipping(i) || report.trySkipping(i+1)
		}
	}

	return true
}

// Helper method to try skipping a level
func (report *Report) trySkipping(index int) bool {
	if index < 0 {
		return false
	}
	tmpValues := []int{}
	if len(report.values[:index]) > 0 {
		tmpValues = append(tmpValues, report.values[:index]...)
	}
	if len(report.values[index+1:]) > 0 {
		tmpValues = append(tmpValues, report.values[index+1:]...)
	}
	tmpReport := Report{values: tmpValues}
	return tmpReport.isSafe(true)
}

func main() {
	var inputPath string

	flag.StringVar(&inputPath, "input-path", "", "Input file")
	flag.Parse()

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := readReports(scanner)

	safeLists := 0
	for _, report := range reports {

		if report.isSafe(false) {
			fmt.Println(report)
			safeLists++
		}
	}

	fmt.Println(safeLists)
}

func readReports(scanner *bufio.Scanner) []Report {
	reports := []Report{}

	for scanner.Scan() {
		line := scanner.Text()
		report := Report{}
		parts := strings.Split(line, " ")
		for _, val := range parts {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("Cannot convert line to integer:", line, err)
			}
			report.addValue(intVal)
		}
		reports = append(reports, report)
	}
	return reports
}
