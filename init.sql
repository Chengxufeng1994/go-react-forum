CREATE DATABASE IF NOT EXISTS go_react_forum DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `go_react_forum`;

DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    username   VARCHAR(20)  NOT NULL,
    email      VARCHAR(100) NOT NULL UNIQUE,
    password   VARCHAR(100) NOT NULL,
    created_at datetime     NOT NULL,
    updated_at datetime     NOT NULL
) DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;

CREATE TABLE IF NOT EXISTS sessions
(
    session_id VARCHAR(100) PRIMARY KEY,
    username   VARCHAR(100) NOT NULL,
    user_id    INT          NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
