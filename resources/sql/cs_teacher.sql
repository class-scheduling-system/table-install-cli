USE `class-scheduling-system`;
CREATE TABLE `cs_teacher`
(
    `teacher_uuid` CHAR(32)     NOT NULL PRIMARY KEY COMMENT '教师主键',
    `unit_uuid`    CHAR(32)     NOT NULL COMMENT '单位主键',
    `user_uuid`    CHAR(32)     NULL COMMENT '用户主键',
    `id`           VARCHAR(32)  NOT NULL COMMENT '教师工号',
    `name`         VARCHAR(32)  NOT NULL COMMENT '教师姓名',
    `english_name` VARCHAR(128) NOT NULL COMMENT '教师英文名',
    `ethnic`       CHAR(32)     NOT NULL COMMENT '教师民族',
    `sex`          BOOLEAN      NOT NULL DEFAULT TRUE COMMENT '教师性别 0:女 1:男',
    `phone`        VARCHAR(16)  NULL COMMENT '教师电话',
    `email`        VARCHAR(255) NULL COMMENT '教师邮箱',
    `job_title`    VARCHAR(32)  NULL COMMENT '教师职称',
    `desc`         VARCHAR(255) NULL COMMENT '教师描述',
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '教师表';

CREATE UNIQUE INDEX `uk_teacher_id` ON `cs_teacher` (`id`) COMMENT '教师工号唯一索引';

ALTER TABLE `cs_teacher`
    ADD CONSTRAINT `fk_teacher_unit_uuid`
        FOREIGN KEY (`unit_uuid`) REFERENCES `cs_department` (`department_uuid`),
    ADD CONSTRAINT `fk_teacher_user_uuid`
        FOREIGN KEY (`user_uuid`) REFERENCES `cs_user` (`user_uuid`);