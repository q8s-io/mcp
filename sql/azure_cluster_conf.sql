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

 Date: 28/04/2020 10:01:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for azure_cluster_conf
-- ----------------------------
DROP TABLE IF EXISTS `azure_cluster_conf`;
CREATE TABLE `azure_cluster_conf` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  `region` varchar(20) NOT NULL,
  `cluster_version` varchar(20) NOT NULL,
  `resource_group_name` varchar(20) NOT NULL,
  `virtual_network_name` varchar(20) NOT NULL,
  `virtual_network_address` varchar(20) NOT NULL,
  `subnet_name` varchar(20) NOT NULL,
  `subnet_address` varchar(20) NOT NULL,
  `secret_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_azure_cluster_conf_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
