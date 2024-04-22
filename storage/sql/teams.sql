CREATE VIEW IF NOT EXISTS categories_quota AS
    SELECT categories.name AS name,
        categories.gender AS gender,
        count(teams.id) AS amount
    FROM categories, teams
    WHERE categories.id = teams.categories_id
    GROUP BY categories.id,
         categories.name,
         categories.gender;

INSERT INTO teams (name, gender, categories_id, cities_id)
VALUES
    ('Bears', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bielsko-Biała')),
    ('Bulldogs', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bieruń')),
    ('Blazers', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Blachownia')),
    ('Ballers', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bytom')),
    ('Bucks', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Będzin')),
    ('Chargers', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Chorzów')),
    ('Cyclones', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Cieszyn')),
    ('Dragons', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czechowice-Dziedzice')),
    ('Cougars', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czeladź')),
    ('Crusaders', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czerwionka-Leszczyny')),
    ('Comets', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Częstochowa')),
    ('Dolphins', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Dąbrowa Górnicza')),
    ('Giants', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Gliwice')),
    ('Ice', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Imielin')),
    ('Jaguars', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Jastrzębie-Zdrój')),
    ('Jets', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Jaworzno')),
    ('Knights', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Kalety')),
    ('Kings', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Katowice')),
    ('Kangaroos', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Knurów')),
    ('Kestrels', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Koniecpol')),
    ('Kodiaks', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Koziegłowy')),
    ('Keepers', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Krzanowice')),
    ('Kites', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Krzepice')),
    ('Raptors', 'MALE', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Kuźnia Raciborska'))
    ON CONFLICT DO NOTHING;