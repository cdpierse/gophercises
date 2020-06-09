package cyoa

import (
	"encoding/json"
	_ "encoding/json"
	"html/template"
	_ "io"
	_ "io/ioutil"
	"log"
	"net/http"
	"os"
)

var defaultHandlerTempl = `
	<!DOCTYPE html>
	<html>

	<head>
		<meta charset="utf-8">
		<title> Choose Your Own Adventure</title>
	</head>

	<body>
		<h1>{{.Title}} </h1>
		{{range .StoryParagraphs}}
		<p>{{.}}</p>
		{{end}}
		<ul>
			{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
			{{end}}
		</ul>

	</body>

	</html>
	`

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTempl))
}

var tpl *template.Template

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}

}

// Story represents a standard choose your own
// adventure story. A story consists of chapter keys (strings)
// which map a chapter name to an instance of a Chapter struct.
type Story map[string]Chapter

// Chapter represents a possible chatper or route in an adventure.
// A chapter consits of a title and array of the story points. Finally
// a chapter has an array of Option structs which the user interacts with
// to progress the story.
type Chapter struct {
	Title           string   `json:"title"`
	StoryParagraphs []string `json:"story"`
	Options         []Option `json:"options"`
}

// Option is a struct that represents the possible options or directions for
// the story
type Option struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

func ReadStory(name string) (Story, error) {
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
