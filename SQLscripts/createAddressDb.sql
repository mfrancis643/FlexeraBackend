DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
    `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
    `streetNameAndNumber` varchar(30) NOT NULL,
    `town` varchar(30) NOT NULL,
    `postcode` varchar(8) NOT NULL,
    PRIMARY KEY (`id`)
                       )
ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;