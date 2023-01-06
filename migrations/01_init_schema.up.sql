
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(50) NOT NULL,
                         `password` varchar(255) NOT NULL,
                         `last_name` varchar(50) NOT NULL,
                         `first_name` varchar(50) NOT NULL,
                         `phone` varchar(20) DEFAULT NULL,
                         `role` enum('guest', 'host', 'admin') NOT NULL DEFAULT 'guest',
                         `avatar` json DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` timestamp NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `email` (`email`)
);
DROP TABLE IF EXISTS `places`;
CREATE TABLE `places` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `name` VARCHAR(50) NOT NULL,
                         `description` VARCHAR(255) NOT NULL,
                         `total_guests` INT NOT NULL,
                         `total_bedrooms` INT NOT NULL,
                         `total_bathrooms` INT NOT NULL,
                         `price_per_night` DOUBLE NOT NULL,
                         `average_rating` DOUBLE NOT NULL,
                         `owner_id` INT NOT NULL,
                         `location_id` INT NOT NULL,
                         `lat` DOUBLE NULL,
                         `lng` DOUBLE NULL,
                         `address` VARCHAR(255) NOT NULL,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` timestamp NULL,

                         PRIMARY KEY (`id`),
                         KEY `owner_id` (`owner_id`) USING BTREE,
                         KEY `location_id` (`location_id`) USING BTREE

);
DROP TABLE IF EXISTS `locations`;
CREATE TABLE `locations`(
                            `id` INT NOT NULL AUTO_INCREMENT,
                            `country` VARCHAR(255) NOT NULL,
                            `state` VARCHAR(255) NOT NULL,
                            `city` VARCHAR(255) NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` DATETIME NULL,
                            PRIMARY KEY (`id`)
);
DROP TABLE IF EXISTS `amenities`;
CREATE TABLE `amenities`(
                            `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                            `name` VARCHAR(255) NOT NULL,
                            `description` VARCHAR(255) NOT NULL,
                            `icon` json NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` DATETIME NULL,
                            PRIMARY KEY (`id`)
);
DROP TABLE IF EXISTS `place_amenities`;
CREATE TABLE `place_amenities` (
                                  `amenity_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `place_id` INT NOT NULL,
                                  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `deleted_at` DATETIME NULL,
                                  PRIMARY KEY (`amenity_id`,`place_id`)
);

DROP TABLE IF EXISTS `bookings`;
CREATE TABLE `bookings`(
                           `id` INT NOT NULL AUTO_INCREMENT,
                           `user_id` INT NOT NULL,
                           `place_id` INT NOT NULL,
                           `checkin_date` DATETIME NOT NULL,
                           `checkout_date` DATETIME NOT NULL,
                           `discount` DOUBLE(8, 2) NOT NULL,
    `total_price` DOUBLE(8, 2) NOT NULL,
    `status` ENUM('pending', 'reserved', 'deposit', 'paid', 'waiting', 'cancelled') NOT NULL DEFAULT 'pending',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`) USING BTREE,
    KEY `place_id` (`place_id`) USING BTREE

);

DROP TABLE IF EXISTS `reviews`;
CREATE TABLE `reviews`(
                          `id` INT NOT NULL AUTO_INCREMENT,
                          `booking_id` INT NOT NULL,
                          `rating` INT NOT NULL,
                          `comment` VARCHAR(255) NOT NULL,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` DATETIME NULL,
                          PRIMARY KEY (`id`)
);