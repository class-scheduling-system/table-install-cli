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
CREATE TABLE `cs_student`
(
    `student_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '学生主键',
    `id`           VARCHAR(32) NOT NULL COMMENT '学号',
    `name`         VARCHAR(32) NOT NULL COMMENT '学生姓名',
    `gender`       BOOLEAN     NOT NULL COMMENT '性别 0:女 1:男',
    `grade`        VARCHAR(32) NOT NULL COMMENT '年级',
    `department`   CHAR(32)    NOT NULL COMMENT '所属学院',
    `major`        CHAR(32)    NOT NULL COMMENT '所属专业',
    `class`        CHAR(32)    NULL COMMENT '班级',
    `user_uuid`    CHAR(32)    NULL COMMENT '对应用户主键',
    `created_at`   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '学生表';

CREATE UNIQUE INDEX `uk_student_id` ON `cs_student` (`id`) COMMENT '学号唯一索引';

-- 添加优化查询的索引
CREATE INDEX `idx_student_name` ON `cs_student` (`name`) COMMENT '学生姓名索引';
CREATE INDEX `idx_student_department` ON `cs_student` (`department`) COMMENT '学院索引';
CREATE INDEX `idx_student_major` ON `cs_student` (`major`) COMMENT '专业索引';
CREATE INDEX `idx_student_class` ON `cs_student` (`class`) COMMENT '班级索引';
CREATE INDEX `idx_student_grade` ON `cs_student` (`grade`) COMMENT '年级索引';
CREATE INDEX `idx_student_department_major_class` ON `cs_student` (`department`, `major`, `class`) COMMENT '院系专业班级组合索引';

ALTER TABLE `cs_student`
    ADD CONSTRAINT `fk_cs_student_cs_department`
        FOREIGN KEY (`department`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_student_cs_major`
        FOREIGN KEY (`major`) REFERENCES `cs_major` (`major_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_student_cs_user`
        FOREIGN KEY (`user_uuid`) REFERENCES `cs_user` (`user_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_administrative_class_cs_student`
        FOREIGN KEY (`class`) REFERENCES `cs_administrative_class` (`administrative_class_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;
