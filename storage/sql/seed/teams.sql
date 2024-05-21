BEGIN;

INSERT INTO teams (name, email, phone, categories_id, cities_id)
VALUES
    ('Bears', 'bears@bielsko-biala.pl', '789123456', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bielsko-Biała')),
    ('Bulldogs', 'bulldogs@bierun.pl', '800234567', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bieruń')),
    ('Blazers', 'blazers@blachownia.pl', '923456789', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Blachownia')),
    ('Ballers', 'ballers@bytom.pl', '701234567', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Bytom')),
    ('Bucks', 'bucks@bedzin.pl', '998765432', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Będzin')),
    ('Chargers', 'chargers@chorzow.pl', '811223344', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Chorzów')),
    ('Cyclones', 'cyclones@cieszyn.pl', '945678123', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Cieszyn')),
    ('Dragons', 'dragons@czechowice-dziedzice.pl', '712345678', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czechowice-Dziedzice')),
    ('Cougars', 'cougars@czeladz.pl', '876543210', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czeladź')),
    ('Crusaders', 'crusaders@czerwionka-leszczyny.pl', '900123456', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Czerwionka-Leszczyny')),
    ('Comets', 'comets@czestochowa.pl', '789123456', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Częstochowa')),
    ('Dolphins', 'dolphins@dabrowa-gornicza.pl', '800234567', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Dąbrowa Górnicza')),
    ('Giants', 'giants@gliwice.pl', '923456789', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Gliwice')),
    ('Ice', 'ice@imielin.pl', '701234567', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Imielin')),
    ('Jaguars', 'jaguars@jastrzebie-zdroj.pl', '998765432', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Jastrzębie-Zdrój')),
    ('Jets', 'jets@jaworzno.pl', '811223344', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Jaworzno')),
    ('Knights', 'knights@kalety.pl', '945678123', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Kalety')),
    ('Kings', 'kings@katowice.pl', '712345678', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Katowice')),
    ('Kangaroos', 'kangaroos@knurow.pl', '876543210', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Knurów')),
    ('Kestrels', 'kestrels@koniecpol.pl', '900123456', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Koniecpol')),
    ('Kodiaks', 'kodiaks@kozieglowy.pl', '789123456', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Koziegłowy')),
    ('Keepers', 'keepers@krzanowice.pl', '800234567', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Krzanowice')),
    ('Kites', 'kites@krzepice.pl', '923456789', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Krzepice')),
    ('Raptors', 'raptors@kuznia-raciborska.pl', '701234567', (SELECT id FROM categories WHERE name = 'U18' AND gender = 'MALE'), (SELECT id FROM cities WHERE name = 'Kuźnia Raciborska'))
    ON CONFLICT DO NOTHING;

COMMIT;