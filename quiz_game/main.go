package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

type problem struct {
	Question string
	Answer   string
}

type quiz struct {
	Problems    []problem
	Correct     int
	Incorrect   int
	QuestionIdx int
}

func main() {
	var filename string
	flag.StringVar(&filename, "f", "problems.csv", "file to read problems for quiz from")
	flag.Parse()
	p := readProblems(filename)
	q := quiz{p, 0, 0, 0}
	for {
		if q.QuestionIdx == len(q.Problems) {
			break
		}
		q.askQuestion()
	}
	log.Printf("All done, you got %v questions right and %v wrong", q.Correct, q.Incorrect)

}
func (q *quiz) askQuestion() {
	log.Printf("Question: What is %s ?", q.Problems[q.QuestionIdx].Question)
	reader := bufio.NewReader(os.Stdin)
	log.Print("Enter your answer:")
	text, _ := reader.ReadString('\n')
	if strings.TrimSpace(text) == strings.TrimSpace(q.Problems[q.QuestionIdx].Answer) {
		q.Correct++
	} else {
		q.Incorrect++
	}
	q.QuestionIdx++

}

func readProblems(filename string) []problem {
	csvFile, err := os.Open(filename)
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	if err != nil {
		log.Panicln(err)
	}
	var problems []problem
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		problems = append(problems, problem{Question: line[0], Answer: line[1]})
	}
	return problems
}
