CREATE DATABASE ecommerce;

CREATE TABLE `users`
(
    `id` bigint(20) NOT NULL auto_increment,
    `name` varchar(100) NOT NULL,
    `mobile_number` bigint(20) NOT NULL,
    `longitude` Double precision(10,4) NOT NULL,
    `latitude` Double precision(10,4) NOT NULL,
    `created_at` bigint(20) NOT NULL,
    `updated_at` bigint(20) NOT NULL,
    primary key (`id`)
);

CREATE TABLE `products`
(
    `id` bigint(20) NOT NULL auto_increment,
    `product_name` varchar(100) NOT NULL,
    `product_description` varchar(250) NOT NULL,
    `product_price` bigint(20) NOT NULL,
    `created_at` bigint(20) NOT NULL,
    `updated_at` bigint(20) NOT NULL,
    primary key (`id`)
);

CREATE TABLE `orders`
(
    `id` bigint(20) NOT NULL auto_increment,
    `user_id` bigint(20) NOT NULL,
    `product_id` bigint(20) NOT NULL,
    foreign key (`user_id`) references users(`id`),
    foreign key (`product_id`) references products(`id`),
    primary key (`id`)
);


CREATE TABLE `product_links`
(
    `id` bigint(20) NOT NULL auto_increment,
    `product_id` bigint(20) NOT NULL,
    `link` varchar(2000) NOT NULL,
    foreign key (`product_id`) references products(`id`),
    primary key (`id`)
);

CREATE TABLE `compressed_product_path`
(
    `id` bigint(20) NOT NULL auto_increment,
    `product_id` bigint(20) NOT NULL,
    `path` varchar(200) NOT NULL,
    foreign key (`product_id`) references products(`id`),
    primary key (`id`)
);

