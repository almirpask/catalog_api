CREATE DATABASE `catalog`;

use `catalog`;

CREATE TABLE `categories` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `catalog`.`categories` (`id`,`name`) VALUES ("6b4c28f4-6831-495a-9444-19c93452faa3","Relogios");
INSERT INTO `catalog`.`categories` (`id`,`name`) VALUES ("7c0ca0d4-ff23-4bd7-b131-c563067c4b43","Eletronicos");

CREATE TABLE `products` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `category_id` varchar(36) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_categories_idx` (`category_id`)
);