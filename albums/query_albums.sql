SELECT
  Album.AlbumId AS AlbumId,
  Album.Title AS AlbumTitle,
  Artist.ArtistId AS ArtistId,
  Artist.Name AS ArtistName,
  COUNT(Track.TrackId) AS TrackCount,
  SUM(Track.Milliseconds) AS Milliseconds
FROM Album
JOIN Artist USING (ArtistId)
JOIN Track USING (AlbumId)
GROUP BY Album.AlbumId
ORDER BY Album.Title
LIMIT 25
OFFSET (:page - 1) * 25
;
