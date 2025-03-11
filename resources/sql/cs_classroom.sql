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
CREATE TABLE `cs_classroom`
(
    `classroom_uuid`            CHAR(32)       NOT NULL PRIMARY KEY COMMENT '教室主键',
    `number`                    VARCHAR(32)    NOT NULL COMMENT '教室编号',
    `name`                      VARCHAR(32)    NOT NULL COMMENT '教室名称',
    `campus_uuid`               CHAR(32)       NOT NULL COMMENT '校区主键',
    `building_uuid`             CHAR(32)       NOT NULL COMMENT '楼栋主键',
    `floor`                 VARCHAR(4) NOT NULL COMMENT '楼层',
    `type`                      CHAR(32)       NOT NULL COMMENT '教室类型',
    `tag`                       JSON           NULL COMMENT '教室标签',
    `capacity`                  INT            NOT NULL COMMENT '教室容量',
    `examination_room`          BOOLEAN        NOT NULL DEFAULT FALSE COMMENT '是否是考场',
    `examination_room_capacity` INT            NULL COMMENT '考场容量',
    `is_multimedia`             BOOLEAN        NOT NULL DEFAULT FALSE COMMENT '是否是多媒体教室',
    `is_air_conditioned`        BOOLEAN        NOT NULL DEFAULT FALSE COMMENT '是否有空调',
    `status`                    BOOLEAN        NOT NULL DEFAULT TRUE COMMENT '教室状态 0:禁用 1:启用',
    `description`               VARCHAR(255)   NULL COMMENT '教室描述',
    `management_department` CHAR(32)   NULL COMMENT '管理部门',
    `area`                      DECIMAL(10, 2) NOT NULL COMMENT '教室面积',
    `tables_chairs_type`    CHAR(32)   NULL COMMENT '桌椅类型',
    `created_at`                TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '教室表';

CREATE UNIQUE INDEX `uk_classroom_number` ON `cs_classroom` (`number`) COMMENT '教室编号唯一索引';

-- 添加优化查询的索引
CREATE INDEX `idx_classroom_campus_building` ON `cs_classroom` (`campus_uuid`, `building_uuid`) COMMENT '校区楼栋组合索引';
CREATE INDEX `idx_classroom_capacity` ON `cs_classroom` (`capacity`) COMMENT '容量索引';
CREATE INDEX `idx_classroom_type` ON `cs_classroom` (`type`) COMMENT '类型索引';
CREATE INDEX `idx_classroom_status` ON `cs_classroom` (`status`) COMMENT '状态索引';
CREATE INDEX `idx_classroom_examination` ON `cs_classroom` (`examination_room`) COMMENT '考场索引';
CREATE INDEX `idx_classroom_multimedia` ON `cs_classroom` (`is_multimedia`) COMMENT '多媒体索引';

ALTER TABLE `cs_classroom`
    ADD CONSTRAINT `fk_cs_classroom_cs_campus`
        FOREIGN KEY (`campus_uuid`) REFERENCES `cs_campus` (`campus_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_building`
        FOREIGN KEY (`building_uuid`) REFERENCES `cs_building` (`building_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_classroom_type`
        FOREIGN KEY (`type`) REFERENCES `cs_classroom_type` (`class_type_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_management_department`
        FOREIGN KEY (`management_department`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_tables_chairs_type`
        FOREIGN KEY (`tables_chairs_type`) REFERENCES `cs_tables_chairs_type` (`tables_chairs_type_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE;
