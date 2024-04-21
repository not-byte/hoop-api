DROP DATABASE IF EXISTS tournament_api;

CREATE DATABASE IF NOT EXISTS tournament_api;

SET DATABASE = tournament_api;

CREATE TYPE IF NOT EXISTS COORDINATES AS (lat FLOAT, long FLOAT);

CREATE TABLE IF NOT EXISTS subscription (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    name VARCHAR(32) NOT NULL UNIQUE,
    has_key BOOLEAN DEFAULT false,
    price DECIMAL(3, 2) DEFAULT 0,
    description VARCHAR(64) NOT NULL,
    quota_tournaments INT2 NOT NULL DEFAULT 0,
    quota_players INT2 NOT NULL DEFAULT 0,
    quota_teams INT2 NOT NULL DEFAULT 0,
    INDEX name_idx (name)
);

CREATE TABLE IF NOT EXISTS tenancy (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    subscription_id INT8 NOT NULL REFERENCES subscription(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL UNIQUE,
    key VARCHAR UNIQUE,
    created TIMESTAMP DEFAULT current_timestamp(),
    INDEX subscription_idx (subscription_id)
);

CREATE TABLE IF NOT EXISTS account (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    email VARCHAR(64) NOT NULL UNIQUE,
    login VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    type INT2 DEFAULT 0,
    is_activated BOOLEAN DEFAULT false,
    activated_timestamp TIMESTAMP,
    activation_key VARCHAR(64) NOT NULL UNIQUE,
    created TIMESTAMP DEFAULT current_timestamp(),
    last_logon_datetime TIMESTAMP,
    last_logon_location COORDINATES,
    INDEX email_idx (email),
    INDEX login_idx (login),
    INDEX type_idx (type)
);

CREATE TABLE IF NOT EXISTS referee (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    account_id INT8 NOT NULL UNIQUE REFERENCES account(id) ON DELETE CASCADE,
    INDEX account_idx (account_id)
);

CREATE TABLE IF NOT EXISTS team (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    tenancy_id INT8 NOT NULL REFERENCES tenancy(id) ON DELETE CASCADE,
    owner_id INT8 NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    name VARCHAR(32) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp(),
    INDEX tenancy_idx (tenancy_id),
    INDEX owner_idx (owner_id)
);

CREATE TABLE IF NOT EXISTS state (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    name VARCHAR(32) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS city (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    state_id INT8 NOT NULL REFERENCES state(id) ON DELETE CASCADE,
    name VARCHAR(32) NOT NULL,
    INDEX state_idx (state_id)
);

CREATE TABLE IF NOT EXISTS location (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    city_id INT8 NOT NULL REFERENCES city(id) ON DELETE CASCADE,
    coordinates COORDINATES NOT NULL,
    INDEX city_idx (city_id)
);

CREATE TABLE IF NOT EXISTS arena (
     id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
     location_id INT8 NOT NULL REFERENCES location(id) ON DELETE CASCADE,
     name VARCHAR(32) NOT NULL,
     url VARCHAR(128) NOT NULL,
     INDEX location_inx (location_id)
);

CREATE TABLE IF NOT EXISTS category (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    gender BOOLEAN DEFAULT false,
    type INT2 DEFAULT 0
);

CREATE TABLE IF NOT EXISTS tournament (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    arena_id INT8 NOT NULL REFERENCES arena(id) ON DELETE CASCADE,
    category_id INT8 NOT NULL REFERENCES category(id) ON DELETE CASCADE,
    tenancy_id INT8 NOT NULL REFERENCES tenancy(id) ON DELETE CASCADE,
    name VARCHAR(32) NOT NULL UNIQUE,
    starting TIMESTAMP NOT NULL,
    max_teams INT2 DEFAULT 8,
    created TIMESTAMP DEFAULT current_timestamp(),
    INDEX arena_idx (arena_id),
    INDEX category_idx (category_id),
    INDEX tenancy_idx (tenancy_id)
);

CREATE TABLE IF NOT EXISTS tournament_category (
    tournament_id INT8 NOT NULL REFERENCES tournament(id) ON DELETE CASCADE,
    category_id INT8 NOT NULL REFERENCES category(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS player (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    account_id INT8 NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    team_id INT8 REFERENCES team(id),
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    full_name VARCHAR(64) AS (CONCAT(first_name, ' ', last_name)) STORED,
    position INT2 NOT NULL,
    height INT2 NOT NULL,
    gender BOOL NOT NULL DEFAULT false,
    signed_agreement BOOLEAN DEFAULT false,
    INDEX account_idx (account_id),
    INDEX team_idx (team_id)
);

CREATE TABLE IF NOT EXISTS roster (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    team_id INT8 NOT NULL REFERENCES team(id) ON DELETE CASCADE,
    INDEX team_idx (team_id)
);

CREATE TABLE IF NOT EXISTS roster_player (
    player_id INT8 NOT NULL REFERENCES player(id) ON DELETE CASCADE,
    roster_id INT8 NOT NULL REFERENCES roster(id) ON DELETE CASCADE,
    INDEX player_idx (player_id),
    INDEX roster_idx (roster_id),
    UNIQUE (roster_id, player_id)
);

CREATE TABLE IF NOT EXISTS game (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    arena_id INT8 NOT NULL REFERENCES arena(id) ON DELETE CASCADE,
    tournament_id INT8 NOT NULL REFERENCES tournament(id) ON DELETE CASCADE,
    home_id INT8 NOT NULL REFERENCES roster(id) ON DELETE CASCADE,
    opponent_id INT8 NOT NULL REFERENCES roster(id) ON DELETE CASCADE,
    INDEX arena_id (arena_id),
    INDEX tournament_idx (tournament_id),
    INDEX home_idx (home_id),
    INDEX opponent_idx (opponent_id),
    UNIQUE (home_id, opponent_id)
);

CREATE TABLE IF NOT EXISTS game_referee (
    game_id INT8 NOT NULL REFERENCES game(id) ON DELETE CASCADE,
    referee_id INT8 NOT NULL REFERENCES referee(id) ON DELETE CASCADE,
    INDEX game_idx (game_id),
    INDEX referee_idx (referee_id),
    UNIQUE (referee_id, game_id)
);

CREATE TABLE IF NOT EXISTS statistics (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    game_id INT8 NOT NULL REFERENCES game(id) ON DELETE CASCADE,
    player_id INT8 NOT NULL REFERENCES player(id) ON DELETE CASCADE,
    assists INT2 NOT NULL DEFAULT 0,
    rebounds_offensive INT2 NOT NULL DEFAULT 0,
    rebounds_defensive INT2 NOT NULL DEFAULT 0,
    field_goal_attempted INT2 NOT NULL DEFAULT 0,
    field_goal_made INT2 NOT NULL DEFAULT 0,
    three_point_attempted INT2 NOT NULL DEFAULT 0,
    three_point_made INT2 NOT NULL DEFAULT 0,
    free_throw_attempted INT2 NOT NULL DEFAULT 0,
    free_throw_made INT2 NOT NULL DEFAULT 0,
    steals INT2 NOT NULL DEFAULT 0,
    blocks INT2 NOT NULL DEFAULT 0,
    fouls INT2 NOT NULL DEFAULT 0,
    turnovers INT2 NOT NULL DEFAULT 0,
    minutes INT2 NOT NULL DEFAULT 0,
    INDEX game_idx (game_id),
    INDEX player_idx (player_id),
    UNIQUE (game_id, player_id)
);

CREATE TABLE IF NOT EXISTS schedule (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    game_id INT8 NOT NULL UNIQUE REFERENCES game(id) ON DELETE CASCADE,
    starting TIMESTAMP,
    INDEX game_idx (game_id)
);

CREATE TABLE IF NOT EXISTS log (
    id INT8 NOT NULL UNIQUE PRIMARY KEY DEFAULT unordered_unique_rowid(),
    game_id INT8 NOT NULL REFERENCES game(id) ON DELETE CASCADE,
    timestamp TIMESTAMP NOT NULL DEFAULT current_timestamp(),
    message VARCHAR,
    INDEX game_idx (game_id)
);

CREATE OR REPLACE VIEW games_playing AS
    SELECT game.id,
           schedule.starting
    FROM game, schedule
    WHERE game.id = schedule.game_id
    AND schedule.starting > current_timestamp();

CREATE OR REPLACE VIEW teams_playing AS
    SELECT game.id,
           roster.id as roster_id,
           team.name
    FROM game, roster, team
    WHERE roster.team_id = team.id;

CREATE OR REPLACE VIEW players_playing AS
    SELECT roster.id,
       player.id as player_id,
       player.full_name
    FROM player, roster, roster_player
    WHERE player.id = roster_player.player_id
    AND roster.id = roster_player.roster_id
    AND player.signed_agreement = true;
