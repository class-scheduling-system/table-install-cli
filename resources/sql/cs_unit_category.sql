USE `class-scheduling-system`;
CREATE TABLE `cs_unit_category`
(
    `unit_category_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '单位类别主键',
    `name`               VARCHAR(32) NOT NULL COMMENT '单位类别名称',
    `order`              INT         NOT NULL DEFAULT 100 COMMENT '单位类别排序',
    `english_name`       VARCHAR(32) NULL COMMENT '单位类别英文名称',
    `short_name`         VARCHAR(32) NULL COMMENT '单位类别简称',
    `is_entity`          BOOLEAN     NOT NULL DEFAULT TRUE COMMENT '是否实体单位类别',
    `created_at`         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '单位类别表';

CREATE UNIQUE INDEX `name_UNIQUE` ON `cs_unit_category` (`name`) COMMENT '单位类别名称唯一索引';