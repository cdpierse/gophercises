package main

import (
	"github.com/cdpierse/gophercises/cyoa"
	"log"
)

func main() {

	s, err := cyoa.ReadStory("gopher.json")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(s["denver"])

}
