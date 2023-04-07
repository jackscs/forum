create table `community`(
                            `id` int(11) not null auto_increment,
                            `community_id` int(10) unsigned not null,
                            `community_name` varchar(128) collate utf8mb4_general_ci not null,
                            `introduction` varchar(256) collate utf8mb4_general_ci not null,
                            `create_time` timestamp not null default  current_timestamp,
                            `update_time` timestamp not null default  current_timestamp on update current_timestamp,
                            primary key (`id`),
                            unique key `idx_community_id` (`community_id`),
                            unique key `idx_community_name` (`community_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 collate utf8mb4_general_ci

insert into `community` values('1','1','Go','Golang','2016-11-01 08:10:10','2016-11-01 08:10:10')
insert into `community` values('2','2','Leetcode','','2016-11-01 08:10:10','2016-11-01 08:10:10')
insert into `community` values('3','3','CS:Go','Rush B','2016-11-01 08:10:10','2016-11-01 08:10:10')
insert into `community` values('4','4','Lol','!','2016-11-01 08:10:10','2016-11-01 08:10:10'


create table `article`(
                          `id` bigint(20) not null auto_increment,
                          `post_id` bigint(20) not null comment '帖子id',
                          `title` varchar(128) collate utf8mb4_general_ci not null comment '内容',
                          `content` varchar(8192) collate utf8mb4_general_ci not null comment '内容',
                          `author_id` bigint(20) not null comment '作者的用户id',
                          `community_id` bigint(20) not null comment '所属社区',
                          `status` tinyint(4) not null default '1' comment '帖子状态',
                          `create_time` timestamp null  default current_timestamp comment '创建时间',
                          `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                          primary key (`id`),
                          unique key `idx_post_id` (`post_id`),
                          key `idx_author_id` (`author_id`),
                          key `idx_community_id`(`community_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 collate=utf8mb4_general_ci;










