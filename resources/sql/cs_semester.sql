USE `class-scheduling-system`;
CREATE TABLE `cs_semester`
(
    `semester_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '学期主键',
    `name`          VARCHAR(32) NOT NULL COMMENT '学期名称',
    `description`   VARCHAR(255) NULL COMMENT '学期描述',
    `start_date`    DATE        NOT NULL COMMENT '学期开始日期',
    `end_date`      DATE        NOT NULL COMMENT '学期结束日期',
    `created_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '学期表';

CREATE UNIQUE INDEX `uk_semester_name` ON `cs_semester` (`name`) COMMENT '学期名称唯一索引';