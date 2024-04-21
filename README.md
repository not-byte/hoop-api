# tournament-api

# Basketball Tournament API

### Storage

sql cockroach database

multiple nodes

transactions

denormalization

database schema

redis cache

### Service

golang rest api

versioning

gzip compression

performance monitor

tests

cdn for photos and other resources

### Environment

docker containers

load balancer

kubernetes

99.9% availability

## Service Design

## Database Schema

This section is an overview of the database schema for the service.
The schema includes types and tables for managing REST API data, following:

- [Types](#schema-types)
- [Subscriptions](#subscription-table)
- [Tenancies](#tenancy-table)
- [Accounts](#account-table)
- [Referees](#referee-table)
- [Teams](#team-table)
- [Players](#player-table)
- [Games](#game-table)
- [Locations](#location-table)
- [Tournaments](#tournament-table)
- [Statistics](#statistics-table)
- [Categories](#category-table)
- [Rosters](#roster-table)
- [Game Referees](#game-referee-table)
- [Logs](#log-table)

### Schema Types

| Type        | Description             |
|-------------|-------------------------|
| Coordinates | (lat FLOAT, long FLOAT) |


### Subscription Table

| Field               | Type     | Description                              |
|---------------------|----------|------------------------------------------|
| `id`                | INT      | Unique identifier for the subscription.  |
| `name`              | VARCHAR  | Name of the subscription plan.           |
| `has_key`           | BOOLEAN  | Indicates if the subscription has a key. |
| `price`             | FLOAT    | Price of the subscription.               |
| `description`       | VARCHAR  | Description of the subscription.         |
| `quota_tournaments` | SMALLINT | Quota for tournaments.                   |
| `quota_players`     | SMALLINT | Quota for players.                       |
| `quota_teams`       | SMALLINT | Quota for teams.                         |

### Tenancy Table

| Field             | Type      | Description                             |
|-------------------|-----------|-----------------------------------------|
| `id`              | INT       | Unique identifier for the tenancy.      |
| `subscription_id` | INT       | References the associated subscription. |
| `name`            | VARCHAR   | Name of the tenancy.                    |
| `key`             | VARCHAR   | Unique key for the tenancy.             |
| `created`         | TIMESTAMP | Timestamp of tenancy creation.          |

### Account Table

| Field                 | Type        | Description                            |
|-----------------------|-------------|----------------------------------------|
| `id`                  | INT         | Unique identifier for the account.     |
| `email`               | VARCHAR     | Email of the account.                  |
| `login`               | VARCHAR     | Login username.                        |
| `password`            | VARCHAR     | Password for account authentication.   |
| `type`                | SMALLINT    | Type of the account.                   |
| `is_activated`        | BOOLEAN     | Indicates if the account is activated. |
| `activated_timestamp` | TIMESTAMP   | Timestamp of account activation.       |
| `activation_key`      | VARCHAR     | Unique activation key for the account. |
| `created`             | TIMESTAMP   | Timestamp of account creation.         |
| `last_logon_datetime` | TIMESTAMP   | Timestamp of last login.               |
| `last_logon_location` | COORDINATES | Last login location coordinates.       |

### Referee Table

| Field        | Type      | Description                        |
|--------------|-----------|------------------------------------|
| `id`         | INT       | Unique identifier for the referee. |
| `account_id` | INT       | References the associated account. |
| `created`    | TIMESTAMP | Timestamp of referee creation.     |

### Team Table

| Field        | Type      | Description                        |
|--------------|-----------|------------------------------------|
| `id`         | INT       | Unique identifier for the team.    |
| `tenancy_id` | INT       | References the associated tenancy. |
| `owner_id`   | INT       | References the account owner.      |
| `name`       | VARCHAR   | Name of the team.                  |
| `created`    | TIMESTAMP | Timestamp of team creation.        |

### Player Table

| Field               | Type     | Description                           |
|---------------------|----------|---------------------------------------|
| `id`                | INT      | Unique identifier for the player.     |
| `account_id`        | INT      | References the associated account.    |
| `team_id`           | INT      | References the associated team.       |
| `first_name`        | VARCHAR  | First name of the player.             |
| `last_name`         | VARCHAR  | Last name of the player.              |
| `full_name`         | VARCHAR  | Concatenation of first and last name. |
| `position`          | SMALLINT | Player's position.                    |
| `height`            | SMALLINT | Player's height.                      |
| `gender`            | BOOLEAN  | Player's gender.                      |
| `signed_agreement`  | BOOLEAN  | Indicates player agreement.           |

### Game Table

| Field               | Type     | Description                            |
|---------------------|----------|----------------------------------------|
| `id`                | INT      | Unique identifier for the game.        |
| `arena_id`          | INT      | References the associated arena.       |
| `tournament_id`     | INT      | References the associated tournament.  |
| `home_id`           | INT      | References the home team's roster.     |
| `opponent_id`       | INT      | References the opponent team's roster. |

### Location Table

| Field         | Type        | Description                         |
|---------------|-------------|-------------------------------------|
| `id`          | INT         | Unique identifier for the location. |
| `city_id`     | INT         | References the associated city.     |
| `coordinates` | COORDINATES | Latitude and longitude coordinates. |

### Tournament Table

| Field         | Type      | Description                           |
|---------------|-----------|---------------------------------------|
| `id`          | INT       | Unique identifier for the tournament. |
| `arena_id`    | INT       | References the associated arena.      |
| `category_id` | INT       | References the associated category.   |
| `tenancy_id`  | INT       | References the associated tenancy.    |
| `name`        | VARCHAR   | Name of the tournament.               |
| `starting`    | TIMESTAMP | Date of the tournament.               |
| `max_teams`   | SMALLINT  | Maximum number of teams allowed.      |
| `created`     | TIMESTAMP | Timestamp of tournament creation.     |

### Category Table

| Field               | Type     | Description                         |
|---------------------|----------|-------------------------------------|
| `id`                | INT      | Unique identifier for the category. |
| `gender`            | BOOLEAN  | Gender category of the tournament.  |
| `type`              | SMALLINT | Type of category.                   |

### Roster Table

| Field               | Type     | Description                       |
|---------------------|----------|-----------------------------------|
| `id`                | INT      | Unique identifier for the roster. |
| `team_id`           | INT      | References the associated team.   |

### Game Referee Table

| Field           | Type | Description                        |
|-----------------|------|------------------------------------|
| `game_id`       | INT8 | References the associated game.    |
| `referee_id`    | INT8 | References the associated referee. |

### Log Table

| Field       | Type      | Description                          |
|-------------|-----------|--------------------------------------|
| `id`        | INT8      | Unique identifier for the log entry. |
| `game_id`   | INT8      | References the associated game.      |
| `timestamp` | TIMESTAMP | Timestamp of the log entry.          |
| `message`   | VARCHAR   | Log message.                         |
