CREATE TABLE `wkf_template` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `name` varchar(64) UNIQUE NOT NULL DEFAULT '' COMMENT '名称',
  `form_structure` text NOT NULL DEFAULT '' COMMENT '模板结构',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;