CREATE TABLE `wkf_process` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
  `icon` varchar(128) NOT NULL DEFAULT '' COMMENT 'ICON',
  `structure` text NOT NULL DEFAULT '' COMMENT '流程结构',
  `classify` int NOT NULL DEFAULT 0 COMMENT '分类',
  `template` varchar(255) NOT NULL DEFAULT '' COMMENT '模板',
  `task` varchar(255) NOT NULL DEFAULT '' COMMENT '任务',
  `notice` varchar(255) NOT NULL DEFAULT '' COMMENT '通知',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;