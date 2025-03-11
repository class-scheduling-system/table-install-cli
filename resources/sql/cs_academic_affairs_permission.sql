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
CREATE TABLE `cs_academic_affairs_permission`
(
    `academic_affairs_permission_uuid` CHAR(32)  NOT NULL PRIMARY KEY COMMENT '教务权限主键',
    `authorized_user`                  CHAR(32)  NOT NULL COMMENT '授权用户',
    `department`                       CHAR(32)  NOT NULL COMMENT '部门（要求该部门为院系）',
    `type`                             TINYINT   NOT NULL COMMENT '类型 0:所有权限, 1:教务权限...',
    `created_at`                       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT = '教务权限表';

ALTER TABLE `cs_academic_affairs_permission`
    ADD CONSTRAINT `cs_academic_affairs_permission_user_uuid_foreign`
        FOREIGN KEY (`authorized_user`) REFERENCES `cs_user` (`user_uuid`) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `cs_academic_affairs_permission_department_uuid_foreign`
        FOREIGN KEY (`department`) REFERENCES `cs_department` (`department_uuid`) ON DELETE CASCADE ON UPDATE CASCADE;
