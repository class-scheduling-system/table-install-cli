USE `class-scheduling-system`;
CREATE TABLE `cs_permission`
(
    `permission_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '权限主键',
    `permission_key`  VARCHAR(255)  NOT NULL COMMENT '权限键 key1.key2.key3......',
    `name`            VARCHAR(128) NOT NULL COMMENT '权限名称',
    `desc`            VARCHAR(255) NULL COMMENT '权限描述'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '权限表';

CREATE UNIQUE INDEX `uk_permission_key` ON `cs_permission` (`permission_key`) COMMENT '权限键唯一索引';