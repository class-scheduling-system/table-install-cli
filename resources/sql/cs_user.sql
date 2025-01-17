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
    `role_uuid`  CHAR(36)     NOT NULL COMMENT '角色UUID',
    `permission` JSON         NOT NULL COMMENT '用户权限',
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

CREATE UNIQUE INDEX `uk_user_email` ON `cs_user` (`email`) COMMENT '用户邮箱唯一索引';
CREATE UNIQUE INDEX `uk_user_phone` ON `cs_user` (`phone`) COMMENT '用户手机号唯一索引';
CREATE UNIQUE INDEX `uk_user_name` ON `cs_user` (`name`) COMMENT '用户名唯一索引';

ALTER TABLE `cs_user`
    ADD CONSTRAINT `fk_user_role_uuid`
        FOREIGN KEY (`role_uuid`) REFERENCES `cs_role` (`role_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;