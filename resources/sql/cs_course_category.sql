USE `class-scheduling-system`;
CREATE TABLE `cs_course_category`
(
    `course_category_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '课程分类主键',
    `name`                VARCHAR(32) NOT NULL COMMENT '课程分类名称',
    `description`         VARCHAR(255) NULL COMMENT '课程分类描述',
    `created_at`          TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '课程分类表';

CREATE UNIQUE INDEX `uk_course_category_name` ON `cs_course_category` (`name`) COMMENT '课程分类名称唯一索引';