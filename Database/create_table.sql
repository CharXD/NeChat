SET NAMES utf8;

CREATE TABLE IF NOT EXISTS `messages` (
  `sender` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `recipient` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content` varchar(1024) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `date` int(13) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `user_infos` (
  `uid` int(6) NOT NULL AUTO_INCREMENT,
  `username` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `salt` varchar(6) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

