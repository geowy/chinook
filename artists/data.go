package artists

import (
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
