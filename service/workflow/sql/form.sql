-- 工单表单数据

CREATE TABLE `wkf_form` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `ticket_id` int NOT NULL DEFAULT 0 COMMENT '工单ID',
  `form_structure` text NOT NULL DEFAULT '' COMMENT '表单结构',
  `form_data` text NOT NULL DEFAULT '' COMMENT '表单数据',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;