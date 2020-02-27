package albums

import (
	"errors"
	"github.com/geowy/chinook/data"
)

type Album struct {
	AlbumId      int
	Title        string
	ArtistId     int
	ArtistName   string
	TrackCount   int
	Milliseconds int
	Tracks       []Track
}

type Track struct {
	Name         string
	Genre        string
	Composer     *string
	Milliseconds int
}

type Artist struct {
	ArtistId int
	Name     string
}

func QueryAlbums(page int) []Album {
	rows := data.Query("albums/query_albums.sql", page)
	defer rows.Close()

	albums := []Album{}
	for rows.Next() {
		album := Album{}

		err := rows.Scan(
			&album.AlbumId,
			&album.Title,
			&album.ArtistId,
			&album.ArtistName,
			&album.TrackCount,
			&album.Milliseconds)
		if err != nil {
			panic(err)
		}

		albums = append(albums, album)
	}

	return albums
}

func QueryAlbum(id int) Album {
	albumRow := data.Query("albums/query_album.sql", id)
	defer albumRow.Close()

	if !albumRow.Next() {
		panic(errors.New("Album not found"))
	}

	album := Album{}
	err := albumRow.Scan(
		&album.AlbumId,
		&album.Title,
		&album.ArtistId,
		&album.ArtistName)
	if err != nil {
		panic(err)
	}

	trackRows := data.Query("albums/query_album_tracks.sql", id)
	if err != nil {
		panic(err)
	}
	defer trackRows.Close()

	for trackRows.Next() {
		track := Track{}

		err = trackRows.Scan(
			&track.Name,
			&track.Genre,
			&track.Composer,
			&track.Milliseconds)
		if err != nil {
			panic(err)
		}

		album.Tracks = append(album.Tracks, track)
	}

	return album
}

func QueryArtists() []Artist {
	rows := data.Query("albums/query_artists.sql")
	defer rows.Close()

	artists := []Artist{}
	for rows.Next() {
		artist := Artist{}

		err := rows.Scan(&artist.ArtistId, &artist.Name)
		if err != nil {
			panic(err)
		}

		artists = append(artists, artist)
	}

	return artists
}

func UpdateAlbum(id int, title string, artistId int) {
	data.Exec("albums/update_album.sql", title, artistId, id)
}
