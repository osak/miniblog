CREATE TABLE posts (
    id UUID PRIMARY KEY,
    slug VARCHAR(99) UNIQUE,
    title VARCHAR(99) NOT NULL,
    body TEXT NOT NULL,

    posted_at BIGINT,

    -- System columns
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX posts_posted_at_idx ON posts (posted_at);
