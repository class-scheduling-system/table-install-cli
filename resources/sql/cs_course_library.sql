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
CREATE TABLE `cs_course_library`
(
    `course_library_uuid` CHAR(32)       NOT NULL PRIMARY KEY COMMENT '课程库主键',
    `id`                  VARCHAR(32)    NOT NULL COMMENT '课程编号',
    `name`                VARCHAR(32)    NOT NULL COMMENT '课程库名称',
    `english_name`        VARCHAR(128)   NULL COMMENT '课程英文名称',
    `category`            CHAR(32)       NULL COMMENT '课程类别',
    `property`            CHAR(32)       NULL COMMENT '课程属性',
    `type`                CHAR(32)       NOT NULL COMMENT '课程类型',
    `nature`              CHAR(32)       NULL COMMENT '课程性质',
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