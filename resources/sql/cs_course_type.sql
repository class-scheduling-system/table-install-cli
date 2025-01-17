USE `class-scheduling-system`;
CREATE TABLE `cs_course_type`
(
    `course_type_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '课程类型主键',
    `name`            VARCHAR(32) NOT NULL COMMENT '课程类型名称',
    `description`     VARCHAR(255) NULL COMMENT '课程类型描述',
    `created_at`      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '课程类型表';

CREATE UNIQUE INDEX `uk_course_type_name` ON `cs_course_type` (`name`) COMMENT '课程类型名称唯一索引';