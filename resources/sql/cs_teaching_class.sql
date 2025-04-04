/*
 * --------------------------------------------------------------------------------
 * Copyright (c) 2022-NOW(至今) 锋楪技术团队
 * Author: 锋楪技术团队 (https://www.frontleaves.com)
 *
 * 本文件包含锋楪技术团队项目的源代码，项目的所有源代码均遵循 MIT 开源许可证协议。
 * --------------------------------------------------------------------------------
 * 许可证声明：
 *
 * 版权所有 (c) 2022-2025 锋楪技术团队。保留所有权利。
 *
 * 本软件是"按原样"提供的，没有任何形式的明示或暗示的保证，包括但不限于
 * 对适销性、特定用途的适用性和非侵权性的暗示保证。在任何情况下，
 * 作者或版权持有人均不承担因软件或软件的使用或其他交易而产生的、
 * 由此引起的或以任何方式与此软件有关的任何索赔、损害或其他责任。
 *
 * 使用本软件即表示您了解此声明并同意其条款。
 *
 * 有关 MIT 许可证的更多信息，请查看项目根目录下的 LICENSE 文件或访问：
 * https://opensource.org/licenses/MIT
 * --------------------------------------------------------------------------------
 * 免责声明：
 *
 * 使用本软件的风险由用户自担。作者或版权持有人在法律允许的最大范围内，
 * 对因使用本软件内容而导致的任何直接或间接的损失不承担任何责任。
 * --------------------------------------------------------------------------------
 */

USE `class-scheduling-system`;
CREATE TABLE `cs_teaching_class`
(
    `teaching_class_uuid`        CHAR(32)     NOT NULL PRIMARY KEY COMMENT '教学班主键',
    `semester_uuid`              CHAR(32)     NOT NULL COMMENT '学期主键',
    `course_uuid`                CHAR(32)     NOT NULL COMMENT '课程主键',
    `teaching_class_code`        VARCHAR(32)  NOT NULL COMMENT '教学班编号',
    `teaching_class_name`        VARCHAR(64)  NOT NULL COMMENT '教学班名称',
    `administrative_classes`     JSON         NOT NULL COMMENT '包含的行政班级(包含班级UUID)',
    `is_administrative`          BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '如果是必修课（区分必修和选修，选修不包含行政班）则该字段为true，否则为false',
    `class_size`                 INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '班级规模',
    `actual_student_count`       INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '实际学生人数',
    `course_department_uuid`     CHAR(32)     NOT NULL COMMENT '开课院系',
    `description`                VARCHAR(255) NULL COMMENT '教学班描述',
    `is_enabled`                 BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否启用(0:禁用,1:启用)',
    `created_at`                 TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                 TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '教学班表';

CREATE UNIQUE INDEX `uk_teaching_class_code` ON `cs_teaching_class` (`teaching_class_code`) COMMENT '教学班编号唯一索引';
CREATE INDEX `idx_teaching_class_semester` ON `cs_teaching_class` (`semester_uuid`) COMMENT '学期索引';
CREATE INDEX `idx_teaching_class_course` ON `cs_teaching_class` (`course_uuid`) COMMENT '课程索引';
CREATE INDEX `idx_teaching_class_department` ON `cs_teaching_class` (`course_department_uuid`) COMMENT '院系索引';
CREATE INDEX `idx_teaching_class_enabled` ON `cs_teaching_class` (`is_enabled`) COMMENT '启用状态索引';

ALTER TABLE `cs_teaching_class`
    ADD CONSTRAINT `fk_teaching_class_semester`
        FOREIGN KEY (`semester_uuid`) REFERENCES `cs_semester` (`semester_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_teaching_class_course`
        FOREIGN KEY (`course_uuid`) REFERENCES `cs_course_library` (`course_library_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_teaching_class_department`
        FOREIGN KEY (`course_department_uuid`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;
