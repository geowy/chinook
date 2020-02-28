SELECT
  Artist.Artistid AS Artistid,
  Artist.Name AS Name
FROM Artist
WHERE Artist.ArtistId = :id
LIMIT 1
;
