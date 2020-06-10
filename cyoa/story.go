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
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`

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
		// panic(err)
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
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option is a struct that represents the possible options or directions for
// the story
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
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

	log.Println(s["intro"].Options)
	return s, nil

}
