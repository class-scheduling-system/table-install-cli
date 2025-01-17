USE `class-scheduling-system`;
CREATE TABLE `cs_classroom`
(
    `classroom_uuid`            CHAR(32)       NOT NULL PRIMARY KEY COMMENT '教室主键',
    `number`                    VARCHAR(32)    NOT NULL COMMENT '教室编号',
    `name`                      VARCHAR(32)    NOT NULL COMMENT '教室名称',
    `campus_uuid`               CHAR(32)       NOT NULL COMMENT '校区主键',
    `building_uuid`             CHAR(32)       NOT NULL COMMENT '楼栋主键',
    `floor`                     INT            NOT NULL COMMENT '楼层',
    `type`                      CHAR(32)       NOT NULL COMMENT '教室类型',
    `tag`                       JSON           NULL COMMENT '教室标签',
    `capacity`                  INT            NOT NULL COMMENT '教室容量',
    `examination_room`          BOOLEAN        NOT NULL DEFAULT FALSE COMMENT '是否是考场',
    `examination_room_capacity` INT            NULL COMMENT '考场容量',
    `is_multimedia`             BOOLEAN        NOT NULL DEFAULT FALSE COMMENT '是否是多媒体教室',
    `is_air_conditioned`        BOOLEAN        NOT NULL DEFAULT FALSE COMMENT '是否有空调',
    `status`                    BOOLEAN        NOT NULL DEFAULT TRUE COMMENT '教室状态 0:禁用 1:启用',
    `description`               VARCHAR(255)   NULL COMMENT '教室描述',
    `management_department`     CHAR(32)       NULL COMMENT '管理部门',
    `area`                      DECIMAL(10, 2) NOT NULL COMMENT '教室面积',
    `tables_chairs_type`        CHAR(32)       NULL COMMENT '桌椅类型',
    `created_at`                TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '教室表';

CREATE UNIQUE INDEX `uk_classroom_number` ON `cs_classroom` (`number`) COMMENT '教室编号唯一索引';

ALTER TABLE `cs_classroom`
    ADD CONSTRAINT `fk_cs_classroom_cs_campus`
        FOREIGN KEY (`campus_uuid`) REFERENCES `cs_campus` (`campus_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_building`
    FOREIGN KEY (`building_uuid`) REFERENCES `cs_building` (`building_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_classroom_type`
        FOREIGN KEY (`type`) REFERENCES `cs_classroom_type` (`class_type_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_management_department`
        FOREIGN KEY (`management_department`) REFERENCES `cs_department` (`department_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_classroom_cs_tables_chairs_type`
        FOREIGN KEY (`tables_chairs_type`) REFERENCES `cs_tables_chairs_type` (`tables_chairs_type_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE;