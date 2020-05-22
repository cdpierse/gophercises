package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
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
	q.askQuestion()
	log.Println(q.QuestionIdx)
}
func (q *quiz) askQuestion() {
	log.Printf("Question: What is %s ?", q.Problems[q.QuestionIdx].Question)
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
