/*
 Navicat Premium Data Transfer

 Source Server         : testdb
 Source Server Type    : MySQL
 Source Server Version : 100420
 Source Host           : 127.0.0.1:3306
 Source Schema         : user_manag

 Target Server Type    : MySQL
 Target Server Version : 100420
 File Encoding         : 65001

 Date: 15/08/2021 21:43:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 118 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'username_afina', '$2a$10$09wICJT0tsTxG0hDsVwwb.6LUq9atxY.8nWioI4k7dSrZZ9BXUi2.', 'name_afina');
INSERT INTO `users` VALUES (8, 'username_alna', '$2a$10$..sJrPhcZ/Q/hH1lzsaVYeu6qtMBVY6fkPmG6CKL2r3qsGk78PZAa', 'name_ana1');
INSERT INTO `users` VALUES (9, 'username_rinda', '$2a$10$1bQTx9L2Ze7wPzXErsud8e..WApQCYOcKUjpwoyY3VtJbOTZe2mjO', 'name_rinda');
INSERT INTO `users` VALUES (15, 'username_ana', '$2a$10$1bQTx9L2Ze7wPzXErsud8e..WApQCYOcKUjpwoyY3VtJbOTZe2mjO', 'name_ana2');

SET FOREIGN_KEY_CHECKS = 1;
