
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(50) NOT NULL,
                         `password` varchar(50) NOT NULL,
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

DROP TABLE IF EXISTS `properties`;
CREATE TABLE `properties` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `name` varchar(100) NOT NULL,
                              `description` text,
                              `icon` json DEFAULT NULL,
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `deleted_at` timestamp NULL,
                              PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `places`;
CREATE TABLE `places` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `name` varchar(50) NOT NULL,
                          `address` varchar(255) NOT NULL,
                          `owner_id` int NOT NULL,
                          `city_id` int DEFAULT NULL,
                          `lat` double DEFAULT NULL,
                          `lng` double DEFAULT NULL,
                          `cover` json NULL DEFAULT NULL,
                          `price_per_night` double DEFAULT '0',
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` timestamp NULL,
                          PRIMARY KEY (`id`),
                          KEY `owner_id` (`owner_id`) USING BTREE,
                          KEY `city_id` (`city_id`) USING BTREE
);

DROP TABLE IF EXISTS `cities`;
CREATE TABLE `cities` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `title` varchar(100) NOT NULL,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` timestamp NULL,
                          PRIMARY KEY (`id`)
);

INSERT INTO `cities` (`title`)
VALUES
    ('An Giang'),
    ('Bắc Giang'),
    ('Bắc Cạn'),
    ('Bạc Liêu'),
    ('Bắc Ninh'),
    ('Bến Tre'),
    ('Bình Định'),
    ('Bình Dương'),
    ('Bình Phước'),
    ('Bình Thuận'),
    ('Cà Mau'),
    ('Cần Thơ'),
    ('Cao Bằng'),
    ('Đà Nẵng'),
    ('Đắk Lắk'),
    ('Đắk Nông'),
    ('Điện Biên'),
    ('Đồng Nai'),
    ('Đồng Tháp'),
    ('Gia Lai'),
    ('Hà Giang'),
    ('Hà Nam'),
    ('Hà Nội'),
    ('Hà Tĩnh'),
    ('Hải Dương'),
    ('Hải Phòng'),
    ('Hậu Giang'),
    ('Hoà Bình'),
    ('Hưng Yên'),
    ('Khánh Hoà'),
    ('Kiên Giang'),
    ('Kon Tum'),
    ('Lai Châu'),
    ('Lâm Đồng'),
    ('Lạng Sơn'),
    ('Lào Cai'),
    ('Long An'),
    ('Nam Định'),
    ('Nghệ An'),
    ('Ninh Bình'),
    ('Ninh Thuận'),
    ('Phú Thọ'),
    ('Phú Yên'),
    ('Quảng Bình'),
    ('Quảng Namm'),
    ('Quãng Ngãi'),
    ('Quãng Ninh'),
    ('Quãng Trị'),
    ('Sóc Trăng'),
    ('Sơn La'),
    ('Tây Ninh'),
    ('Thái Bình'),
    ('Thái Nguyên'),
    ('Thanh Hoá'),
    ('Huế'),
    ('Tiền Giang'),
    ('Hồ Chí Minh'),
    ('Trà Vinh'),
    ('Tuyên Quang'),
    ('Vĩnh Long'),
    ('Vĩnh Phúc'),
    ('Vũng Tàu'),
    ('Yên Bái');