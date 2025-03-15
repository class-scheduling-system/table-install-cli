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
CREATE TABLE `cs_department`
(
    `department_uuid`            CHAR(32)     NOT NULL PRIMARY KEY COMMENT '部门主键',
    `department_code`            VARCHAR(32)  NOT NULL COMMENT '部门编码',
    `department_name`            VARCHAR(64)  NOT NULL COMMENT '部门名称',
    `department_order`           INT          NOT NULL DEFAULT 100 COMMENT '部门排序',
    `department_english_name`    VARCHAR(128) NULL COMMENT '部门英文名称',
    `department_short_name`      VARCHAR(32)  NULL COMMENT '部门简称',
    `department_address`         VARCHAR(255) NULL COMMENT '部门地址',
    `is_entity`                  BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否实体部门',
    `administrative_head`        VARCHAR(32)  NULL COMMENT '行政负责人',
    `party_committee_head`       VARCHAR(32)  NULL COMMENT '党委负责人',
    `establishment_date`         DATE         NOT NULL DEFAULT (CURRENT_DATE) COMMENT '成立日期',
    `expiration_date`            DATE         NULL COMMENT '失效日期',
    `unit_category`              CHAR(32)     NOT NULL COMMENT '单位类别',
    `unit_type`                  CHAR(32)     NOT NULL COMMENT '单位办别',
    `parent_department`          CHAR(32)     NULL COMMENT '上级部门',
    `assigned_teaching_building` JSON         NULL COMMENT '分配教学楼',
    `is_teaching_college`        BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '开课院系',
    `is_attending_college`       BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '上课院系',
    `fixed_phone`                VARCHAR(32)  NULL COMMENT '固定电话',
    `remark`                     TEXT         NULL COMMENT '备注',
    `is_teaching_office`         BOOLEAN      NOT NULL DEFAULT FALSE COMMENT '开课教研室',
    `is_enabled`                 BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否启用',
    `created_at`                 TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                 TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '部门表';

CREATE UNIQUE INDEX `uk_department_code` ON `cs_department` (`department_code`) COMMENT '部门编码唯一索引';
CREATE UNIQUE INDEX `uk_department_name` ON `cs_department` (`department_name`) COMMENT '部门名称唯一索引';

-- 添加优化查询的索引
CREATE INDEX `idx_department_parent` ON `cs_department` (`parent_department`) COMMENT '上级部门索引';
CREATE INDEX `idx_department_unit_type` ON `cs_department` (`unit_type`) COMMENT '单位办别索引';
CREATE INDEX `idx_department_unit_category` ON `cs_department` (`unit_category`) COMMENT '单位类别索引';
CREATE INDEX `idx_department_teaching_college` ON `cs_department` (`is_teaching_college`) COMMENT '开课院系索引';
CREATE INDEX `idx_department_attending_college` ON `cs_department` (`is_attending_college`) COMMENT '上课院系索引';
CREATE INDEX `idx_department_teaching_office` ON `cs_department` (`is_teaching_office`) COMMENT '开课教研室索引';
CREATE INDEX `idx_department_is_enabled` ON `cs_department` (`is_enabled`) COMMENT '启用状态索引';
CREATE INDEX `idx_department_order` ON `cs_department` (`department_order`) COMMENT '部门排序索引';

ALTER TABLE `cs_department`
    ADD CONSTRAINT `fk_department_unit_category`
        FOREIGN KEY (`unit_category`) REFERENCES `cs_unit_category` (`unit_category_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_department_unit_type`
        FOREIGN KEY (`unit_type`) REFERENCES `cs_unit_type` (`unit_type_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_department_parent_department`
        FOREIGN KEY (`parent_department`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE;
