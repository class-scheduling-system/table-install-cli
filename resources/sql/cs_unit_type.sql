USE `class-scheduling-system`;
CREATE TABLE `cs_unit_type`
(
    `unit_type_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '单位办别主键',
    `name`           VARCHAR(32)  NOT NULL COMMENT '单位名称',
    `english_name`   VARCHAR(128) NULL COMMENT '单位英文名称',
    `short_name`     VARCHAR(32)  NULL COMMENT '单位简称',
    `order`          INT          NOT NULL DEFAULT 100 COMMENT '单位排序',
    `created_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '单位办别表';

CREATE UNIQUE INDEX `name_UNIQUE` ON `cs_unit_type` (`name`) COMMENT '单位名称唯一索引';