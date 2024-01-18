SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `btf`;
CREATE DATABASE IF NOT EXISTS `btf` DEFAULT CHARACTER SET utf8mb4;
USE `btf`;

-- ----------------------------
-- Table structure for t_user_info
-- ----------------------------
DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info`  (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
     `avatar` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像地址',
     `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '邮箱',
     `intro` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '个人简介',
     `website` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '个人网站',
     `is_disable` tinyint(1) NULL DEFAULT NULL COMMENT '是否禁用(0-否 1-是)',
     `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `t_user_info` (`id`, `nickname`, `avatar`, `email`, `intro`, `website`, `is_disable`)
VALUES (1, '管理员', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', 'admin@qq.com', '我是管理员用户！', 'https://www.raein11.top', 0);

INSERT INTO `t_user_info` (`id`, `nickname`, `avatar`, `email`, `intro`, `website`, `is_disable`)
VALUES (2, '普通用户', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', 'user@qq.com', '我是普通用户！', 'https://www.raein11.top', 0);

INSERT INTO `t_user_info` (`id`, `nickname`, `avatar`, `email`, `intro`, `website`, `is_disable`)
VALUES (3, '测试用户', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', 'test@qq.com', '我是测试用户！', 'https://www.raein11.top', 0);

-- ----------------------------
-- Table structure for t_user_auth
-- ----------------------------
DROP TABLE IF EXISTS `t_user_auth`;
CREATE TABLE `t_user_auth`  (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `user_info_id` bigint NULL DEFAULT NULL COMMENT '用户信息ID',
     `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户名',
     `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '密码',
     `login_type` tinyint(1) NULL DEFAULT NULL COMMENT '登录类型',
     `last_login_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上次登录时间',
     `ip_address` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '登录IP地址',
     `ip_source` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT 'IP来源',
     `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 245 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `t_user_auth` (`id`, `user_info_id`, `username`, `password`, `login_type`, `ip_address`, `ip_source`)
VALUES (1, 1, 'admin', '$2a$10$np.P54Jep7GB/H5vG1PcbudYcxAAf1iiBf7NzsQJT9ZfsYz6tFPcm', 1, '192.168.50.31', '内网IP');
INSERT INTO `t_user_auth` (`id`, `user_info_id`, `username`, `password`, `login_type`, `ip_address`, `ip_source`)
VALUES (2, 2, 'user', '$2a$10$9vHpoeT7sF4j9beiZfPsOe0jJ67gOceO2WKJzJtHRZCjNJajl7Fhq', 1, '172.21.242.1:48716', '');
INSERT INTO `t_user_auth` (`id`, `user_info_id`, `username`, `password`, `login_type`, `ip_address`, `ip_source`)
VALUES (3, 3, 'test', '$2a$10$FmU4jxwDlibSL9pdt.AsuODkbB4gLp3IyyXeoMmW/XALtT/HdwTsi', 1, '192.168.50.31', '内网IP');

-- ----------------------------
-- Table structure for t_user_role
-- ----------------------------
DROP TABLE IF EXISTS `t_user_role`;
CREATE TABLE `t_user_role`  (
     `user_id` bigint NULL DEFAULT NULL,
     `role_id` bigint NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `t_user_role` VALUES (1, 1);
INSERT INTO `t_user_role` VALUES (2, 2);
INSERT INTO `t_user_role` VALUES (3, 3);

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role`  (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '角色名',
     `label` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '角色描述',
     `is_disable` tinyint(1) NULL DEFAULT NULL COMMENT '是否禁用(0-否 1-是)',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `t_role` VALUES (1, '管理员', 'admin', 0);
INSERT INTO `t_role` VALUES (2, '用户', 'user', 0);
INSERT INTO `t_role` VALUES (3, '测试', 'test', 0);


SET FOREIGN_KEY_CHECKS = 1;