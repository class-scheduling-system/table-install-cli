USE `class-scheduling-system`;
CREATE TABLE `cs_major`
(
    `major_uuid`        CHAR(32)     NOT NULL PRIMARY KEY COMMENT '专业主键',
    `major_name`        VARCHAR(32)  NOT NULL COMMENT '专业名称',
    `major_description` VARCHAR(255) NULL COMMENT '专业描述',
    `major_code`        VARCHAR(32)  NOT NULL COMMENT '专业代码',
    `major_status`      BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '专业状态 0:禁用 1:启用',
    `created_at`        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '专业表';

CREATE UNIQUE INDEX `uk_major_name` ON `cs_major` (`major_name`) COMMENT '专业名称唯一索引';
CREATE UNIQUE INDEX `uk_major_code` ON `cs_major` (`major_code`) COMMENT '专业代码唯一索引';