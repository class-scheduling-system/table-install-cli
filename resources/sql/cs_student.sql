USE `class-scheduling-system`;
CREATE TABLE `cs_student`
(
    `student_uuid` CHAR(32)    NOT NULL PRIMARY KEY COMMENT '学生主键',
    `id`   VARCHAR(32) NOT NULL COMMENT '学号',
    `name` VARCHAR(32) NOT NULL COMMENT '学生姓名',
    `gender`       BOOLEAN     NOT NULL COMMENT '性别 0:女 1:男',
    `grade`        VARCHAR(32) NOT NULL COMMENT '年级',
    `department`   CHAR(32)    NOT NULL COMMENT '学院',
    `major`        CHAR(32)    NOT NULL COMMENT '专业',
    `class`        VARCHAR(32) NOT NULL COMMENT '班级',
    `user_uuid`    CHAR(32)    NULL COMMENT '对应用户主键',
    `created_at`   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '学生表';

CREATE UNIQUE INDEX `uk_student_id` ON `cs_student` (`id`) COMMENT '学号唯一索引';

ALTER TABLE `cs_student`
    ADD CONSTRAINT `fk_cs_student_cs_department`
        FOREIGN KEY (`department`) REFERENCES `cs_department` (`department_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_student_cs_major`
        FOREIGN KEY (`major`) REFERENCES `cs_major` (`major_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_cs_student_cs_user`
        FOREIGN KEY (`user_uuid`) REFERENCES `cs_user` (`user_uuid`) ON DELETE RESTRICT ON UPDATE CASCADE;