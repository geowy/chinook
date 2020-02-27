package albums

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

func AlbumIndexHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageData := struct {
		Albums   []Album
		NextPage int
		PrevPage int
	}{
		QueryAlbums(page),
		page + 1,
		page - 1}

	render(w, "albums/index.html", pageData)
}

func ShowAlbumHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	pageData := struct {
		Album Album
	}{
		QueryAlbum(id)}

	render(w, "albums/show.html", pageData)
}

func EditAlbumHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	if r.Method == "POST" {
		title := r.FormValue("title")
		artistId, err := strconv.Atoi(r.FormValue("artist-id"))
		if err != nil {
			panic(err)
		}

		UpdateAlbum(id, title, artistId)

		http.Redirect(w, r, "/albums/show?id="+idStr, http.StatusSeeOther)
	} else {
		pageData := struct {
			Album   Album
			Artists []Artist
		}{
			QueryAlbum(id),
			QueryArtists()}

		render(w, "albums/edit.html", pageData)
	}
}
