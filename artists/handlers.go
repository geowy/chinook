package artists

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func render(w io.Writer, filename string, data interface{}) {
	log.Print("Rendering ", filename)
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

func EditArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		UpdateArtist(id, name)
		http.Redirect(w, r, "/artists", http.StatusSeeOther)
	} else {
		pageData := struct {
			Artist Artist
		}{
			QueryArtist(id),
		}

		render(w, "artists/edit.html", pageData)
	}
}
