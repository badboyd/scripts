CREATE TABLE authors (
  id   INTEGER PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);


-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE people (id int);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE people;