/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Version : 50712
 Source Host           : localhost
 Source Database       : ppgo_api_demo_gin

 Target Server Version : 50712
 File Encoding         : utf-8

 Date: 10/18/2017 22:21:17 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `ppgo_member`
-- ----------------------------
DROP TABLE IF EXISTS `ppgo_member`;
CREATE TABLE `ppgo_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login_name` varchar(50) NOT NULL DEFAULT '0' COMMENT '登录名',
  `password` varchar(64) NOT NULL DEFAULT '0' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `ppgo_member`
-- ----------------------------
BEGIN;
INSERT INTO `ppgo_member` VALUES ('1', 'haodaquan', '1234'), ('3', 'hell31', 'g2223'), ('4', 'hell31', 'g2223'), ('5', 'hell31', 'g2223'), ('6', 'hell31', 'g2223');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
