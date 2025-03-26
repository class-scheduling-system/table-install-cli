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

USE `class-scheduling-system`;
CREATE TABLE `cs_scheduling_conflict`
(
    `conflict_uuid`            CHAR(32)        NOT NULL PRIMARY KEY COMMENT '冲突主键',
    `semester_uuid`            CHAR(32)        NOT NULL COMMENT '学期主键',
    `first_assignment_uuid`    CHAR(32)        NOT NULL COMMENT '第一个排课主键',
    `second_assignment_uuid`   CHAR(32)        NOT NULL COMMENT '第二个排课主键',
    `conflict_type`            TINYINT UNSIGNED NOT NULL COMMENT '冲突类型: 1-教师冲突 2-教室冲突 3-班级冲突 4-其他冲突',
    `conflict_time`            JSON            NOT NULL COMMENT '冲突时间',
    `description`              VARCHAR(255)    NULL COMMENT '冲突描述',
    `resolution_status`        TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '解决状态: 0-未解决 1-已解决 2-忽略',
    `resolution_method`        TINYINT UNSIGNED NULL COMMENT '解决方法: 1-调整第一个课程 2-调整第二个课程 3-同时调整 4-其他',
    `resolution_notes`         VARCHAR(255)    NULL COMMENT '解决备注',
    `resolved_by`              CHAR(32)        NULL COMMENT '解决人',
    `resolved_at`              TIMESTAMP       NULL COMMENT '解决时间',
    `created_at`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '排课冲突表';

-- 添加索引以优化查询性能
CREATE INDEX `idx_conflict_semester` ON `cs_scheduling_conflict` (`semester_uuid`) COMMENT '学期索引';
CREATE INDEX `idx_conflict_first_assignment` ON `cs_scheduling_conflict` (`first_assignment_uuid`) COMMENT '第一个排课索引';
CREATE INDEX `idx_conflict_second_assignment` ON `cs_scheduling_conflict` (`second_assignment_uuid`) COMMENT '第二个排课索引';
CREATE INDEX `idx_conflict_type` ON `cs_scheduling_conflict` (`conflict_type`) COMMENT '冲突类型索引';
CREATE INDEX `idx_conflict_resolution_status` ON `cs_scheduling_conflict` (`resolution_status`) COMMENT '解决状态索引';
CREATE INDEX `idx_conflict_resolved_by` ON `cs_scheduling_conflict` (`resolved_by`) COMMENT '解决人索引';

-- 添加复合索引，用于快速查找某个排课的所有冲突
CREATE INDEX `idx_conflict_assignments` ON `cs_scheduling_conflict` (`first_assignment_uuid`, `second_assignment_uuid`) COMMENT '排课组合索引';
CREATE INDEX `idx_conflict_status_type` ON `cs_scheduling_conflict` (`resolution_status`, `conflict_type`) COMMENT '状态类型组合索引';

-- 添加外键约束
ALTER TABLE `cs_scheduling_conflict`
    ADD CONSTRAINT `fk_scheduling_conflict_semester_uuid`
        FOREIGN KEY (`semester_uuid`) REFERENCES `cs_semester` (`semester_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_scheduling_conflict_first_assignment`
        FOREIGN KEY (`first_assignment_uuid`) REFERENCES `cs_class_assignment` (`class_assignment_uuid`)
            ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_scheduling_conflict_second_assignment`
        FOREIGN KEY (`second_assignment_uuid`) REFERENCES `cs_class_assignment` (`class_assignment_uuid`)
            ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_scheduling_conflict_resolved_by`
        FOREIGN KEY (`resolved_by`) REFERENCES `cs_user` (`user_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE;

-- 创建视图，用于快速查看未解决的冲突
CREATE OR REPLACE VIEW `v_unresolved_conflicts` AS
SELECT 
    sc.conflict_uuid,
    sc.conflict_type,
    sc.description,
    sc.resolution_status,
    ca1.teaching_class_name AS first_class_name,
    ca2.teaching_class_name AS second_class_name,
    t1.name AS first_teacher_name,
    t2.name AS second_teacher_name,
    cr1.name AS first_classroom_name,
    cr2.name AS second_classroom_name,
    sc.conflict_time,
    sc.created_at
FROM 
    `cs_scheduling_conflict` sc
    JOIN `cs_class_assignment` ca1 ON sc.first_assignment_uuid = ca1.class_assignment_uuid
    JOIN `cs_class_assignment` ca2 ON sc.second_assignment_uuid = ca2.class_assignment_uuid
    JOIN `cs_teacher` t1 ON ca1.teacher_uuid = t1.teacher_uuid
    JOIN `cs_teacher` t2 ON ca2.teacher_uuid = t2.teacher_uuid
    JOIN `cs_classroom` cr1 ON ca1.classroom_uuid = cr1.classroom_uuid
    JOIN `cs_classroom` cr2 ON ca2.classroom_uuid = cr2.classroom_uuid
WHERE 
    sc.resolution_status = 0
ORDER BY 
    sc.created_at DESC;