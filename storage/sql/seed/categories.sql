BEGIN;

INSERT INTO categories (name, gender, team_limit) VALUES
    ('U18', 'MALE', 8),
    ('OPEN', 'MALE', 24)
    ON CONFLICT DO NOTHING;

COMMIT;