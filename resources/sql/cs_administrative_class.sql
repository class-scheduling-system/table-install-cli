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
CREATE TABLE `cs_administrative_class`
(
    `administrative_class_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '行政班主键',
    `department_uuid`           CHAR(32)     NOT NULL COMMENT '所属部门/院系',
    `major_uuid`                CHAR(32)     NOT NULL COMMENT '所属专业',
    `class_code`                VARCHAR(32)  NOT NULL COMMENT '班级编号',
    `class_name`                VARCHAR(64)  NOT NULL COMMENT '班级名称',
    `grade_uuid`                CHAR(32)     NOT NULL COMMENT '年级UUID',
    `is_administrative`         BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否为行政班(0:否[虚拟班],1:是[行政班])',
    `student_count`             INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '学生人数',
    `counselor_uuid`            CHAR(32)     NULL COMMENT '辅导员UUID',
    `monitor_uuid`              CHAR(32)     NULL COMMENT '班长UUID',
    `is_enabled`                BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否启用(0:禁用,1:启用)',
    `description`               VARCHAR(255) NULL COMMENT '班级描述',
    `created_at`                TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '行政班表';

CREATE UNIQUE INDEX `uk_class_code` ON `cs_administrative_class` (`class_code`) COMMENT '班级编号唯一索引';
CREATE INDEX `idx_administrative_class_department` ON `cs_administrative_class` (`department_uuid`) COMMENT '院系索引';
CREATE INDEX `idx_administrative_class_major` ON `cs_administrative_class` (`major_uuid`) COMMENT '专业索引';
CREATE INDEX `idx_administrative_class_grade_uuid` ON `cs_administrative_class` (`grade_uuid`) COMMENT '年级索引';
CREATE INDEX `idx_administrative_class_department_major` ON `cs_administrative_class` (`department_uuid`, `major_uuid`) COMMENT '院系专业组合索引';
CREATE INDEX `idx_administrative_class_counselor` ON `cs_administrative_class` (`counselor_uuid`) COMMENT '辅导员索引';
CREATE INDEX `idx_administrative_class_monitor` ON `cs_administrative_class` (`monitor_uuid`) COMMENT '班长索引';
CREATE INDEX `idx_administrative_class_enabled` ON `cs_administrative_class` (`is_enabled`) COMMENT '启用状态索引';

ALTER TABLE `cs_administrative_class`
    ADD CONSTRAINT `fk_administrative_class_department`
        FOREIGN KEY (`department_uuid`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_administrative_class_major`
        FOREIGN KEY (`major_uuid`) REFERENCES `cs_major` (`major_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_administrative_class_counselor`
        FOREIGN KEY (`counselor_uuid`) REFERENCES `cs_teacher` (`teacher_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_administrative_class_monitor`
        FOREIGN KEY (`monitor_uuid`) REFERENCES `cs_student` (`student_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_administrative_class_grade`
        FOREIGN KEY (`grade_uuid`) REFERENCES `cs_grade` (`grade_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE `cs_student`
    ADD CONSTRAINT `fk_cs_administrative_class_cs_student`
        FOREIGN KEY (`class`) REFERENCES `cs_administrative_class` (`administrative_class_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;
