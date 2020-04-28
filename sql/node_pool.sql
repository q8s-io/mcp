/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : mcp

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 28/04/2020 10:02:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for node_pool
-- ----------------------------
DROP TABLE IF EXISTS `node_pool`;
CREATE TABLE `node_pool` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `cluster_conf_id` int(10) unsigned DEFAULT NULL,
  `pool_role` tinyint(3) unsigned NOT NULL,
  `pool_name` varchar(20) NOT NULL,
  `instance_type` varchar(30) NOT NULL,
  `instance_size_gb` smallint(5) unsigned NOT NULL,
  `replicas` smallint(5) unsigned NOT NULL,
  `os_image` varchar(20) NOT NULL,
  `os_disk_type` varchar(30) NOT NULL,
  `os_disk_size_gb` smallint(5) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_node_pool_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
