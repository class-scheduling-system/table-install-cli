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
CREATE TABLE `cs_major`
(
    `major_uuid`        CHAR(32)          NOT NULL PRIMARY KEY COMMENT '专业主键',
    `major_name`        VARCHAR(32)       NOT NULL COMMENT '专业名称',
    `major_description` VARCHAR(255)      NULL COMMENT '专业描述',
    `major_code`        VARCHAR(32)       NOT NULL COMMENT '专业代码',
    `major_status`      BOOLEAN           NOT NULL DEFAULT TRUE COMMENT '专业状态 0:禁用 1:启用',
    `department_uuid`   CHAR(32)          NOT NULL COMMENT '学院外键',
    `education_years`   SMALLINT UNSIGNED NOT NULL COMMENT '学制（年）',
    `training_level`    VARCHAR(32)       NOT NULL COMMENT '培养层次（例如：本科，专科）',
    `created_at`        TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '专业表';

ALTER TABLE `cs_major`
    ADD CONSTRAINT `fk_department_uuid`
        FOREIGN KEY (`department_uuid`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;

CREATE UNIQUE INDEX `uk_major_name` ON `cs_major` (`major_name`) COMMENT '专业名称唯一索引';
CREATE UNIQUE INDEX `uk_major_code` ON `cs_major` (`major_code`) COMMENT '专业代码唯一索引';
