CREATE database IF NOT EXISTS ecomm ;

use ecomm;

CREATE TABLE IF NOT EXISTS `users` (
    `uid` bigint PRIMARY KEY NOT NULL ,
    `name` varchar(255) NOT NULL ,
    `password` varchar(255) NOT NULL ,
    `email` varchar(255) NOT NULL ,
    `phone` int NOT NULL ,
    `sex` int2 NOT NULL
);

CREATE TABLE IF NOT EXISTS `orders` (
    `id` int PRIMARY KEY  NOT NULL  AUTO_INCREMENT,
    `payment_method` varchar(255) NOT NULL ,
    `shipping_price` decimal(10,2) NOT NULL ,
    `total_price` decimal(10,2) NOT NULL ,
    `created_at` datetime DEFAULT (now()),
    `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS `product` (
    `id` int PRIMARY KEY NOT NULL  AUTO_INCREMENT,
    `uid` bigint NOT NULL ,
    `name` varchar(255) NOT NULL ,
    `image` varchar(255) NOT NULL ,
    `category` varchar(255) NOT NULL ,
    `description` varchar(255) NOT NULL ,
    `price` decimal(10,2) NOT NULL ,
    `count_in_stock` int NOT NULL,
    `created_at` datetime DEFAULT (now()),
    `updated_at` datetime
);

ALTER TABLE `product` ADD FOREIGN KEY (`uid`) REFERENCES `users` (`uid`);

CREATE TABLE IF NOT EXISTS `order_item` (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    order_id int NOT NULL ,
    product_id int NOT NULL ,
    name varchar(255) NOT NULL ,
    image varchar(255) NOT NULL ,
    price int NOT NULL
);

ALTER TABLE `order_item` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);
ALTER TABLE `order_item` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

CREATE TABLE IF NOT EXISTS `good` (
    `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `price` int NOT NULL,
    `created_at` datetime DEFAULT (now()),
    `count` int NOT NULL
);