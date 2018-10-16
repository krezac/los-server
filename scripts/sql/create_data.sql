USE los;

-- cleanup
-- note the reverse order becase of keys
DELETE FROM users;
DELETE FROM competitions;
DELETE FROM competition_categories;
DELETE FROM competition_types;
DELETE FROM ranges;

-- create

-- source of data: http://zbranekvalitne.cz/strelecka-mapa/data.json (see real data generator in scripts/data)
INSERT INTO ranges (ID, NAME, LATITUDE, LONGITUDE, URL) VALUES
	(1, 'SSK Čelákovice TEST', 50.148425000000, 14.735038888900, 'https://zbranekvalitne.cz/strelnice/ssk-celakovice'),
	(2, 'SSK Žalany TEST', 50.596953611100, 13.893513333300, 'https://zbranekvalitne.cz/strelnice/zalany'),
	(3, 'AVIM Praha TEST', 50.091916666700, 14.441044444400, 'https://zbranekvalitne.cz/strelnice/avim-praha');

INSERT INTO competition_categories (ID, CODE, NAME) VALUES
	(1, 'P', 'Pohárová'),
	(2, 'K', 'Klubová'),
	(3, 'L', 'Losík'),
	(4, 'M', 'Mistrovství ČR');

INSERT INTO competition_types (ID, CODE, NAME) VALUES
	(1, 'KZ', 'Krátká zbraň'),
	(2, 'Pu', 'Puška'),
	(3, 'Br', 'Brokovnice');

INSERT INTO competitions (ID, NAME, EVENT_DATE, RANGE_ID, CATEGORY_ID, TYPE_ID) VALUES
	(1, '10 ran a dost 27', '2018-10-27', 1, 1, 1),
	(2, 'Velka rana', '2018-10-28', 1, 2, 2);

INSERT INTO users (ID, LOGIN, PASSWORD) VALUES
	(1, 'user1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK'),
	(2, 'user2@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK');
