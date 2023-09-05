CREATE DATABASE ecommerce;

CREATE TABLE `users`
(
    `id` bigint(20) NOT NULL auto_increment,
    `name` varchar(100) NOT NULL,
    `mobile_number` bigint(20) NOT NULL,
    `longitude` Double precision(10,4) NOT NULL,
    `latitude` Double precision(10,4) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `updated_at` timestamp,
    PRIMARY KEY (`id`)
);

