CREATE TABLE posts (
  id      	SERIAL PRIMARY KEY,
  title		VARCHAR(255) NOT NULL,
  content 	TEXT NOT NULL,
  author  	VARCHAR(255) NOT NULL
);
