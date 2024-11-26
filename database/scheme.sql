-- Active: 1732632262439@@127.0.0.1@5432@olympics
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(64) PRIMARY KEY,
    email VARCHAR(128) UNIQUE NOT NULL,
    username VARCHAR(64) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL
);

CREATE TYPE difficulty AS ENUM ('beginners', 'intermediates', 'professionals');

CREATE TABLE IF NOT EXISTS exercises (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    image_path VARCHAR(256) NOT NULL
);

CREATE TYPE parameter_type AS ENUM ('duration', 'times');

CREATE TABLE IF NOT EXISTS exercise_parameter (
    exercise_id BIGINT REFERENCES exercises ON DELETE CASCADE, 
    type PARAMETER_TYPE NOT NULL,
    value VARCHAR(64),
    PRIMARY KEY (exercise_id, type)
);

CREATE TABLE IF NOT EXISTS trainings ();