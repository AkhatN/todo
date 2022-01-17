DROP TABLE IF EXISTS todo;
CREATE TABLE IF NOT EXISTS todo (
    id SERIAL,
    description TEXT NOT NULL,
    created_at TEXT NOT NULL,
    completed_at TEXT,
    PRIMARY KEY (id)
);