# tournament-api

# Basketball Tournament API

### Storage

sql cockroach database

multiple nodes

transactions

denormalization

database schema

redis cache

# Service

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
The schema includes tables for managing REST API data, following:

- [Subscriptions,](#subscription-table)
- [Tenancies,](#tenancy-table)
- [Accounts,](#account-table)
- [Locations,](#location-table)
- [Tournaments,](#tournament-table)
- [Statistics,](#statistics-table)
- [Teams,](#team-table)
- [Players,](#player-table)
- [Games,](#game-table)
- [Referees,](#referee-table)

### Subscription Table

| Field               | Type     | Details                                  |
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

| Field             | Type      | Details                                 |
|-------------------|-----------|-----------------------------------------|
| `id`              | INT       | Unique identifier for the tenancy.      |
| `subscription_id` | INT       | References the associated subscription. |
| `name`            | VARCHAR   | Name of the tenancy.                    |
| `key`             | VARCHAR   | Unique key for the tenancy.             |
| `created`         | TIMESTAMP | Timestamp of tenancy creation.          |

### Account Table

| Field                 | Type        | Details                                |
|-----------------------|-------------|----------------------------------------|
| `id`                  | INT         | Unique identifier for the account.     |
| `email`               | VARCHAR     | Email of the account.                  |
| `login`               | VARCHAR     | Login username.                        |
| `password`            | VARCHAR     | Password for account authentication.   |
| `first_name`          | VARCHAR     | First name of the account holder.      |
| `last_name`           | VARCHAR     | Last name of the account holder.       |
| `full_name`           | VARCHAR     | Concatenation of first and last name.  |
| `type`                | SMALLINT    | Type of the account.                   |
| `is_activated`        | BOOLEAN     | Indicates if the account is activated. |
| `activated_timestamp` | TIMESTAMP   | Timestamp of account activation.       |
| `activation_key`      | VARCHAR     | Unique activation key for the account. |
| `created`             | TIMESTAMP   | Timestamp of account creation.         |
| `last_logon_datetime` | TIMESTAMP   | Timestamp of last login.               |
| `last_logon_location` | COORDINATES | Last login location coordinates.       |

### State Table

| Field               | Type     | Details                                |
|---------------------|----------|----------------------------------------|
| `id`                | INT      | Unique identifier for the state.       |
| `name`              | VARCHAR  | Name of the state.                     |

### City Table

| Field               | Type     | Details                                |
|---------------------|----------|----------------------------------------|
| `id`                | INT      | Unique identifier for the city.        |
| `state_id`          | INT      | References the associated state.       |
| `name`              | VARCHAR  | Name of the city.                      |

### Location Table

| Field         | Type        | Details                             |
|---------------|-------------|-------------------------------------|
| `id`          | INT         | Unique identifier for the location. |
| `city_id`     | INT         | References the associated city.     |
| `name`        | VARCHAR     | Name of the location.               |
| `coordinates` | COORDINATES | Latitude and longitude coordinates. |

### Tournament Table

| Field         | Type      | Details                               |
|---------------|-----------|---------------------------------------|
| `id`          | INT       | Unique identifier for the tournament. |
| `tenancy_id`  | INT       | References the associated tenancy.    |
| `location_id` | INT       | References the associated location.   |
| `name`        | VARCHAR   | Name of the tournament.               |
| `date`        | TIMESTAMP | Date of the tournament.               |
| `max_teams`   | SMALLINT  | Maximum number of teams allowed.      |
| `created`     | TIMESTAMP | Timestamp of tournament creation.     |

### Category Table

| Field               | Type     | Details                                |
|---------------------|----------|----------------------------------------|
| `tournament_id`     | INT      | References the associated tournament.  |
| `gender`            | SMALLINT | Gender category of the tournament.     |
| `type`              | SMALLINT | Type of category.                      |

### Statistics Table

| Field                   | Type | Details                               |
|-------------------------|------|---------------------------------------|
| `id`                    | INT  | Unique identifier for the statistics. |
| `assists`               | INT  |                                       |
| `rebounds_offensive`    | INT2 |                                       |
| `rebounds_defensive`    | INT2 |                                       |
| `field_goal_attempted`  | INT2 |                                       |
| `field_goal_made`       | INT  |                                       |
| `three_point_attempted` | INT  |                                       |
| `three_point_made`      | INT  |                                       |
| `free_throw_attempted`  | INT  |                                       |
| `free_throw_made`       | INT  |                                       |
| `steals`                | INT  |                                       |
| `blocks`                | INT  |                                       |
| `fouls`                 | INT  |                                       |
| `turnovers`             | INT  |                                       |
| `minutes`               | INT  |                                       |

### Team Table

| Field        | Type      | Details                            |
|--------------|-----------|------------------------------------|
| `id`         | INT       | Unique identifier for the team.    |
| `tenancy_id` | INT       | References the associated tenancy. |
| `owner_id`   | INT       | References the account owner.      |
| `name`       | VARCHAR   | Name of the team.                  |
| `created`    | TIMESTAMP | Timestamp of team creation.        |

### Team Statistics Table

| Field           | Type     | Details                               |
|-----------------|----------|---------------------------------------|
| `id`            | INT      | Unique identifier for the team stats. |
| `team_id`       | INT      | References the associated team.       |
| `game_id`       | INT      | References the associated game.       |
| `statistics_id` | INT      | References the associated statistics. |

### Player Table

| Field               | Type     | Details                                |
|---------------------|----------|----------------------------------------|
| `id`                | INT      | Unique identifier for the player.      |
| `account_id`        | INT      | References the associated account.     |
| `team_id`           | INT      | References the associated team.        |
| `position`          | SMALLINT | Player's position.                     |
| `height`            | SMALLINT | Player's height.                       |
| `agreement`         | BOOLEAN  | Indicates player agreement.            |

### Player Statistics Table

| Field           | Type | Details                                 |
|-----------------|------|-----------------------------------------|
| `id`            | INT  | Unique identifier for the player stats. |
| `player_id`     | INT  | References the associated player.       |
| `game_id`       | INT  | References the associated game.         |
| `statistics_id` | INT  | References the associated statistics.   |

### Roster Table

| Field               | Type     | Details                                |
|---------------------|----------|----------------------------------------|
| `id`                | INT      | Unique identifier for the roster.      |
| `team_id`           | INT      | References the associated team.        |

### Game Table

| Field               | Type     | Details                                |
|---------------------|----------|----------------------------------------|
| `id`                | INT      | Unique identifier for the game.        |
| `tournament_id`     | INT      | References the associated tournament.  |
| `home_id`           | INT      | References the home team's roster.     |
| `opponent_id`       | INT      | References the opponent team's roster. |

### Schedule Table

| Field      | Type      | Details                             |
|------------|-----------|-------------------------------------|
| `id`       | INT       | Unique identifier for the schedule. |
| `game_id`  | INT       | References the associated game.     |
| `starting` | TIMESTAMP | Starting time of the game.          |

### Log Table

| Field       | Type      | Details                              |
|-------------|-----------|--------------------------------------|
| `id`        | INT       | Unique identifier for the log entry. |
| `game_id`   | INT       | References the associated game.      |
| `timestamp` | TIMESTAMP | Timestamp of the log entry.          |
| `message`   | VARCHAR   | Log message.                         |

### Referee Table

| Field        | Type      | Details                            |
|--------------|-----------|------------------------------------|
| `id`         | INT       | Unique identifier for the referee. |
| `account_id` | INT       | References the associated account. |
| `created`    | TIMESTAMP | Timestamp of referee creation.     |
