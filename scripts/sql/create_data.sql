USE los;

-- cleanup
-- note the reverse order becase of keys
DELETE FROM users;
DELETE FROM competitions;
DELETE FROM competition_types;
DELETE FROM ranges;

-- create

-- source of data: http://zbranekvalitne.cz/strelecka-mapa/data.json
INSERT INTO ranges (ID, NAME, LATITUDE, LONGITUDE, ACTIVE) VALUES
	(1, 'SSK Čelákovice', 50.148425000000, 14.735038888900, 1),
	(2, 'SSK Žalany', 50.596953611100, 13.893513333300, 1),
	(3, 'AVIM Praha', 50.091916666700, 14.441044444400, 1);

INSERT INTO competition_types (ID, CODE, NAME) VALUES
	(1, 'P', 'Pohárová'),
	(2, 'K', 'Klubová'),
	(3, 'L', 'Losík'),
	(4, 'M', 'Mistrovství ČR');

INSERT INTO competitions (ID, NAME, EVENT_DATE, RANGE_ID, TYPE_ID) VALUES
	(1, '10 ran a dost 27', '2018-10-27', 1, 2);

INSERT INTO users (ID, LOGIN, PASSWORD) VALUES
	(1, 'user1', 'pass1');
