USE `class-scheduling-system`;
CREATE TABLE `cs_course_library`
(
    `course_library_uuid` CHAR(32)       NOT NULL PRIMARY KEY COMMENT '课程库主键',
    `id`                  VARCHAR(32)    NOT NULL COMMENT '课程编号',
    `name`                VARCHAR(32)    NOT NULL COMMENT '课程库名称',
    `english_name`        VARCHAR(128)   NOT NULL COMMENT '课程英文名称',
    `category`            CHAR(32)       NOT NULL COMMENT '课程类别',
    `property`            CHAR(32)       NOT NULL COMMENT '课程属性',
    `type`                CHAR(32)       NOT NULL COMMENT '课程类型',
    `nature`              CHAR(32)       NOT NULL COMMENT '课程性质',
    `department`          CHAR(32)       NOT NULL COMMENT '开课学院',
    `is_enabled`          BOOLEAN        NOT NULL DEFAULT TRUE COMMENT '是否启用',
    `total_hours`         DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '总学时',
    `week_hours`          DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '周学时',
    `theory_hours`        DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '理论学时',
    `experiment_hours`    DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '实验学时',
    `practice_hours`      DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '实践学时',
    `computer_hours`      DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '上机学时',
    `other_hours`         DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '其他学时',
    `credit`              DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '学分',
    `description`         TEXT           NULL COMMENT '课程库描述',
    `edit_user`           CHAR(32)       NULL COMMENT '编辑人',
    `created_at`          TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '课程库表';

CREATE UNIQUE INDEX `uk_course_library_id` ON `cs_course_library` (`id`) COMMENT '课程编号唯一索引';
CREATE UNIQUE INDEX `uk_course_library_name` ON `cs_course_library` (`name`) COMMENT '课程库名称唯一索引';

ALTER TABLE `cs_course_library`
    ADD CONSTRAINT `fk_course_library_category`
        FOREIGN KEY (`category`) REFERENCES `cs_course_category` (`course_category_uuid`),
    ADD CONSTRAINT `fk_course_library_property`
        FOREIGN KEY (`property`) REFERENCES `cs_course_property` (`course_property_uuid`),
    ADD CONSTRAINT `fk_course_library_type`
        FOREIGN KEY (`type`) REFERENCES `cs_course_type` (`course_type_uuid`),
    ADD CONSTRAINT `fk_course_library_nature`
        FOREIGN KEY (`nature`) REFERENCES `cs_course_nature` (`course_nature_uuid`),
    ADD CONSTRAINT `fk_course_library_department`
        FOREIGN KEY (`department`) REFERENCES `cs_department` (`department_uuid`),
    ADD CONSTRAINT `fk_course_library_edit_user`
        FOREIGN KEY (`edit_user`) REFERENCES `cs_user` (`user_uuid`);