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
 * 本软件是“按原样”提供的，没有任何形式的明示或暗示的保证，包括但不限于
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
CREATE TABLE `cs_class_assignment`
(
    `class_assignment_uuid`           CHAR(32)          NOT NULL PRIMARY KEY COMMENT '排课主键',
    `semester_uuid`                   CHAR(32)          NOT NULL COMMENT '学期主键',
    `course_uuid`                     CHAR(32)          NOT NULL COMMENT '课程主键',
    `teacher_uuid`                    CHAR(32)          NOT NULL COMMENT '教师主键',
    `classroom_uuid`                  CHAR(32)          NOT NULL COMMENT '教室主键',
    `teaching_class_composition_uuid` JSON              NOT NULL COMMENT '教学班组成',
    `course_ownership`                VARCHAR(32)       NOT NULL COMMENT '课程归属',
    `teaching_class_name`             VARCHAR(64)       NOT NULL COMMENT '教学班名称',
    `credit_hour_type`                CHAR(32)          NOT NULL COMMENT '学时类型',
    `teaching_hours`                  DECIMAL(10, 2)    NOT NULL COMMENT '教学学时',
    `scheduled_hours`                 DECIMAL(10, 2)    NOT NULL COMMENT '排课学时',
    `total_hours`                     DECIMAL(10, 2)    NOT NULL COMMENT '总学时',
    `scheduling_priority`             SMALLINT UNSIGNED NOT NULL DEFAULT 100 COMMENT '排课优先级',
    `class_size`                      INT UNSIGNED      NOT NULL COMMENT '班级规模',
    `teaching_campus`                 CHAR(32)          NOT NULL COMMENT '教学校区',
    `class_time`                      JSON              NOT NULL COMMENT '上课时间',
    `consecutive_sessions`            TINYINT UNSIGNED  NOT NULL DEFAULT 2 COMMENT '连堂节数',
    `classroom_type`                  CHAR(32)          NOT NULL COMMENT '教室类型',
    `designated_classroom` CHAR(32) NULL COMMENT '指定教室',
    `designated_teaching_building`    CHAR(32)          NULL COMMENT '指定教学楼',
    `specified_time`                  JSON              NULL COMMENT '指定时间',
    `created_at`                      TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                      TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '排课表';

-- 添加索引以优化查询性能
CREATE INDEX `idx_class_assignment_semester` ON `cs_class_assignment` (`semester_uuid`) COMMENT '学期索引';
CREATE INDEX `idx_class_assignment_teacher` ON `cs_class_assignment` (`teacher_uuid`) COMMENT '教师索引';
CREATE INDEX `idx_class_assignment_classroom` ON `cs_class_assignment` (`classroom_uuid`) COMMENT '教室索引';
CREATE INDEX `idx_class_assignment_course` ON `cs_class_assignment` (`course_uuid`) COMMENT '课程索引';
CREATE INDEX `idx_class_assignment_priority` ON `cs_class_assignment` (`scheduling_priority`) COMMENT '优先级索引';
CREATE INDEX `idx_class_assignment_campus` ON `cs_class_assignment` (`teaching_campus`) COMMENT '校区索引';

ALTER TABLE `cs_class_assignment`
    ADD CONSTRAINT `fk_class_assignment_semester_uuid`
        FOREIGN KEY (`semester_uuid`) REFERENCES `cs_semester` (`semester_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_course_uuid`
        FOREIGN KEY (`course_uuid`) REFERENCES `cs_course_library` (`course_library_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_teacher_uuid`
        FOREIGN KEY (`teacher_uuid`) REFERENCES `cs_teacher` (`teacher_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_classroom_uuid`
        FOREIGN KEY (`classroom_uuid`) REFERENCES `cs_classroom` (`classroom_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_credit_hour_type`
        FOREIGN KEY (`credit_hour_type`) REFERENCES `cs_credit_hour_type` (`credit_hour_type_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_teaching_campus`
        FOREIGN KEY (`teaching_campus`) REFERENCES `cs_campus` (`campus_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_classroom_type`
        FOREIGN KEY (`classroom_type`) REFERENCES `cs_classroom_type` (`class_type_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_designated_classroom`
        FOREIGN KEY (`designated_classroom`) REFERENCES `cs_classroom` (`classroom_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_class_assignment_designated_teaching_building`
        FOREIGN KEY (`designated_teaching_building`) REFERENCES `cs_building` (`building_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE;
