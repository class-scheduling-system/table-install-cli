USE `class-scheduling-system`;
CREATE TABLE `cs_tables_chairs_type`
(
    `tables_chairs_type_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '桌椅类型主键',
    `name`                    VARCHAR(32)  NOT NULL COMMENT '桌椅类型名称',
    `description`             VARCHAR(255) NULL COMMENT '桌椅类型描述',
    `base64_img`              TEXT         NULL COMMENT '桌椅类型图片',
    `created_at`              TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`              TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '桌椅类型表';

CREATE UNIQUE INDEX `uk_tables_chairs_type_name` ON `cs_tables_chairs_type` (`name`) COMMENT '桌椅类型名称唯一索引';