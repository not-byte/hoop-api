BEGIN;

INSERT INTO teams (name, description, categories_id, cities_id)
VALUES
    ('Bears', 'A fierce and determined team known for their powerful defense.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bielsko-Biała')),
    ('Bulldogs', 'A tenacious team with a never-give-up attitude.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bieruń')),
    ('Blazers', 'An agile and fast-paced team that loves to score quickly.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Blachownia')),
    ('Ballers', 'A skilled and versatile team known for their ball-handling skills.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bytom')),
    ('Bucks', 'A well-rounded team with a strategic approach to the game.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Będzin')),
    ('Chargers', 'A high-energy team known for their aggressive offense.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Chorzów')),
    ('Cyclones', 'A dynamic team with a whirlwind style of play.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Cieszyn')),
    ('Dragons', 'A fierce and relentless team with a fiery spirit.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czechowice-Dziedzice')),
    ('Cougars', 'A stealthy and strategic team known for their quick reflexes.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czeladź')),
    ('Crusaders', 'A brave and disciplined team with a strong defensive lineup.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czerwionka-Leszczyny')),
    ('Comets', 'A fast and elusive team known for their quick transitions.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Częstochowa')),
    ('Dolphins', 'A smooth and agile team with excellent coordination.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Dąbrowa Górnicza')),
    ('Giants', 'A powerhouse team known for their strength and endurance.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Gliwice')),
    ('Ice', 'A cool and composed team with a focus on precision.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Imielin')),
    ('Jaguars', 'A sleek and agile team with a predatory instinct.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Jastrzębie-Zdrój')),
    ('Jets', 'A fast and high-flying team with a knack for speed.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Jaworzno')),
    ('Knights', 'A noble and steadfast team with strong teamwork.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Kalety')),
    ('Kings', 'A dominant and authoritative team known for their control of the game.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Katowice')),
    ('Kangaroos', 'A lively and energetic team with great jumping ability.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Knurów')),
    ('Kestrels', 'A sharp and keen team with excellent vision on the field.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Koniecpol')),
    ('Kodiaks', 'A strong and resilient team with a robust defense.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Koziegłowy')),
    ('Keepers', 'A vigilant and protective team with an unbreakable defense.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Krzanowice')),
    ('Kites', 'A high-flying team with a knack for soaring above the competition.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Krzepice')),
    ('Raptors', 'A fast and aggressive team with sharp instincts.', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Kuźnia Raciborska'))
    ON CONFLICT DO NOTHING;

COMMIT;
