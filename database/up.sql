DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    id VARCHAR(32) PRIMARY KEY,
    text VARCHAR(255) NOT NULL
);