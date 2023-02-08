package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFileName := flag.String("CSV", "Problem.CSV", "A CSV file in the form of 'question,answer'")
	flag.Parse()
	_ = csvFileName

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file : %s \n", *csvFileName))

	} else {
		fmt.Print(" Open sucessfully")
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(" Failed to parse the provided CSV file.")
	}
	fmt.Println(lines)
	Problems := parseLines(lines)

	timeLimit := 2

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	correct := 0
	for i, p := range Problems {
		fmt.Printf("Problem #%d: %s = \n  ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n ", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d", correct, len(Problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}

		}

	}

	fmt.Printf("You scored %d out of %d", correct, len(Problems))
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type Problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
