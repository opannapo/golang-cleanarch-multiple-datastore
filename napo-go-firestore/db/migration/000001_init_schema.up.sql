# CLI script migration golang-migrate
# migrate -path /media/WinLin/GOLANG-PROJECT/GIT/golang-cleanarch-multiple-datastore/napo-go-firestore/db/migration -database "mysql://napouser:napouser@tcp(localhost:3306)/napo-firestore?multiStatements=true" -verbose up

CREATE TABLE IF NOT EXISTS `credential`
(
    `id`        int(11)      NOT NULL AUTO_INCREMENT,
    `key`       varchar(100) NOT NULL,
    `signature` varchar(100) DEFAULT NULL,
    `type`      int(11)      DEFAULT NULL COMMENT '1 username, 2 email, 3 phone',
    `user_id`   int(11)      DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `credential_UN` (`key`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS `topic`
(
    `id`            int(11) NOT NULL AUTO_INCREMENT,
    `title`         varchar(100) DEFAULT NULL,
    `content`       text         DEFAULT NULL,
    `source`        varchar(100) DEFAULT NULL,
    `topic_type_id` int(11)      DEFAULT NULL COMMENT 'topic_type id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS `topic_type`
(
    `id`    int(11)      NOT NULL AUTO_INCREMENT,
    `label` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;



CREATE TABLE IF NOT EXISTS `topic_type_ref`
(
    `id`            int(11) NOT NULL AUTO_INCREMENT,
    `topic_id`      int(11) DEFAULT NULL COMMENT 'Topic Id',
    `topic_type_id` int(11) DEFAULT NULL COMMENT 'Topic type id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS `user`
(
    `id`              int(11) NOT NULL AUTO_INCREMENT,
    `first_name`      varchar(100) DEFAULT NULL,
    `last_name`       varchar(100) DEFAULT NULL,
    `profile_picture` text         DEFAULT NULL,
    `birth_date`      int(11)      DEFAULT NULL,
    `phone`           varchar(100) DEFAULT NULL,
    `email`           varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS `user_following_topic`
(
    `id`            int(11) NOT NULL AUTO_INCREMENT,
    `user_id`       int(11) DEFAULT NULL,
    `topic_type_id` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX uft_user_id_idx (user_id),
    INDEX uft_topic_type_id_idx (topic_type_id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;


# CREATE INDEX user_following_topic_idx_user_id USING HASH ON user_following_topic (user_id);