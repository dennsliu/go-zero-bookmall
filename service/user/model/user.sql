CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `number` varchar(255) NOT NULL DEFAULT '' COMMENT '学号',
    `username` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户名称',
    `email` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender` char(5)  NOT NULL COMMENT '男｜女｜未公开',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `number_unique` (`number`)
    UNIQUE KEY `number_unique` (`username`)
    UNIQUE KEY `number_unique` (`email`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;