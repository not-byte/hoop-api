--> Drop tables
DROP TABLE IF EXISTS teams_players;
DROP TABLE IF EXISTS team_players;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS recoveries;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS audits;
DROP TABLE IF EXISTS categories;

--> drop customs types
DROP TYPE IF EXISTS position_enum;
DROP TYPE IF EXISTS account_enum;
DROP TYPE IF EXISTS category_enum;
DROP TYPE IF EXISTS gender_enum;