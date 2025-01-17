USE `class-scheduling-system`;
CREATE TABLE `cs_building`
(
    `building_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '教学楼主键',
    `building_name` VARCHAR(32) NOT NULL COMMENT '教学楼名称',
    `campus_uuid`   CHAR(32)    NOT NULL COMMENT '校区主键',
    `is_status`     BOOLEAN     NOT NULL DEFAULT TRUE COMMENT '教学楼状态 0:禁用 1:启用',
    `created_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '教学楼表';

ALTER TABLE `cs_building`
    ADD CONSTRAINT `fk_cs_building_cs_campus`
        FOREIGN KEY (`campus_uuid`) REFERENCES `cs_campus` (`campus_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE;