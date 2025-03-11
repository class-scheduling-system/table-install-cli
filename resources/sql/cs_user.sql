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
CREATE TABLE `cs_user`
(
    `user_uuid`  CHAR(32)     NOT NULL PRIMARY KEY COMMENT '用户主键',
    `name`       VARCHAR(32)  NOT NULL COMMENT '用户名',
    `password`   CHAR(60)     NOT NULL COMMENT '用户密码',
    `email`      VARCHAR(255) NOT NULL COMMENT '用户邮箱',
    `phone`      VARCHAR(11)  NOT NULL COMMENT '用户手机号',
    `status`     BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '用户状态 0:禁用 1:启用',
    `ban`        BOOLEAN      NOT NULL DEFAULT FALSE COMMENT '用户是否被封禁 0:未封禁 1:已封禁',
    `role_uuid`  CHAR(32)     NOT NULL COMMENT '角色UUID',
    `permission` JSON NULL COMMENT '用户权限',
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '用户表';

CREATE UNIQUE INDEX `uk_user_email` ON `cs_user` (`email`) COMMENT '用户邮箱唯一索引';
CREATE UNIQUE INDEX `uk_user_phone` ON `cs_user` (`phone`) COMMENT '用户手机号唯一索引';
CREATE UNIQUE INDEX `uk_user_name` ON `cs_user` (`name`) COMMENT '用户名唯一索引';

-- 添加优化查询的索引
CREATE INDEX `idx_user_status` ON `cs_user` (`status`) COMMENT '用户状态索引';
CREATE INDEX `idx_user_ban` ON `cs_user` (`ban`) COMMENT '用户封禁状态索引';
CREATE INDEX `idx_user_role_uuid` ON `cs_user` (`role_uuid`) COMMENT '用户角色索引';

ALTER TABLE `cs_user`
    ADD CONSTRAINT `fk_user_role_uuid`
        FOREIGN KEY (`role_uuid`) REFERENCES `cs_role` (`role_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;