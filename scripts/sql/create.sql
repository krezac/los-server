-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.3.10-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- ******** DONE ******** -- Dumping database structure for los
-- ******** DONE ********  DROP DATABASE IF EXISTS `los`;
-- ******** DONE ******** CREATE DATABASE IF NOT EXISTS `los` /*!40100 DEFAULT CHARACTER SET utf8 */;
-- ******** DONE ******** USE `los`;

-- Dumping structure for table los.competitions
-- ******** DONE ******** DROP TABLE IF EXISTS `competitions`;
-- ******** DONE ******** CREATE TABLE IF NOT EXISTS `competitions` (
-- ******** DONE ********   `ID` int(11) NOT NULL AUTO_INCREMENT,
-- ******** DONE ********   `NAME` tinytext NOT NULL,
-- ******** DONE ********   `DATE` date NOT NULL,
-- ******** DONE ********   `RANGE_ID` int(11) NOT NULL,
-- ******** DONE ********   `TYPE_ID` int(11) NOT NULL,
-- ******** DONE ********   PRIMARY KEY (`ID`),
-- ******** DONE ********   KEY `FK_COMPETITIONS_ranges` (`RANGE_ID`),
-- ******** DONE ********   KEY `FK_competitions_competition_types` (`TYPE_ID`),
-- ******** DONE ********   CONSTRAINT `FK_COMPETITIONS_ranges` FOREIGN KEY (`RANGE_ID`) REFERENCES `ranges` (`ID`),
-- ******** DONE ********   CONSTRAINT `FK_competitions_competition_types` FOREIGN KEY (`TYPE_ID`) REFERENCES `competition_types` (`ID`)
-- ******** DONE ******** ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ******** DONE ******** -- Dumping data for table los.competitions: ~0 rows (approximately)
-- ******** DONE ******** DELETE FROM `competitions`;
-- ******** DONE ******** /*!40000 ALTER TABLE `competitions` DISABLE KEYS */;
-- ******** DONE ******** /*!40000 ALTER TABLE `competitions` ENABLE KEYS */;

-- ******** DONE ******** -- Dumping structure for table los.competition_types
-- ******** DONE ******** DROP TABLE IF EXISTS `competition_types`;
-- ******** DONE ******** CREATE TABLE IF NOT EXISTS `competition_types` (
-- ******** DONE ********   `ID` int(11) NOT NULL AUTO_INCREMENT,
-- ******** DONE ********   `CODE` varchar(10) NOT NULL,
-- ******** DONE ********   `NAME` tinytext NOT NULL,
-- ******** DONE ********   PRIMARY KEY (`ID`),
-- ******** DONE ********   UNIQUE KEY `UQ_COMPETITION_TYPE_CODE` (`CODE`)
-- ******** DONE ******** ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- Dumping data for table los.competition_types: ~3 rows (approximately)
-- ******** DONE ******** DELETE FROM `competition_types`;
-- ******** DONE ******** /*!40000 ALTER TABLE `competition_types` DISABLE KEYS */;
-- ******** DONE ******** INSERT INTO `competition_types` (`ID`, `CODE`, `NAME`) VALUES
-- ******** DONE ******** 	(1, 'K', 'Klubová'),
-- ******** DONE ******** 	(2, 'P', 'Pohárová'),
-- ******** DONE ******** 	(3, 'M', 'Mistrovství ČR');
-- ******** DONE ******** /*!40000 ALTER TABLE `competition_types` ENABLE KEYS */;

-- Dumping structure for table los.competitors
DROP TABLE IF EXISTS `competitors`;
CREATE TABLE IF NOT EXISTS `competitors` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `PREFIX` tinytext NOT NULL DEFAULT '',
  `FIRST_NAME` tinytext NOT NULL,
  `LAST_NAME` tinytext NOT NULL,
  `SUFFIX` tinytext NOT NULL DEFAULT '',
  `NICK_NAME` tinytext NOT NULL DEFAULT '',
  `EMAIL` tinytext NOT NULL DEFAULT '',
  `PHONE` tinytext NOT NULL DEFAULT '',
  `LICENCE` tinytext NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table los.competitors: ~0 rows (approximately)
DELETE FROM `competitors`;
/*!40000 ALTER TABLE `competitors` DISABLE KEYS */;
/*!40000 ALTER TABLE `competitors` ENABLE KEYS */;

-- Dumping structure for table los.divisions
DROP TABLE IF EXISTS `divisions`;
CREATE TABLE IF NOT EXISTS `divisions` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `CODE` varchar(10) NOT NULL,
  `NAME` tinytext NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `UQ_DIVISION_CODE` (`CODE`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- Dumping data for table los.divisions: ~4 rows (approximately)
DELETE FROM `divisions`;
/*!40000 ALTER TABLE `divisions` DISABLE KEYS */;
INSERT INTO `divisions` (`ID`, `CODE`, `NAME`) VALUES
	(1, 'Pi', 'Pistole'),
	(2, 'MPi', 'Malá pistole'),
	(3, 'Re', 'Revolver'),
	(4, 'MRe', 'Malý revolver');
/*!40000 ALTER TABLE `divisions` ENABLE KEYS */;

-- Dumping structure for table los.participants
DROP TABLE IF EXISTS `participants`;
CREATE TABLE IF NOT EXISTS `participants` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `COMPETITOR_ID` int(11) NOT NULL DEFAULT 0,
  `SQUAD_ID` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`ID`),
  KEY `FK_participants_squads` (`SQUAD_ID`),
  KEY `FK_participants_competitors` (`COMPETITOR_ID`),
  CONSTRAINT `FK_participants_competitors` FOREIGN KEY (`COMPETITOR_ID`) REFERENCES `competitors` (`ID`),
  CONSTRAINT `FK_participants_squads` FOREIGN KEY (`SQUAD_ID`) REFERENCES `squads` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table los.participants: ~0 rows (approximately)
DELETE FROM `participants`;
/*!40000 ALTER TABLE `participants` DISABLE KEYS */;
/*!40000 ALTER TABLE `participants` ENABLE KEYS */;

-- ******** DONE ******** -- Dumping structure for table los.ranges
-- ******** DONE ******** DROP TABLE IF EXISTS `ranges`;
-- ******** DONE ******** CREATE TABLE IF NOT EXISTS `ranges` (
-- ******** DONE ******** `ID` int(11) NOT NULL AUTO_INCREMENT,
-- ******** DONE ******** `NAME` tinytext NOT NULL,
-- ******** DONE ******** `LATITUDE` decimal(10,7) NOT NULL,
-- ******** DONE ******** `LONGITUDE` decimal(10,7) NOT NULL,
-- ******** DONE ******** `ACTIVE` tinyint(1) NOT NULL DEFAULT 1,
-- ******** DONE ******** PRIMARY KEY (`ID`)
-- ******** DONE ******** ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ******** DONE ******** -- Dumping data for table los.ranges: ~2 rows (approximately)
-- ******** DONE ******** DELETE FROM `ranges`;
-- ******** DONE ******** /*!40000 ALTER TABLE `ranges` DISABLE KEYS */;
-- ******** DONE ******** INSERT INTO `ranges` (`ID`, `NAME`, `LATITUDE`, `LONGITUDE`, `ACTIVE`) VALUES
-- ******** DONE ********	(1, 'Čelákovice', 50.1488390, 14.7349610, 1),
-- ******** DONE ********	(2, 'Žalany', 50.5967990, 13.8933250, 1);
-- ******** DONE ******** /*!40000 ALTER TABLE `ranges` ENABLE KEYS */;

-- Dumping structure for table los.squads
DROP TABLE IF EXISTS `squads`;
CREATE TABLE IF NOT EXISTS `squads` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `COMPETITION_ID` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK_squads_competitions` (`COMPETITION_ID`),
  CONSTRAINT `FK_squads_competitions` FOREIGN KEY (`COMPETITION_ID`) REFERENCES `competitions` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table los.squads: ~0 rows (approximately)
DELETE FROM `squads`;
/*!40000 ALTER TABLE `squads` DISABLE KEYS */;
/*!40000 ALTER TABLE `squads` ENABLE KEYS */;

-- Dumping structure for table los.stages
DROP TABLE IF EXISTS `stages`;
CREATE TABLE IF NOT EXISTS `stages` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NAME` tinytext NOT NULL,
  `COMPETITION_ID` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK__competitions` (`COMPETITION_ID`),
  CONSTRAINT `FK__competitions` FOREIGN KEY (`COMPETITION_ID`) REFERENCES `competitions` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table los.stages: ~0 rows (approximately)
DELETE FROM `stages`;
/*!40000 ALTER TABLE `stages` DISABLE KEYS */;
/*!40000 ALTER TABLE `stages` ENABLE KEYS */;

-- Dumping structure for table los.targets
DROP TABLE IF EXISTS `targets`;
CREATE TABLE IF NOT EXISTS `targets` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `STAGE_ID` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK_targets_stages` (`STAGE_ID`),
  CONSTRAINT `FK_targets_stages` FOREIGN KEY (`STAGE_ID`) REFERENCES `stages` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table los.targets: ~0 rows (approximately)
DELETE FROM `targets`;
/*!40000 ALTER TABLE `targets` DISABLE KEYS */;
/*!40000 ALTER TABLE `targets` ENABLE KEYS */;

-- Dumping structure for table los.users
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table los.users: ~0 rows (approximately)
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
