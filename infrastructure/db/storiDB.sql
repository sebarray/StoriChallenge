-- --------------------------------------------------------
-- Host:                         us-cdbr-east-06.cleardb.net
-- Versión del servidor:         5.6.50-log - MySQL Community Server (GPL)
-- SO del servidor:              Linux
-- HeidiSQL Versión:             12.5.0.6677
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Volcando estructura de base de datos para heroku_97c81d6e787e777
CREATE DATABASE IF NOT EXISTS `heroku_97c81d6e787e777` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `heroku_97c81d6e787e777`;

-- Volcando estructura para tabla heroku_97c81d6e787e777.transactions
CREATE TABLE IF NOT EXISTS `transactions` (
  `ID` int(11) NOT NULL DEFAULT '0',
  `AMOUNT` float NOT NULL,
  `DATE` date NOT NULL,
  `USER_ID` varchar(40) NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  KEY `FK_transactions_user` (`USER_ID`),
  CONSTRAINT `FK_transactions_user` FOREIGN KEY (`USER_ID`) REFERENCES `user` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Volcando datos para la tabla heroku_97c81d6e787e777.transactions: ~5 rows (aproximadamente)
INSERT INTO `transactions` (`ID`, `AMOUNT`, `DATE`, `USER_ID`) VALUES
	(1, 40, '2023-07-30', 'dc2e420c-2f38-11ee-8124-0242ac1101b5'),
	(2, -30, '2023-07-30', 'dc2e420c-2f38-11ee-8124-0242ac1101b5'),
	(3, 60, '2023-06-30', 'dc2e420c-2f38-11ee-8124-0242ac1101b5'),
	(4, 50, '2023-05-30', 'dc2e420c-2f38-11ee-8124-0242ac1101b4'),
	(5, -79, '2023-07-30', 'dc2e420c-2f38-11ee-8124-0242ac1101b4');

-- Volcando estructura para tabla heroku_97c81d6e787e777.user
CREATE TABLE IF NOT EXISTS `user` (
  `ID` varchar(40) NOT NULL,
  `EMAIL` varchar(120) NOT NULL,
  `NAME` varchar(80) NOT NULL,
  `LAST_NAME` varchar(80) NOT NULL,
  `PASSWORD` varchar(300) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Volcando datos para la tabla heroku_97c81d6e787e777.user: ~2 rows (aproximadamente)
INSERT INTO `user` (`ID`, `EMAIL`, `NAME`, `LAST_NAME`, `PASSWORD`) VALUES
	('dc2e420c-2f38-11ee-8124-0242ac1101b4', 'sebarray98@gmail.com', 'sebastian', 'de mello', 'DASDDDDDDDDDDDDDDDDDDDDA'),
	('dc2e420c-2f38-11ee-8124-0242ac1101b5', 's.wiselink@gmail.com', 'wiselink', 'w', 'DDDDASSSSSSSDDDDDDDDDDDDDDDD');

-- Volcando estructura para procedimiento heroku_97c81d6e787e777.get_txn
DELIMITER //
CREATE PROCEDURE `get_txn`(
	IN `email_param` VARCHAR(80)
)
BEGIN




  IF email_param = "" THEN
    SELECT
      us.EMAIL,
      CONCAT(
        '[',
        GROUP_CONCAT(
          CONCAT(
            '{"id": ', txn.ID, 
            ',"date": "', txn.DATE, '",',
            '"amount": ', txn.AMOUNT, 
            '}'
          )
          ORDER BY txn.DATE
          SEPARATOR ','
        ),
        ']'
      ) AS txn_data,
 AVG(CASE WHEN txn.AMOUNT < 0 THEN txn.AMOUNT END) AS average_negative_amount,
  AVG(CASE WHEN txn.AMOUNT >= 0 THEN txn.AMOUNT END) AS average_positive_amount,
  SUM(txn.AMOUNT) AS total_amount,
  us.NAME
    FROM
      heroku_97c81d6e787e777.user AS us
    JOIN
      heroku_97c81d6e787e777.transactions AS txn
    ON
      txn.USER_ID = us.ID
    GROUP BY
      us.EMAIL
    ORDER BY
      us.EMAIL;
      
 ELSEIF email_param!= "" then 
  
  
  
    SELECT
      us.EMAIL,
      CONCAT(
        '[',
        GROUP_CONCAT(
          CONCAT(
            '{"id": ', txn.ID, 
            ',"date": "', txn.DATE, '",',
            '"amount": ', txn.AMOUNT, 
            '}'
          )
          ORDER BY txn.DATE
          SEPARATOR ','
        ),
        ']'
      ) AS txn_data
      ,
 AVG(CASE WHEN txn.AMOUNT < 0 THEN txn.AMOUNT END) AS average_negative_amount,
  AVG(CASE WHEN txn.AMOUNT >= 0 THEN txn.AMOUNT END) AS average_positive_amount,
  SUM(txn.AMOUNT) AS total_amount,
  us.NAME
    FROM
      heroku_97c81d6e787e777.user AS us
    JOIN
      heroku_97c81d6e787e777.transactions AS txn
    ON
      txn.USER_ID = us.ID
    WHERE
      us.EMAIL = email_param
    GROUP BY
      us.EMAIL
    ORDER BY
      us.EMAIL;
  END IF;


END//
DELIMITER ;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
