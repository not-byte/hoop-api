BEGIN;

CREATE TYPE IF NOT EXISTS position_enum AS ENUM ('PG', 'SG', 'SF', 'PF', 'C');
CREATE TYPE IF NOT EXISTS category_enum AS ENUM ('U18', 'OPEN');
CREATE TYPE IF NOT EXISTS gender_enum AS ENUM ('MALE', 'FEMALE');

CREATE TABLE IF NOT EXISTS categories (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    name TEXT NOT NULL,
    gender gender_enum NOT NULL,
    team_limit INT NOT NULL DEFAULT 0,
    UNIQUE (name, gender),
    INDEX categories_name (name)
);

CREATE TABLE IF NOT EXISTS cities (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    state TEXT NOT NULL,
    INDEX cities_name (name)
);

CREATE TABLE IF NOT EXISTS teams (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    categories_id BIGINT DEFAULT NULL REFERENCES categories (id) ON DELETE CASCADE,
    cities_id BIGINT DEFAULT NULL REFERENCES cities (id) ON DELETE CASCADE,
    name TEXT NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT now(),
    INDEX teams_name (name)
);

CREATE TABLE IF NOT EXISTS players (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    teams_id BIGINT DEFAULT NULL REFERENCES teams (id) ON DELETE CASCADE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    full_name TEXT AS (CONCAT(first_name, ' ', last_name)) STORED,
    birthday DATE,
    gender gender_enum NOT NULL,
    number SMALLINT NOT NULL,
    height SMALLINT NOT NULL,
    weight SMALLINT DEFAULT NULL,
    wingspan SMALLINT DEFAULT NULL,
    position position_enum NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (teams_id, number),
    INDEX players_number (number)
);

CREATE TABLE IF NOT EXISTS teams_players (
    teams_id BIGINT NOT NULL REFERENCES teams (id) ON DELETE CASCADE,
    players_id BIGINT NOT NULL UNIQUE REFERENCES players (id) ON DELETE CASCADE,
    PRIMARY KEY (teams_id, players_id)
);

COMMIT;

