CREATE TABLE bundles (
    id SERIAL PRIMARY KEY,

    priority INT DEFAULT 0,
    size INT NOT NULL,
    status INT DEFAULT 0
)