package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type quizInfo struct {
	question string
	answer   string
}

func main() {
	count := 0
	csvFileName := flag.String("csv", "problems.csv", "csv file")
	timelimit := flag.Int("limit", 20, "Time Limit is 30 seconds")
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	flag.Parse()

	csvInput, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Println("Failed to open the CSV file : ", err)
		os.Exit(1)
	}

	read := csv.NewReader(csvInput)
	records, err := read.ReadAll()
	if err != nil {
		fmt.Println("Parse filed")
		os.Exit(1)
	}

	quiz := parseRecord(records)

loop:
	for i, p := range quiz {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		anschannel := make(chan string)
		go takeinput(anschannel)

		select {
		case <-timer.C:
			fmt.Println("\n Time is up")
			break loop

		case ans := <-anschannel:
			if p.answer == ans {
				count++
			}
		}

	}

	fmt.Printf("You scored %d out of %d \n", count, len(quiz))
}

func parseRecord(records [][]string) []quizInfo {
	tempquiz := make([]quizInfo, len(records))
	for i, eachrecord := range records {
		tempquiz[i].question = eachrecord[0]
		tempquiz[i].answer = strings.TrimSpace(eachrecord[1])
	}
	return tempquiz
}

func takeinput(anschannel chan string) {
	var str string
	fmt.Scanln(&str)
	anschannel <- str
}
