package main

import (
	"github.com/geowy/chinook/albums"
	"github.com/geowy/chinook/artists"
	"github.com/geowy/chinook/data"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Unknown route " + r.URL.Path)
	http.Redirect(w, r, "/albums", http.StatusMovedPermanently)
}

func logging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Started " + r.Method + " " + r.URL.String())
		handler(w, r)
		log.Print("Finished " + r.Method + " " + r.URL.String())
	}
}

func panicRecovery(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Print(err)
			}
		}()
		handler(w, r)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	data.Start()
	defer data.Stop()

	http.HandleFunc("/", logging(rootHandler))
	http.HandleFunc("/albums", panicRecovery(logging(albums.AlbumIndexHandler)))
	http.HandleFunc("/albums/show", panicRecovery(logging(albums.ShowAlbumHandler)))
	http.HandleFunc("/albums/edit", panicRecovery(logging(albums.EditAlbumHandler)))
	http.HandleFunc("/artists", panicRecovery(logging(artists.ArtistIndexHandler)))

	log.Print("Server listening on http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
