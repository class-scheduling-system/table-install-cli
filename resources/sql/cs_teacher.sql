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
CREATE TABLE `cs_teacher`
(
    `teacher_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '教师主键',
    `unit_uuid`    CHAR(32)     NOT NULL COMMENT '单位主键',
    `user_uuid`    CHAR(32)     NULL COMMENT '用户主键',
    `id`           VARCHAR(64)  NOT NULL COMMENT '教师工号',
    `name`         VARCHAR(32)  NOT NULL COMMENT '教师姓名',
    `english_name` VARCHAR(256) NOT NULL COMMENT '教师英文名',
    `ethnic`       CHAR(32)     NOT NULL COMMENT '教师民族',
    `sex`          BOOLEAN      NOT NULL COMMENT '教师性别 0:女 1:男',
    `type`         CHAR(32)     NOT NULL COMMENT '教师类型',
    `phone`        VARCHAR(16)  NULL COMMENT '教师电话',
    `email`        VARCHAR(255) NULL COMMENT '教师邮箱',
    `job_title`    VARCHAR(32)  NULL COMMENT '教师职称',
    `desc`         VARCHAR(255) NULL COMMENT '教师描述',
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '教师表';

CREATE UNIQUE INDEX `uk_teacher_id` ON `cs_teacher` (`id`) COMMENT '教师工号唯一索引';

-- 添加优化查询的索引
CREATE INDEX `idx_teacher_name` ON `cs_teacher` (`name`) COMMENT '教师姓名索引';
CREATE INDEX `idx_teacher_unit` ON `cs_teacher` (`unit_uuid`) COMMENT '教师单位索引';
CREATE INDEX `idx_teacher_job_title` ON `cs_teacher` (`job_title`) COMMENT '教师职称索引';
CREATE INDEX `idx_teacher_sex` ON `cs_teacher` (`sex`) COMMENT '教师性别索引';

ALTER TABLE `cs_teacher`
    ADD CONSTRAINT `fk_teacher_unit_uuid`
        FOREIGN KEY (`unit_uuid`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_teacher_user_uuid`
        FOREIGN KEY (`user_uuid`) REFERENCES `cs_user` (`user_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_teacher_type`
        FOREIGN KEY (`type`) REFERENCES `cs_teacher_type` (`teacher_type_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;
