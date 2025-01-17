USE `class-scheduling-system`;
CREATE TABLE `cs_classroom_tag`
(
    `class_tag_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '教室标签主键',
    `name`           VARCHAR(32)  NOT NULL COMMENT '教室标签名称',
    `description`    VARCHAR(255) NULL COMMENT '教室标签描述',
    `created_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '教室标签表';

CREATE UNIQUE INDEX `uk_classroom_tag_name` ON `cs_classroom_tag` (`name`) COMMENT '教室标签名称唯一索引';