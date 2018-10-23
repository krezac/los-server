USE los;

-- cleanup
-- note the reverse order becase of keys
DELETE FROM shots;
DELETE FROM attempts;
DELETE FROM competitors;
DELETE FROM users;
DELETE FROM squads;
DELETE FROM hits;
DELETE FROM hit_type_values;
DELETE FROM hit_values;
DELETE FROM hit_types;
DELETE FROM targets;
DELETE FROM target_types;
DELETE FROM situations;
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

INSERT INTO ranges (ID, NAME, LATITUDE, LONGITUDE, URL, ACTIVE) VALUES
	(4, 'Not Active TEST', 50.091916666700, 14.441044444400, 'https://zbranekvalitne.cz/strelnice/avim-praha', 0);


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
	(1, '10 ran a dost 29', '2018-10-29', 1, 1, 1),
	(2, 'Velka rana', '2018-10-28', 1, 2, 2);

INSERT INTO situations (ID, NUMBER, NAME, COMPETITION_ID) VALUES
	(1, 1, 'Utek ze zajeti', 1),
	(2, 2, 'Dostavnik', 1);

INSERT INTO target_types (ID, NAME, SPECIAL_UI) VALUES
	(1, 'LOS papir 2 rany', "ui_los_2"),
	(2, 'Popper', "ui_popper");

INSERT INTO targets (ID, NUMBER, NAME, SITUATION_ID, TARGET_TYPE_ID) VALUES
	(1, 1, 'Terc 1', 1, 1),
	(2, 2, 'Terc 2', 1, 1),
	(3, 3, 'Terc 3', 1, 2);

INSERT INTO hit_types (ID, NAME, CUMMULATIVE) VALUES
	(1, 'Zasah papir', 0),
	(2, 'Zasah kov', 1);

INSERT INTO hit_values (ID, CODE, NAME, VALUE) VALUES
	(1, 'A', 'Alpha', 0),
	(2, 'C', 'Charlie', 1),
	(3, 'D', 'Delta', 2),
	(4, 'MT', 'Mis terc', 5),
	(5, 'NT', 'Neterc', 10),
	(6, 'PP', 'Popper', 0),
	(7, 'MP', 'Miss popper', 10),
	(8, 'Proc', 'Procedura', 3);

INSERT INTO hit_type_values (ID, NUMBER, HIT_TYPE_ID, HIT_VALUE_ID) VALUES
	(1, 1, 1, 1), -- papir
	(2, 2, 1, 2),
	(3, 3, 1, 3),
	(4, 4, 1, 4),
	(5, 1, 2, 6), -- popper
	(6, 2, 2, 7);

INSERT INTO hits (ID, NUMBER, TARGET_TYPE_ID, HIT_TYPE_ID, COUNT) VALUES 
	(1, 1, 1, 1, 2),
	(3, 1, 2, 2, 1);

INSERT INTO squads(ID, NUMBER, NAME, COMPETITION_ID) VALUES
	(1, 1, 'Squad 1', 1);

INSERT INTO users (ID, LOGIN, PASSWORD, ROLE_COMPETITOR, ROLE_JUDGE, ROLE_DIRECTOR, ROLE_ADMIN) VALUES
	(1, 'admin1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 1, 1, 1),
	(2, 'director1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 1, 1, 0),
	(3, 'judge1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 1, 0, 0),
	(4, 'user1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 0, 0, 0);

INSERT INTO competitors (ID, FIRST_NAME, LAST_NAME, NICKNAME, EMAIL, LICENCE, SQUAD_ID) VALUES
	(1, 'Karel', 'Ctvrty', 'Bouchac', 'a@b.cz', 'AB12345678', 1);

INSERT INTO attempts (ID, COMPETITOR_ID, SITUATION_ID, TIME, JUDGE_ID) VALUES
	(1, 1, 1, 12.34, 2);

INSERT INTO shots (ID, ATTEMPT_ID, TARGET_ID, HIT_VALUE_ID, COUNT) VALUES
	(1, 1, 1, 1, 1);
