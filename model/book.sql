--- CREATE DATABASE `bookstore`;
USE bookstore;
DROP TABLE `books`;

CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT, 
  `isbn` varchar(13) NOT NULL,
  `author` varchar(128),
  `title` varchar(256) NOT NULL,
  `publisher` varchar(128),
  `year` varchar(32) NOT NULL DEFAULT '2000',
  `tags`  varchar(128), 
  `language` varchar(64),
  `rating` tinyint(4),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- books info
--

INSERT INTO `books` (`isbn`, `author`, `title`, `publisher`, `year`, `tags`, `language`, `rating`) VALUES
('9781617291784', 'Brian Ketelsen & Erik St. Martin & William Kennedy', 'Go in Action', 'Manning Publications', '2015', 'Web & Programming', 'english',4),
('9780321573513', 'Robert Sedgewick & Kevin Wayne', 'Algorithms', 'Addison-Wesley Professional', '2011', 'Algorithms & Programming', 'english',5),
('9780321637734', 'W. Richard Stevens & Stephen A. Rago', 'Advanced Programming in the UNIX Environment', 'Addison-Wesley', '2013', 'Unix & Operating Systems', 'english',4);
