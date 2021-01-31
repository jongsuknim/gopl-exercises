package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(w io.Writer) {
	var trackList = template.Must(template.New("tracklist").Parse(`
<table>
<tr style='text-align: left'>
  <th><a href='/sort?key=Title'>Title</a></th>
  <th><a href='/sort?key=Artist'>Artist</a></th>
  <th><a href='/sort?key=Album'>Album</a></th>
  <th><a href='/sort?key=Year'>Year</a></th>
  <th><a href='/sort?key=Length'>Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))
	if err := trackList.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var tracks = []*Track{
	{"Go", "Deliash", "From the roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		printTracks(w)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/sort", sortHandler)

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title == y.Title {
			return x.Year < y.Year
		}
		return x.Title < y.Title
	}})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	key, value := r.Form["key"]
	if !value {
		return
	}

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		switch key[0] {
		case "Title":
			return x.Title < y.Title
		case "Artist":
			return x.Artist < y.Artist
		case "Album":
			return x.Album < y.Album
		case "Year":
			return x.Year < y.Year
		case "Length":
			return x.Length < y.Length
		}
		return true
	}})

	printTracks(w)
}
