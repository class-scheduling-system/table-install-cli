USE `class-scheduling-system`;
CREATE TABLE `cs_credit_hour_type`
(
    `credit_hour_type_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '学时类型主键',
    `name`                 VARCHAR(32) NOT NULL COMMENT '学时类型名称',
    `description`          VARCHAR(255) NULL COMMENT '学时类型描述',
    `created_at`           TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`           TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '学时类型表';

CREATE UNIQUE INDEX `uk_credit_hour_type_name` ON `cs_credit_hour_type` (`name`) COMMENT '学时类型名称唯一索引';