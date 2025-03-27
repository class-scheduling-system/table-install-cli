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

CREATE TABLE `cs_teacher_course_qualification`
(
    `qualification_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '资格主键',
    `teacher_uuid`      CHAR(32)    NOT NULL COMMENT '教师主键',
    `course_uuid`       CHAR(32)    NOT NULL COMMENT '课程主键',
    `qualification_level` TINYINT    NOT NULL DEFAULT 1 COMMENT '资格等级 1:初级 2:中级 3:高级',
    `is_primary`        BOOLEAN     NOT NULL DEFAULT FALSE COMMENT '是否主讲教师',
    `teach_years`       TINYINT     NOT NULL DEFAULT 0 COMMENT '教授年限',
    `status`            TINYINT     NOT NULL DEFAULT 1 COMMENT '状态 0:待审核 1:已审核 2:已驳回',
    `remarks`           VARCHAR(255) NULL COMMENT '备注说明',
    `approved_by`       CHAR(32)    NULL COMMENT '审核人',
    `approved_at`       TIMESTAMP   NULL COMMENT '审核时间',
    `created_at`        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT '教师课程资格表';

-- 创建唯一索引确保教师和课程的组合唯一
CREATE UNIQUE INDEX `uk_teacher_course` ON `cs_teacher_course_qualification` (`teacher_uuid`, `course_uuid`) 
    COMMENT '教师课程唯一索引';

-- 添加优化查询的索引
CREATE INDEX `idx_qualification_teacher` ON `cs_teacher_course_qualification` (`teacher_uuid`) 
    COMMENT '教师索引';
CREATE INDEX `idx_qualification_course` ON `cs_teacher_course_qualification` (`course_uuid`) 
    COMMENT '课程索引';
CREATE INDEX `idx_qualification_status` ON `cs_teacher_course_qualification` (`status`) 
    COMMENT '状态索引';
CREATE INDEX `idx_qualification_level` ON `cs_teacher_course_qualification` (`qualification_level`) 
    COMMENT '资格等级索引';
CREATE INDEX `idx_qualification_is_primary` ON `cs_teacher_course_qualification` (`is_primary`) 
    COMMENT '主讲教师索引';

-- 添加外键约束
ALTER TABLE `cs_teacher_course_qualification`
    ADD CONSTRAINT `fk_qualification_teacher`
        FOREIGN KEY (`teacher_uuid`) REFERENCES `cs_teacher` (`teacher_uuid`)
            ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_qualification_course`
        FOREIGN KEY (`course_uuid`) REFERENCES `cs_course_library` (`course_library_uuid`)
            ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_qualification_approved_by`
        FOREIGN KEY (`approved_by`) REFERENCES `cs_user` (`user_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE;