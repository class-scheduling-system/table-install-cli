USE `class-scheduling-system`;
CREATE TABLE `cs_campus`
(
    `campus_uuid`    CHAR(32)     NOT NULL PRIMARY KEY COMMENT '校区主键',
    `campus_name`    VARCHAR(32)  NOT NULL COMMENT '校区名称',
    `campus_code`    VARCHAR(32)  NOT NULL COMMENT '校区编码',
    `campus_desc`    VARCHAR(255) NULL COMMENT '校区描述',
    `campus_status`  BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '校区状态 0:禁用 1:启用',
    `campus_address` VARCHAR(255) NOT NULL COMMENT '校区地址',
    `created_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '校区表';

CREATE UNIQUE INDEX `uk_campus_name` ON `cs_campus` (`campus_name`) COMMENT '校区名称唯一索引';
CREATE UNIQUE INDEX `uk_campus_code` ON `cs_campus` (`campus_code`) COMMENT '校区编码唯一索引';