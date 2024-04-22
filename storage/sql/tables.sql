CREATE TYPE IF NOT EXISTS position_enum AS ENUM ('PG', 'SG', 'SF', 'PF', 'C');
CREATE TYPE IF NOT EXISTS account_enum AS ENUM ('PLAYER', 'REFEREE', 'ADMIN');
CREATE TYPE IF NOT EXISTS category_enum AS ENUM ('U18', 'OPEN');
CREATE TYPE IF NOT EXISTS gender_enum AS ENUM ('MALE', 'FEMALE');

CREATE TABLE IF NOT EXISTS permissions (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    type account_enum NOT NULL DEFAULT 'PLAYER',
    flags BIT(8) NOT NULL DEFAULT B'00000000'
);

CREATE TABLE IF NOT EXISTS cities (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    state TEXT NOT NULL,
    INDEX cities_name (name)
);

CREATE TABLE IF NOT EXISTS accounts (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    permissions_id BIGINT NOT NULL UNIQUE REFERENCES permissions (id) ON DELETE CASCADE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    token TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT now(),
    logged_on TIMESTAMP,
    verified BOOLEAN NOT NULL DEFAULT false,
    INDEX accounts_email (email),
    INDEX accounts_verified (verified),
    INDEX accounts_token (token)
);

CREATE TABLE IF NOT EXISTS recoveries (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    accounts_id BIGINT NOT NULL REFERENCES accounts (id) ON DELETE CASCADE,
    created_on TIMESTAMP NOT NULL DEFAULT now(),
    recovered_on TIMESTAMP DEFAULT NULL,
    verified BOOLEAN NOT NULL DEFAULT false,
    verification_code TEXT NOT NULL UNIQUE,
    INDEX accounts_verication_code (verification_code)
);

CREATE TABLE IF NOT EXISTS categories (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    name TEXT NOT NULL,
    gender gender_enum NOT NULL,
    team_limit INT NOT NULL DEFAULT 0,
    UNIQUE (name, gender),
    INDEX categories_name (name)
);

CREATE TABLE IF NOT EXISTS teams (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    cities_id BIGINT DEFAULT NULL REFERENCES cities (id) ON DELETE CASCADE,
    categories_id BIGINT DEFAULT NULL REFERENCES categories (id) ON DELETE CASCADE,
    name TEXT NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    gender gender_enum NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT now(),
    INDEX teams_name (name)
);

CREATE TABLE IF NOT EXISTS players (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    accounts_id BIGINT NOT NULL UNIQUE REFERENCES accounts (id) ON DELETE CASCADE,
    teams_id BIGINT DEFAULT NULL REFERENCES teams (id) ON DELETE CASCADE,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    full_name TEXT AS (CONCAT(first_name, ' ', last_name)) STORED,
    birthday TIMESTAMP NOT NULL,
    gender gender_enum NOT NULL,
    number SMALLINT NOT NULL,
    height SMALLINT NOT NULL,
    weight SMALLINT NOT NULL DEFAULT 0,
    wingspan SMALLINT NOT NULL DEFAULT 0,
    position position_enum NOT NULL,
    UNIQUE (teams_id, number),
    INDEX players_number (number)
);

CREATE TABLE IF NOT EXISTS teams_players (
    teams_id BIGINT NOT NULL REFERENCES teams (id) ON DELETE CASCADE,
    players_id BIGINT NOT NULL UNIQUE REFERENCES players (id) ON DELETE CASCADE,
    PRIMARY KEY (teams_id, players_id)
);

CREATE TABLE IF NOT EXISTS audits (
    id BIGINT NOT NULL UNIQUE DEFAULT unique_rowid() PRIMARY KEY,
    time TIMESTAMP NOT NULL DEFAULT NOW(),
    status SMALLINT NOT NULL DEFAULT 0,
    message TEXT NOT NULL
);