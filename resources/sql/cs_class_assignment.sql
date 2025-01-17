USE `class-scheduling-system`;
CREATE TABLE `cs_class_assignment`
(
    `class_assignment_uuid`           CHAR(32)          NOT NULL PRIMARY KEY COMMENT '排课主键',
    `semester_uuid`                   CHAR(32)          NOT NULL COMMENT '学期主键',
    `course_uuid`                     CHAR(32)          NOT NULL COMMENT '课程主键',
    `teacher_uuid`                    CHAR(32)          NOT NULL COMMENT '教师主键',
    `classroom_uuid`                  CHAR(32)          NOT NULL COMMENT '教室主键',
    `teaching_class_composition_uuid` JSON              NOT NULL COMMENT '教学班组成',
    `course_ownership`                VARCHAR(32)       NOT NULL COMMENT '课程归属',
    `teaching_class_name`             VARCHAR(64)       NOT NULL COMMENT '教学班名称',
    `credit_hour_type`                CHAR(32)          NOT NULL COMMENT '学时类型',
    `teaching_hours`                  DECIMAL(10, 2)    NOT NULL COMMENT '教学学时',
    `scheduled_hours`                 DECIMAL(10, 2)    NOT NULL COMMENT '排课学时',
    `total_hours`                     DECIMAL(10, 2)    NOT NULL COMMENT '总学时',
    `scheduling_priority`             SMALLINT UNSIGNED NOT NULL DEFAULT 100 COMMENT '排课优先级',
    `class_size`                      INT UNSIGNED      NOT NULL COMMENT '班级规模',
    `teaching_campus`                 CHAR(32)          NOT NULL COMMENT '教学校区',
    `class_time`                      JSON              NOT NULL COMMENT '上课时间',
    `consecutive_sessions`            TINYINT UNSIGNED  NOT NULL DEFAULT 2 COMMENT '连堂节数',
    `classroom_type`                  CHAR(32)          NOT NULL COMMENT '教室类型',
    `designated_classroom`            CHAR(32)          NULL COMMENT '是否指定教室',
    `designated_teaching_building`    CHAR(32)          NULL COMMENT '指定教学楼',
    `specified_time`                  JSON              NULL COMMENT '指定时间',
    `created_at`                      TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                      TIMESTAMP         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '排课表';

ALTER TABLE `cs_class_assignment`
    ADD CONSTRAINT `fk_class_assignment_semester_uuid`
        FOREIGN KEY (`semester_uuid`) REFERENCES `cs_semester` (`semester_uuid`),
    ADD CONSTRAINT `fk_class_assignment_course_uuid`
        FOREIGN KEY (`course_uuid`) REFERENCES `cs_course_library` (`course_library_uuid`),
    ADD CONSTRAINT `fk_class_assignment_teacher_uuid`
        FOREIGN KEY (`teacher_uuid`) REFERENCES `cs_teacher` (`teacher_uuid`),
    ADD CONSTRAINT `fk_class_assignment_classroom_uuid`
        FOREIGN KEY (`classroom_uuid`) REFERENCES `cs_classroom` (`classroom_uuid`),
    ADD CONSTRAINT `fk_class_assignment_credit_hour_type`
        FOREIGN KEY (`credit_hour_type`) REFERENCES `cs_credit_hour_type` (`credit_hour_type_uuid`),
    ADD CONSTRAINT `fk_class_assignment_teaching_campus`
        FOREIGN KEY (`teaching_campus`) REFERENCES `cs_campus` (`campus_uuid`),
    ADD CONSTRAINT `fk_class_assignment_classroom_type`
        FOREIGN KEY (`classroom_type`) REFERENCES `cs_classroom_type` (`class_type_uuid`),
    ADD CONSTRAINT `fk_class_assignment_designated_classroom`
        FOREIGN KEY (`designated_classroom`) REFERENCES `cs_classroom` (`classroom_uuid`),
    ADD CONSTRAINT `fk_class_assignment_designated_teaching_building`
        FOREIGN KEY (`designated_teaching_building`) REFERENCES `cs_building` (`building_uuid`);