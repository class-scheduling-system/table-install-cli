USE `class-scheduling-system`;
CREATE TABLE `cs_role`
(
    `role_uuid`   CHAR(32)    NOT NULL PRIMARY KEY COMMENT '角色主键',
    `role_name`   VARCHAR(32) NOT NULL COMMENT '角色名',
    `role_status` BOOLEAN     NOT NULL DEFAULT TRUE COMMENT '角色状态 0:禁用 1:启用',
    `permission`  JSON        NULL COMMENT '角色权限',
    `created_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '角色表';

CREATE UNIQUE INDEX `role_name` ON `cs_role` (`role_name`) COMMENT '角色名唯一索引';