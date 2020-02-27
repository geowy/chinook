SELECT
  Album.AlbumId AS AlbumId,
  Album.Title AS AlbumTitle,
  Artist.ArtistId AS ArtistId,
  Artist.Name AS ArtistName
FROM Album
JOIN Artist USING (ArtistId)
WHERE Album.AlbumId = :id
LIMIT 1
;
