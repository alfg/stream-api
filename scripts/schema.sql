CREATE TABLE `user` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`first_name` TEXT,
	`last_name` TEXT,
	`email` TEXT UNIQUE
);

CREATE TABLE `client` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`api_key` TEXT,
	`api_secret` TEXT,
	`email` TEXT UNIQUE
	`domain` TEXT UNIQUE
);

CREATE TABLE `stream` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`title` TEXT,
	`type` TEXT,
	`description`	TEXT,
  `private` BOOLEAN,
  `stream_name` TEXT UNIQUE,
  `stream_key` TEXT UNIQUE
);
