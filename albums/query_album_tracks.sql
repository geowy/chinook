SELECT
  Track.Name AS Name,
  Genre.Name AS Genre,
  Track.Composer AS Composer,
  Track.Milliseconds AS Milliseconds
FROM Track
JOIN Genre USING (GenreId)
WHERE Track.AlbumId = :id
;
