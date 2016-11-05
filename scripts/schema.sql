CREATE TABLE `user` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`first_name` TEXT,
	`last_name` TEXT,
	`email` TEXT UNIQUE
);

CREATE TABLE `stream` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`stream_name` TEXT,
	`type` TEXT,
	`description`	TEXT,
  `url` TEXT UNIQUE,
  `key` TEXT,
  `private` BOOLEAN
);
