DROP TABLE IF EXISTS `roles`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `roles` (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR() NOT NULL
);

INSERT INTO `roles` VALUES (1, "Admin"), (2, "Member"), (3, "Guest");

CREATE TABLE `users` (
  id INT PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(75) NOT NULL UNIQUE,
  username VARCHAR(100) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  token TEXT,
  role_id NOT NULL DEFAULT 3,
  FOREIGN KEY (role_id) REFERENCES roles(id)
);
