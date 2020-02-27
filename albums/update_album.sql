UPDATE Album
SET
  Title = :title,
  ArtistId = :artistid
WHERE AlbumId = :id
;
