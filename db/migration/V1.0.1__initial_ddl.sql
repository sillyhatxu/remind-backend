CREATE TABLE IF NOT EXISTS `remind`
(
  `id`                 bigint(48)   NOT NULL AUTO_INCREMENT,
  `remind_time`        timestamp(3) NULL     DEFAULT NULL,
  `remind_type`        int          NOT NULL,
  `remind_data`        text         NOT NULL,
  `remind_advance`     int          NOT NULL,
  `remind_frequency`   int          NOT NULL,
  `is_lunar`           TINYINT(1)   not null,
  `status`             int          NOT NULL,
  `created_time`       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `last_modified_time` timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  PRIMARY KEY (`id`),
  INDEX (`remind_type`),
  INDEX (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;