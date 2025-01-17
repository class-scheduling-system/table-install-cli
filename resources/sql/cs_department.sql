USE `class-scheduling-system`;
CREATE TABLE `cs_department`
(
    `department_uuid`            CHAR(32)     NOT NULL PRIMARY KEY COMMENT '部门主键',
    `department_code`            VARCHAR(32)  NOT NULL COMMENT '部门编码',
    `department_name`            VARCHAR(64)  NOT NULL COMMENT '部门名称',
    `department_order`           INT          NOT NULL DEFAULT 100 COMMENT '部门排序',
    `department_english_name`    VARCHAR(128) NULL COMMENT '部门英文名称',
    `department_short_name`      VARCHAR(32)  NULL COMMENT '部门简称',
    `department_address`         VARCHAR(255) NULL COMMENT '部门地址',
    `is_entity`                  BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否实体部门',
    `administrative_head`        VARCHAR(32)  NULL COMMENT '行政负责人',
    `party_committee_head`       VARCHAR(32)  NULL COMMENT '党委负责人',
    `establishment_date`         DATE         NOT NULL DEFAULT (CURRENT_DATE) COMMENT '成立日期',
    `expiration_date`            DATE         NULL COMMENT '失效日期',
    `unit_category`              CHAR(32)     NOT NULL COMMENT '单位类别',
    `unit_type`                  CHAR(32)     NOT NULL COMMENT '单位办别',
    `parent_department`          CHAR(32)     NULL COMMENT '上级部门',
    `assigned_teaching_building` CHAR(32)     NULL COMMENT '分配教学楼',
    `is_teaching_college`        BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '开课院系',
    `is_attending_college`       BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '上课院系',
    `fixed_phone`                VARCHAR(32)  NULL COMMENT '固定电话',
    `remark`                     TEXT         NULL COMMENT '备注',
    `is_teaching_office`         BOOLEAN      NOT NULL DEFAULT FALSE COMMENT '开课教研室',
    `is_enabled`                 BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '是否启用',
    `created_at`                 TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                 TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '部门表';

CREATE UNIQUE INDEX `uk_department_code` ON `cs_department` (`department_code`) COMMENT '部门编码唯一索引';

ALTER TABLE `cs_department`
    ADD CONSTRAINT `fk_department_unit_category`
        FOREIGN KEY (`unit_category`) REFERENCES `cs_unit_category` (`unit_category_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_department_unit_type`
        FOREIGN KEY (`unit_type`) REFERENCES `cs_unit_type` (`unit_type_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_department_parent_department`
        FOREIGN KEY (`parent_department`) REFERENCES `cs_department` (`department_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_department_assigned_teaching_building`
        FOREIGN KEY (`assigned_teaching_building`) REFERENCES `cs_building` (`building_uuid`)
            ON DELETE RESTRICT ON UPDATE CASCADE;
