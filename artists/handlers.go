package artists

import (
	"html/template"
	"io"
	"net/http"
	"strconv"
)

func render(w io.Writer, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func ArtistIndexHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageData := struct {
		Artists  []Artist
		NextPage int
		PrevPage int
	}{
		QueryArtists(page),
		page + 1,
		page - 1,
	}

	render(w, "artists/index.html", pageData)
}
