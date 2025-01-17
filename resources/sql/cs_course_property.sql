USE `class-scheduling-system`;
CREATE TABLE `cs_course_property`
(
    `course_property_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '课程属性主键',
    `name`                VARCHAR(32) NOT NULL COMMENT '课程属性名称',
    `description`         VARCHAR(255) NULL COMMENT '课程属性描述',
    `created_at`          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '课程属性表';

CREATE UNIQUE INDEX `uk_course_property_name` ON `cs_course_property` (`name`) COMMENT '课程属性名称唯一索引';