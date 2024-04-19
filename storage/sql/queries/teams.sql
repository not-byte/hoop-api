CREATE VIEW IF NOT EXISTS categories_quota AS
    SELECT category.name AS name,
    count(teams.id) AS amount,
    --LIMIT SECTION HERE
    FROM categories, teams
    WHERE categories.id = teams.categories_id
    GROUP BY categories.id;