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
CREATE TABLE `cs_teacher_preferences`
(
    `preference_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '教师喜好主键',
    `teacher_uuid`    CHAR(32)     NOT NULL COMMENT '教师主键',
    `semester_uuid`   CHAR(32)     NOT NULL COMMENT '学期主键',
    `day_of_week`    TINYINT      NOT NULL COMMENT '星期几（1-7）',
    `time_slot`      TINYINT      NOT NULL COMMENT '第几节课（1-12）',
    `preference_level` TINYINT     NOT NULL COMMENT '偏好程度（1：最不期望，2：尽量避免，3：可接受，4：较期望，5：非常期望）',
    `reason`         VARCHAR(255)  NULL COMMENT '偏好原因',
    `created_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '教师课程时间偏好表';

-- 添加优化查询的索引
CREATE INDEX `idx_teacher_preferences_teacher` ON `cs_teacher_preferences` (`teacher_uuid`) COMMENT '教师索引';
CREATE INDEX `idx_teacher_preferences_semester` ON `cs_teacher_preferences` (`semester_uuid`) COMMENT '学期索引';
CREATE INDEX `idx_teacher_preferences_time` ON `cs_teacher_preferences` (`day_of_week`, `time_slot`) COMMENT '时间索引';

ALTER TABLE `cs_teacher_preferences`
    ADD CONSTRAINT `fk_teacher_preferences_teacher`
        FOREIGN KEY (`teacher_uuid`) REFERENCES `cs_teacher` (`teacher_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_teacher_preferences_semester`
        FOREIGN KEY (`semester_uuid`) REFERENCES `cs_semester` (`semester_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;
