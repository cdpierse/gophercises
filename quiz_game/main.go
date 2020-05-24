package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"time"
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
	var timeLimit int
	flag.StringVar(&filename, "f", "problems.csv", "file to read problems for quiz from")
	flag.IntVar(&timeLimit, "timelimit", 30, "time limit in seconds to solve problems within")
	flag.Parse()
	p := readProblems(filename)
	q := quiz{p, 0, 0, 0}
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	for {
		log.Printf("Question: What is %s ?", q.Problems[q.QuestionIdx].Question)
		answerCh := make(chan string)
		go q.askQuestion(answerCh)
		select {
		case <-timer.C:
			log.Printf("Time Up, you scored %v questions right and %v wrong out of %v questions",
				q.Correct,
				q.Incorrect,
				len(q.Problems))
			return
		case _ = <-answerCh:

		}

	}

}
func (q *quiz) askQuestion(c chan string) {
	reader := bufio.NewReader(os.Stdin)
	log.Print("Enter your answer:")
	text, _ := reader.ReadString('\n')
	if strings.TrimSpace(text) == strings.TrimSpace(q.Problems[q.QuestionIdx].Answer) {
		q.Correct++
	} else {
		q.Incorrect++
	}
	q.QuestionIdx++
	c <- "received" + text

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
