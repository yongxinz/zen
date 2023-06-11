-- rms.sys_user definition

CREATE TABLE `sys_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `role_id` int NOT NULL DEFAULT 0 COMMENT '角色ID',
  `avatar` varchar(512) NOT NULL DEFAULT '' COMMENT '头像',
  `sex` tinyint NOT NULL DEFAULT 0 COMMENT '性别',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `dept_id` int NOT NULL DEFAULT 0 COMMENT '部门',
  `post_id` int NOT NULL DEFAULT 0 COMMENT '岗位',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_sys_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
