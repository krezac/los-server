USE los;

-- cleanup
-- note the reverse order becase of keys
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS competitions;
DROP TABLE IF EXISTS competition_types;
DROP TABLE IF EXISTS ranges;

-- create

CREATE TABLE IF NOT EXISTS ranges (
  ID int(11) NOT NULL AUTO_INCREMENT,
  NAME tinytext NOT NULL,
  LATITUDE decimal(15,12) NOT NULL,
  LONGITUDE decimal(15,12) NOT NULL,
  ACTIVE tinyint(1) NOT NULL DEFAULT 1,
  CREATED_TS timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (ID)
) DEFAULT CHARSET=utf8 WITH SYSTEM VERSIONING;

CREATE TABLE IF NOT EXISTS competition_types (
  ID int(11) NOT NULL AUTO_INCREMENT,
  CODE varchar(10) NOT NULL,
  NAME tinytext NOT NULL,
  CREATED_TS timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (ID),
  UNIQUE KEY UQ_COMPETITION_TYPE_CODE (CODE)
) DEFAULT CHARSET=utf8 WITH SYSTEM VERSIONING;

CREATE TABLE IF NOT EXISTS competitions (
  ID int(11) NOT NULL AUTO_INCREMENT,
  NAME tinytext NOT NULL,
  EVENT_DATE date NOT NULL,
  RANGE_ID int(11) NOT NULL,
  TYPE_ID int(11) NOT NULL,
  CREATED_TS timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (ID),
  KEY FK_COMPETITIONS_ranges (RANGE_ID),
  KEY FK_competitions_competition_types (TYPE_ID),
  CONSTRAINT FK_COMPETITIONS_ranges FOREIGN KEY (RANGE_ID) REFERENCES ranges (ID),
  CONSTRAINT FK_competitions_competition_types FOREIGN KEY (TYPE_ID) REFERENCES competition_types (ID)
) DEFAULT CHARSET=utf8 WITH SYSTEM VERSIONING;

CREATE TABLE IF NOT EXISTS users (
  ID int(11) NOT NULL AUTO_INCREMENT,
  LOGIN tinytext NOT NULL,
  PASSWORD tinytext NOT NULL,
  ACTIVE tinyint(1) NOT NULL DEFAULT 1,
  CREATED_TS timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (ID)
)  DEFAULT CHARSET=utf8 WITH SYSTEM VERSIONING;