/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 04/06/2022 22:21:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dycomment
-- ----------------------------
DROP TABLE IF EXISTS `dycomment`;
CREATE TABLE `dycomment`  (
  `id` int(0) NOT NULL,
  `userid` int(0) NOT NULL,
  `videoid` int(0) NOT NULL,
  `content` varchar(140) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `date` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dycomment
-- ----------------------------
INSERT INTO `dycomment` VALUES (1, 2, 1, 'OK', '2022-05-23 20:12:05');

-- ----------------------------
-- Table structure for dyuser
-- ----------------------------
DROP TABLE IF EXISTS `dyuser`;
CREATE TABLE `dyuser`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `followcount` int(0) NOT NULL DEFAULT 0,
  `followercount` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dyuser
-- ----------------------------
INSERT INTO `dyuser` VALUES (1, 'admin', '95f003865b4cb5c93e0997b1cfeb9a77', 1, 1);
INSERT INTO `dyuser` VALUES (2, 'noa', '95f003865b4cb5c93e0997b1cfeb9a77', 0, 0);
INSERT INTO `dyuser` VALUES (4, 'hime', '95f003865b4cb5c93e0997b1cfeb9a77', 0, 0);
INSERT INTO `dyuser` VALUES (6, 'himesaka', '95f003865b4cb5c93e0997b1cfeb9a77', 0, 0);

-- ----------------------------
-- Table structure for dyvideo
-- ----------------------------
DROP TABLE IF EXISTS `dyvideo`;
CREATE TABLE `dyvideo`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `authorid` int(0) NULL DEFAULT NULL,
  `playurl` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `coverurl` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `favoritecount` int(0) NULL DEFAULT 0,
  `commentcount` int(0) NOT NULL DEFAULT 0,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `createtime` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dyvideo
-- ----------------------------
INSERT INTO `dyvideo` VALUES (1, 1, 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 5, 1, 'bear', 1653445662);
INSERT INTO `dyvideo` VALUES (6, 1, 'http://rcqzzkaz6.hd-bkt.clouddn.com/1654003460_bear.mp4', 'http://rcqzzkaz6.hd-bkt.clouddn.com/default_cover.png', 0, 0, 'bear', 1654003460);
INSERT INTO `dyvideo` VALUES (7, 6, 'http://rcqzzkaz6.hd-bkt.clouddn.com/1654068013_newbear.mp4', 'http://rcqzzkaz6.hd-bkt.clouddn.com/default_cover.png', 0, 0, 'newbear', 1654068013);
INSERT INTO `dyvideo` VALUES (8, 6, 'http://rcqzzkaz6.hd-bkt.clouddn.com/1654069934_AkemiHomura.mp4', 'http://rcqzzkaz6.hd-bkt.clouddn.com/default_cover.png', 0, 0, 'AkemiHomura', 1654069934);
INSERT INTO `dyvideo` VALUES (9, 6, 'http://rcqzzkaz6.hd-bkt.clouddn.com/1654317441_aa.mp4', 'http://rcqzzkaz6.hd-bkt.clouddn.com/default_cover.png', 0, 0, 'aa', 1654317441);
INSERT INTO `dyvideo` VALUES (10, 6, 'http://rcqzzkaz6.hd-bkt.clouddn.com/1654350134_aaabbbb.mp4', 'http://rcqzzkaz6.hd-bkt.clouddn.com/default_cover.png', 0, 0, 'aaabbbb', 1654350134);
INSERT INTO `dyvideo` VALUES (12, 6, 'http://rcqzzkaz6.hd-bkt.clouddn.com/1654351950_bca.mp4', 'http://rcqzzkaz6.hd-bkt.clouddn.com/default_cover.png', 0, 0, 'bca', 1654351950);
INSERT INTO `dyvideo` VALUES (13, 6, 'http://rcqzzkaz6.hd-bkt.clouddn.com/1654352380_ccaaca.mp4', 'http://rcqzzkaz6.hd-bkt.clouddn.com/default_cover.png', 0, 0, 'ccaaca', 1654352380);

SET FOREIGN_KEY_CHECKS = 1;
