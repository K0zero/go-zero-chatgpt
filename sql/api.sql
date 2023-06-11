/*
 Navicat MySQL Data Transfer

 Source Server         : looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : looklook_order

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/03/2022 17:15:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for openai
-- ----------------------------
DROP TABLE IF EXISTS `openai`;
CREATE TABLE `openai` (
                                  `id` bigint NOT NULL AUTO_INCREMENT,
                                  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `del_state` tinyint NOT NULL DEFAULT '0',
                                  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                                  `api` char(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
                                  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户id',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `idx_sn` (`api`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='OPENAI API';

SET FOREIGN_KEY_CHECKS = 1;
