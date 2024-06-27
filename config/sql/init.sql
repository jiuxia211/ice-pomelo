CREATE TABLE tiny_id (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `biz_type` bigint NOT NULL,
    `max_id` bigint NOT NULL,
    `step` bigint NOT NULL,
    `version` bigint NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
	`updated_at` datetime(3) DEFAULT NULL,
	`deleted_at` datetime(3) DEFAULT NULL
)ENGINE = InnoDB AUTO_INCREMENT = 1000 CHARSET = utf8mb4;

CREATE TABLE user (
    id bigint PRIMARY KEY,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    avatar_url varchar(255),
    `created_at` datetime(3) DEFAULT NULL,
	`updated_at` datetime(3) DEFAULT NULL,
	`deleted_at` datetime(3) DEFAULT NULL
)ENGINE = InnoDB CHARSET = utf8mb4;

CREATE TABLE video (
    id bigint PRIMARY KEY,
    user_id bigint NOT NULL,
    video_url varchar(255) NOT NULL,
    cover_url varchar(255),
    title varchar(255) NOT NULL,
    introduction TEXT,
    `created_at` datetime(3) DEFAULT NULL,
	`updated_at` datetime(3) DEFAULT NULL,
	`deleted_at` datetime(3) DEFAULT NULL
)ENGINE = InnoDB CHARSET = utf8mb4;
