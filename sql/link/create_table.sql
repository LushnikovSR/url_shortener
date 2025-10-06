CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    original_link TEXT NOT NULL,
    short_link TEXT NOT NULL
);