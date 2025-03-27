/*
 * --------------------------------------------------------------------------------
 * Copyright (c) 2022-NOW(至今) 锋楪技术团队
 * Author: 锋楪技术团队 (https://www.frontleaves.com)
 *
 * 本文件包含锋楪技术团队项目的源代码，项目的所有源代码均遵循 MIT 开源许可证协议。
 * --------------------------------------------------------------------------------
 * 许可证声明：
 *
 * 版权所有 (c) 2022-2025 锋楪技术团队。保留所有权利。
 *
 * 本软件是"按原样"提供的，没有任何形式的明示或暗示的保证，包括但不限于
 * 对适销性、特定用途的适用性和非侵权性的暗示保证。在任何情况下，
 * 作者或版权持有人均不承担因软件或软件的使用或其他交易而产生的、
 * 由此引起的或以任何方式与此软件有关的任何索赔、损害或其他责任。
 *
 * 使用本软件即表示您了解此声明并同意其条款。
 *
 * 有关 MIT 许可证的更多信息，请查看项目根目录下的 LICENSE 文件或访问：
 * https://opensource.org/licenses/MIT
 * --------------------------------------------------------------------------------
 * 免责声明：
 *
 * 使用本软件的风险由用户自担。作者或版权持有人在法律允许的最大范围内，
 * 对因使用本软件内容而导致的任何直接或间接的损失不承担任何责任。
 * --------------------------------------------------------------------------------
 */

USE `class-scheduling-system`;

-- 创建请求日志表
CREATE TABLE cs_request_log (
  `request_log_uuid` CHAR(32) NOT NULL PRIMARY KEY COMMENT '请求日志主键',
  `user_uuid` CHAR(32) DEFAULT NULL COMMENT '用户UUID',
  `request_ip` VARCHAR(50) NOT NULL COMMENT '请求IP地址',
  `user_agent` VARCHAR(500) NOT NULL COMMENT '用户代理信息',
  `request_url` VARCHAR(500) NOT NULL COMMENT '请求URL',
  `request_method` VARCHAR(10) NOT NULL COMMENT '请求方法(GET/POST等)',
  `request_params` TEXT COMMENT '请求参数',
  `request_body` TEXT COMMENT '请求体',
  `response_code` INT NOT NULL COMMENT '响应状态码',
  `error_message` VARCHAR(255) COMMENT '错误信息',
  `execution_time` BIGINT NOT NULL COMMENT '执行时间(毫秒)',
  `request_time` TIMESTAMP NOT NULL COMMENT '请求时间',
  `response_time` TIMESTAMP NOT NULL COMMENT '响应时间',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='请求日志表';

-- 创建索引
CREATE INDEX `idx_request_log_user_uuid` ON `cs_request_log` (`user_uuid`) COMMENT '用户UUID索引';
CREATE INDEX `idx_request_log_request_time` ON `cs_request_log` (`request_time`) COMMENT '请求时间索引';
CREATE INDEX `idx_request_log_request_ip` ON `cs_request_log` (`request_ip`) COMMENT 'IP地址索引';
CREATE INDEX `idx_request_log_request_url` ON `cs_request_log` (`request_url`(255)) COMMENT '请求URL索引';

-- 添加外键约束
ALTER TABLE `cs_request_log`
    ADD CONSTRAINT `fk_request_log_user_uuid`
        FOREIGN KEY (`user_uuid`) REFERENCES `cs_user` (`user_uuid`)
            ON DELETE SET NULL ON UPDATE CASCADE;