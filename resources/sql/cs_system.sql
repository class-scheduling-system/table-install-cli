USE `class-scheduling-system`;
CREATE TABLE `cs_system`
(
    `system_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '系统主键',
    `system_key`  VARCHAR(32) NOT NULL COMMENT '系统键',
    `system_val`  TEXT        NULL COMMENT '系统值',
    `created_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '系统表';

CREATE UNIQUE INDEX `uk_system_key` ON `cs_system` (`system_key`) COMMENT '系统键唯一索引';