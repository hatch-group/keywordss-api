DROP TABLE IF EXISTS story_keyword;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS keywords;
DROP TABLE IF EXISTS stories;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name)
);

CREATE TABLE stories (
  id BIGINT NOT NULL,
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  status INT NOT NULL,
  created DATETIME NOT NULL,
  updated DATETIME NOT NULL,
  count INT NOT NULL,
  user_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (title(255), body(255)),
  FOREIGN KEY (user_id)
    REFERENCES users(id)
);

CREATE TABLE keywords (
  id BIGINT NOT NULL,
  word VARCHAR(255) NOT NULL,
  type INT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (word)
);

CREATE TABLE likes (
  id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  story_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id)
    REFERENCES users(id),
  FOREIGN KEY (story_id)
    REFERENCES stories(id)
);

CREATE TABLE story_keyword (
  id BIGINT NOT NULL,
  story_id BIGINT NOT NULL,
  keyword_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (story_id)
    REFERENCES stories(id),
  FOREIGN KEY (keyword_id)
    REFERENCES keywords(id)
);


