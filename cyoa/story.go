package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("hello")
	s, err := readStory("gopher.json")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(s)
}

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

func readStory(name string) (Story, error) {
	dat, err := os.Open(name)
	// dat, err := ioutil.ReadFile(name)
	if err != nil {
		log.Panicln(err)
	}
	var s Story
	d := json.NewDecoder(dat)
	if err := d.Decode(&s); err != nil {
		return nil, err
	}
	return s, nil

}
