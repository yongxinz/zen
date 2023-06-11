CREATE TABLE `sys_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `role_name` varchar(32) NOT NULL DEFAULT '' COMMENT '角色名称',
  `role_key` varchar(32) NOT NULL DEFAULT '' COMMENT '权限字符',
  `sort` tinyint NOT NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;