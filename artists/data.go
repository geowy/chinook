package artists

import (
	"errors"
	"github.com/geowy/chinook/data"
)

type Artist struct {
	ArtistId int
	Name     string
}

func QueryArtists(page int) []Artist {
	rows := data.Query("artists/query_artists.sql", page)
	defer rows.Close()

	artists := []Artist{}
	for rows.Next() {
		artist := Artist{}

		err := rows.Scan(
			&artist.ArtistId,
			&artist.Name,
		)
		if err != nil {
			panic(err)
		}

		artists = append(artists, artist)
	}

	return artists
}

func QueryArtist(id int) Artist {
	row := data.Query("artists/query_artist.sql", id)
	defer row.Close()

	if !row.Next() {
		panic(errors.New("Album not found"))
	}

	artist := Artist{}
	err := row.Scan(
		&artist.ArtistId,
		&artist.Name,
	)
	if err != nil {
		panic(err)
	}

	return artist
}

func UpdateArtist(id int, name string) {
	data.Exec("artists/update_artist.sql", name, id)
}
