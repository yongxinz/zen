CREATE TABLE `wkf_task` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
  `category` varchar(16) NOT NULL DEFAULT '' COMMENT '类型',
  `content` text NOT NULL DEFAULT '' COMMENT '内容',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;