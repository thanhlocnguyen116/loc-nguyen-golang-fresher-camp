/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE `categories` (
  `cate_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`cate_id`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8;

CREATE TABLE `foods` (
  `food_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `cat_id` bigint(20) NOT NULL,
  `price` float NOT NULL,
  `image` json NOT NULL,
  `has_liked` tinyint(1) DEFAULT NULL,
  `rating` float DEFAULT NULL,
  `rating_count` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`food_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE `restaurants` (
  `restaurant_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `cat_id` bigint(20) NOT NULL,
  `ship_time` float NOT NULL,
  `free_ship` tinyint(1) NOT NULL,
  `has_liked` tinyint(1) DEFAULT NULL,
  `rating` float DEFAULT NULL,
  `rating_count` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`restaurant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone` int(11) NOT NULL,
  `address` varchar(255) NOT NULL,
  `image` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `categories` (`cate_id`, `name`) VALUES
(1, 'Chicken');
INSERT INTO `categories` (`cate_id`, `name`) VALUES
(2, 'Burger');


INSERT INTO `foods` (`food_id`, `name`, `cat_id`, `price`, `image`, `has_liked`, `rating`, `rating_count`, `created_at`, `updated_at`) VALUES
(1, 'Chiken Hawaiian', 1, 10.35, '{\"image\": \"none\"}', 0, 4.5, 10, '2022-01-23 04:36:17', '2022-01-23 04:36:17');


INSERT INTO `restaurants` (`restaurant_id`, `name`, `address`, `cat_id`, `ship_time`, `free_ship`, `has_liked`, `rating`, `rating_count`, `created_at`, `updated_at`) VALUES
(1, 'McDonald\'s', '1234 New York', 1, 14, 1, 1, 4.5, 69, '2022-01-23 04:30:39', '2022-01-23 04:30:39');


INSERT INTO `users` (`user_id`, `name`, `password`, `email`, `phone`, `address`, `image`, `created_at`, `updated_at`) VALUES
(1, 'Loc Nguyen', '123456', 'loc.nguyen@gmail.com', 12345689, 'Mễ Trì - Hà Nội', NULL, '2022-01-23 04:38:30', '2022-01-23 04:38:52');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;