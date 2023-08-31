CREATE TABLE IF NOT EXISTS segments (
    slug VARCHAR primary key
);

CREATE TABLE IF NOT EXISTS users (
    id BIGINT primary key
);

CREATE TABLE IF NOT EXISTS user_segments (
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE,
    segment_slug VARCHAR REFERENCES segments (slug) ON DELETE CASCADE,
    adding_time TIMESTAMP not null,
    removal_time TIMESTAMP,
    PRIMARY KEY (user_id, segment_slug)
);