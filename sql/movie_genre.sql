-- Example case create statement:
CREATE TABLE genres (
  id INTEGER PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE movies (
  id INTEGER PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE moviesGenres (
  movieId INTEGER REFERENCES movies(id),
  genreId INTEGER REFERENCES genres(id),
  PRIMARY KEY(movieId, genreId),
  FOREIGN KEY (movieId) REFERENCES movies(id)
  FOREIGN KEY (genreId) REFERENCES genres(id)
);

-- find movies which is not comedy
INSERT INTO genres(id, name) VALUES(1, 'Comedy');
INSERT INTO genres(id, name) VALUES(2, 'Action');
INSERT INTO genres(id, name) VALUES(3, 'Thriller');
INSERT INTO genres(id, name) VALUES(4, 'Science Fiction');

INSERT INTO movies(id, name) VALUES(1, 'Greater Boys');
INSERT INTO movies(id, name) VALUES(2, 'Men In Black');
INSERT INTO movies(id, name) VALUES(3, 'Taxi Night');

INSERT INTO moviesGenres(movieId, genreId) VALUES(1, 1);
INSERT INTO moviesGenres(movieId, genreId) VALUES(1, 2);
INSERT INTO moviesGenres(movieId, genreId) VALUES(1, 3);
INSERT INTO moviesGenres(movieId, genreId) VALUES(1, 4);
INSERT INTO moviesGenres(movieId, genreId) VALUES(2, 2);
INSERT INTO moviesGenres(movieId, genreId) VALUES(2, 3);
INSERT INTO moviesGenres(movieId, genreId) VALUES(2, 4);
INSERT INTO moviesGenres(movieId, genreId) VALUES(3, 2);
INSERT INTO moviesGenres(movieId, genreId) VALUES(3, 3);

-- Expected output (rows in any order):
-- name          noOfGenres
-- ------------------------
-- Source Code   3


-- munculkan movies dengan genres
-- 