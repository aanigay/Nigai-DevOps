CREATE TABLE IF NOT EXISTS users
(
    id       serial PRIMARY KEY,
    name     varchar(64) NOT NULL UNIQUE,
    password varchar(64) NOT NULL,
    role     varchar(32) NOT NULL
);

INSERT INTO users (name, password, role)
VALUES ('root', 'root', 'admin');

CREATE TABLE IF NOT EXISTS posts
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    body        TEXT         NOT NULL,
    createdAt   TIMESTAMP    NOT NULL DEFAULT NOW(),
    userId      BIGINT       NOT NULL,
    commentsIds BIGINT[]     NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS comments
(
    id        SERIAL PRIMARY KEY,
    body      TEXT      NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
    userId    BIGINT    NOT NULL,
    postId    BIGINT    NOT NULL
)