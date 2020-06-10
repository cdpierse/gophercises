package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cdpierse/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "Port at which to serve CYOA")
	filename := flag.String("filename", "gopher.json", "Story file from which CYOA game is basedƒ")
	flag.Parse()
	story, err := cyoa.ReadJSONStory(*filename)
	if err != nil {
		log.Panicln(err)
	}

	h := cyoa.NewHandler(story)
	log.Printf("Starting the server at port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
