DROP TABLE IF EXISTS `people`;
CREATE TABLE `people` (
    `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(30) NOT NULL,
    `balance` float NOT NULL,
    `email` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
                      )
ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;