-- 任务执行表

CREATE TABLE `wkf_execution` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `task_id` int NOT NULL DEFAULT 0 COMMENT '任务ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '任务名称',
  `category` varchar(16) NOT NULL DEFAULT '' COMMENT '任务类型',
  `result` text NOT NULL DEFAULT '' COMMENT '执行结果',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;