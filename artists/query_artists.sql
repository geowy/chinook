SELECT
  Artist.ArtistId AS ArtistId,
  Artist.Name AS Name
FROM Artist
ORDER BY Artist.Name
LIMIT 25
OFFSET (:page - 1) * 25
;
