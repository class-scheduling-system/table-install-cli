USE `class-scheduling-system`;
CREATE TABLE `cs_academic_affairs_permission`
(
    `academic_affairs_permission_uuid` CHAR(32)  NOT NULL PRIMARY KEY COMMENT '教务权限主键',
    `authorized_user`                  CHAR(32)  NOT NULL COMMENT '授权用户',
    `department`                       CHAR(32)  NOT NULL COMMENT '部门（要求该部门为院系）',
    `type`                             TINYINT   NOT NULL COMMENT '类型 0:所有权限, 1:教务权限...',
    `created_at`                       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`                       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '教务权限表';

ALTER TABLE `cs_academic_affairs_permission`
    ADD CONSTRAINT `cs_academic_affairs_permission_user_uuid_foreign`
        FOREIGN KEY (`authorized_user`) REFERENCES `cs_user` (`user_uuid`) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `cs_academic_affairs_permission_department_uuid_foreign`
        FOREIGN KEY (`department`) REFERENCES `cs_department` (`department_uuid`) ON DELETE CASCADE ON UPDATE CASCADE;