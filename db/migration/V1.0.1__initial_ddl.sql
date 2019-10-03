CREATE TABLE IF NOT EXISTS `remind`
(
  `id`                 bigint(48)   NOT NULL AUTO_INCREMENT,
  `title`              varchar(200) NOT NULL,
  `description`        TEXT         NOT NULL,
  `type`               int          NOT NULL,
  `status`             int          NOT NULL,
  `remind_time`        timestamp(3)          DEFAULT NULL,
  `remind_type`        int          NOT NULL,
  `created_time`       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `last_modified_time` timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  PRIMARY KEY (`id`),
  INDEX (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;