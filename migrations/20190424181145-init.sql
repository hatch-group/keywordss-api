-- +migrate Up

CREATE TABLE users (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE stories (
  id BIGINT NOT NULL AUTO_INCREMENT,
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  posted_time TIMESTAMP NOT NULL,
  keywords JSON NOT NULL,
  user_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (title(255), body(255)),
  FOREIGN KEY (user_id)
    REFERENCES users(id)
);

CREATE TABLE keywords (
  id BIGINT NOT NULL AUTO_INCREMENT,
  word VARCHAR(255) NOT NULL,
  type INT NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (word)
);

CREATE TABLE likes (
  id BIGINT NOT NULL AUTO_INCREMENT,
  user_id BIGINT NOT NULL,
  story_id BIGINT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id)
    REFERENCES users(id),
  FOREIGN KEY (story_id)
    REFERENCES stories(id)
);

-- +migrate Down

DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS keywords;
DROP TABLE IF EXISTS stories;
DROP TABLE IF EXISTS users;
