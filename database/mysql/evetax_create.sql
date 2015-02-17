/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping database structure for evetax
CREATE DATABASE IF NOT EXISTS `evetax` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `evetax`;


-- Dumping structure for table evetax.lootpastes
CREATE TABLE IF NOT EXISTS `lootpastes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `charactername` varchar(64) NOT NULL,
  `rawpaste` text NOT NULL,
  `pastecomment` text NOT NULL,
  `totalvalue` bigint(20) NOT NULL,
  `taxamount` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
