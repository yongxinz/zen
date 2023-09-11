-- 工单流转

CREATE TABLE `wkf_circulation` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `ticket_id` int NOT NULL DEFAULT 0 COMMENT '工单ID',
  `state` varchar(255) NOT NULL DEFAULT '' COMMENT '状态信息',
  `source` varchar(255) NOT NULL DEFAULT '' COMMENT '源节点',
  `target` varchar(255) NOT NULL DEFAULT '' COMMENT '目标节点',
  `circulation` varchar(255) NOT NULL DEFAULT '' COMMENT '流转信息',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '流转状态，0拒绝，1同意，2其他',
  `handler_id` int NOT NULL DEFAULT 0 COMMENT '处理人ID',
  `handler_name` varchar(32) NOT NULL DEFAULT '' COMMENT '处理人姓名',
  `cost_duration` int NOT NULL DEFAULT 0 COMMENT '处理时长',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;