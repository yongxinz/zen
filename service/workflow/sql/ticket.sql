-- 工单信息

CREATE TABLE `wkf_ticket` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `classify_id` int NOT NULL DEFAULT 0 COMMENT '分类',
  `process_id` int NOT NULL DEFAULT 0 COMMENT '流程',
  `is_end` tinyint NOT NULL DEFAULT 0 COMMENT '是否结束，0未结束，1已结束',
  `is_denied` tinyint NOT NULL DEFAULT 0 COMMENT '是否拒绝，0否，1是',
  `state` varchar(255) NOT NULL DEFAULT '' COMMENT '状态信息',
  `related_person` varchar(255) NOT NULL DEFAULT '' COMMENT '关联人员',
  `urge_count` int NOT NULL DEFAULT 0 COMMENT '催办次数',
  `urge_lasttime` int NOT NULL DEFAULT 0 COMMENT '上次催办时间',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;