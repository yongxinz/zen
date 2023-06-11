CREATE TABLE `sys_post` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `post_name` varchar(64) NOT NULL DEFAULT '' COMMENT '岗位名称',
  `post_code` varchar(16) NOT NULL DEFAULT '' COMMENT '岗位编码',
  `sort` tinyint NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;