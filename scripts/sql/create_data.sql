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
	(1, 1, 'Squad 1', 1),
	(2, 2, 'Squad 2', 1),
	(3, 3, 'Squad 3', 1);

INSERT INTO users (ID, LOGIN, PASSWORD, ROLE_COMPETITOR, ROLE_JUDGE, ROLE_DIRECTOR, ROLE_ADMIN) VALUES
	(1, 'admin1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 1, 1, 1),
	(2, 'director1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 1, 1, 0),
	(3, 'judge1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 1, 0, 0),
	(4, 'user1@dev.los', '$2a$10$pF0t.w3y.zrra2tj9j5U1eU9XkVXoze83jRbmJlqJfAqhEeU1SvwK', 1, 0, 0, 0);

INSERT INTO competitors (ID, FIRST_NAME, LAST_NAME, NICKNAME, EMAIL, LICENCE, SQUAD_ID) VALUES
	(1, 'Karel', 'Prazsky', 'Bouchac', 'a@b.cz', 'AB000001', 1),
	(2, 'Karina', 'Brnenska', 'Bouchac', 'a@b.cz', 'AB000002', 1),
	(3, 'Radmila', 'Ostravska', 'Bouchac', 'a@b.cz', 'AB000003', 1),
	(4, 'Diana', 'Plzenska', 'Bouchac', 'a@b.cz', 'AB000004', 1),
	(5, 'Dalimil', 'Ustecky', 'Bouchac', 'a@b.cz', 'AB000005', 1),
	(6, 'Kaspar', 'Mikulovsky', 'Bouchac', 'a@b.cz', 'AB000006', 1),
	(7, 'Vilma', 'Horazdovicka', 'Bouchac', 'a@b.cz', 'AB000007', 1),
	(8, 'Cestmir', 'Susicky', 'Bouchac', 'a@b.cz', 'AB000008', 1),
	(9, 'Vladan', 'Strakonicky', 'Bouchac', 'a@b.cz', 'AB000009', 1),
	(10, 'Bretislav', 'Kadansky', 'Bouchac', 'a@b.cz', 'AB000010', 1),
	(11, 'Hynek', 'Lounsky', 'Bouchac', 'a@b.cz', 'AB0000011', 2),
	(12, 'Nela', 'Slanska', 'Bouchac', 'a@b.cz', 'AB000012', 2),
	(13, 'Blazej', 'Pardubicky', 'Bouchac', 'a@b.cz', 'AB000013', 2),
	(14, 'Jarmila', 'Hradecka', 'Bouchac', 'a@b.cz', 'AB000014', 2),
	(15, 'Dobromila', 'Frydecka', 'Bouchac', 'a@b.cz', 'AB000015', 2),
	(16, 'Vanda', 'Mistecka', 'Bouchac', 'a@b.cz', 'AB000016', 2),
	(17, 'Veronika', 'Olomoucka', 'Bouchac', 'a@b.cz', 'AB000017', 2),
	(18, 'Milada', 'CHrudimska', 'Bouchac', 'a@b.cz', 'AB000020', 2),
	(21, 'Bedrich', 'Verounsky', 'Bouchac', 'a@b.cz', 'AB000021', 3),
	(22, 'Anezka', 'Kladenska', 'Bouchac', 'a@b.cz', 'AB000022', 3),
	(23, 'Kamil', 'Maloborsky', 'Bouchac', 'a@b.cz', 'AB000023', 3),
	(24, 'Stela', 'Brezanska', 'Bouchac', 'a@b.cz', 'AB000024', 3),
	(25, 'Kazimir', 'Vestecky', 'Bouchac', 'a@b.cz', 'AB000025', 3),
	(26, 'Mirek', 'Podebradsky', 'Bouchac', 'a@b.cz', 'AB000026', 3),
	(27, 'Tomas', 'Nymbursky', 'Bouchac', 'a@b.cz', 'AB000027', 3),
	(28, 'Gabriela', 'Brodska', 'Bouchac', 'a@b.cz', 'AB000028', 3),
	(29, 'Frantiska', 'Klatovska', 'Bouchac', 'a@b.cz', 'AB000029', 3),
	(30, 'Viktorie', 'Kacerovkska', 'Bouchac', 'a@b.cz', 'AB0000030', 3),
	(31, 'Andela', 'Rudska', 'Bouchac', 'a@b.cz', 'AB000031', 3),
	(32, 'Rehor', 'Revnicky', 'Bouchac', 'a@b.cz', 'AB000032', 3);

INSERT INTO attempts (ID, COMPETITOR_ID, SITUATION_ID, TIME, JUDGE_ID) VALUES
	(1, 1, 1, 12.34, 2);

INSERT INTO shots (ID, ATTEMPT_ID, TARGET_ID, HIT_VALUE_ID, COUNT) VALUES
	(1, 1, 1, 1, 1);
