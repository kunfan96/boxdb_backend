/*
 Navicat Premium Dump SQL

 Source Server         : 虚拟机Mysql
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : 192.168.238.130:3306
 Source Schema         : boxdb

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 28/06/2025 14:50:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for boxdb_tenancy
-- ----------------------------
DROP TABLE IF EXISTS `boxdb_tenancy`;
CREATE TABLE `boxdb_tenancy` (
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '租户名字',
  `status` int DEFAULT NULL COMMENT '状态\n0:正常1:禁用',
  `create_user_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '创建用户id',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of boxdb_tenancy
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for boxdb_tenancy_user
-- ----------------------------
DROP TABLE IF EXISTS `boxdb_tenancy_user`;
CREATE TABLE `boxdb_tenancy_user` (
  `id` varchar(255) NOT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `tenancy_id` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of boxdb_tenancy_user
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for boxdb_users
-- ----------------------------
DROP TABLE IF EXISTS `boxdb_users`;
CREATE TABLE `boxdb_users` (
  `id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '主键 uuid',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名，唯一',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码，sha存储',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手机号，唯一',
  `role` int NOT NULL COMMENT '0：超级管理员（可以创建管理员和用户，创建租户）1：管理员（可以创建用户，创建租户）2：普通用户（不可以注册，只有<超级>管理员添加才可以登录）',
  `status` int NOT NULL COMMENT '0:禁用1:启用2:未启用',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of boxdb_users
-- ----------------------------
BEGIN;
INSERT INTO `boxdb_users` (`id`, `name`, `password`, `phone`, `role`, `status`, `create_time`, `update_time`) VALUES ('87c809dc-52dd-449e-bead-4acdb1b25d81', 'Admin', '$2a$10$rkYq8OZELcJNjz0GzFucEegZ2ZL5KRFTAWI0F80ksPLfxo8XGKM5q', '15876870780', 0, 1, '2025-06-17 08:06:13', '2025-06-17 08:20:49');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
