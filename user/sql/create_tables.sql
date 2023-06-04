CREATE DATABASE user_service;

\c user_service;

SELECT current_database();

CREATE TABLE IF NOT EXISTS "user" (
  username text NOT NULL,
  "password" text NOT NULL,
  PRIMARY KEY (username)
);
