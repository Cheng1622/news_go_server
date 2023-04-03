create table `user` (
  `id` bigint(20) not null auto_increment,
  `user_id` bigint(20) not null,
  -- utf8mb4_general_ci 这个就是排序的时候有点速度上的优势 仅此而已
  `username` varchar(64) collate utf8mb4_general_ci not null,
  `password` varchar(64) collate utf8mb4_general_ci not null,
  `email` varchar(64) collate utf8mb4_general_ci not null,
  -- `gender` tinyint(4) not null default '0',
  -- 这个就是插入数据的时候 可以自动设置时间
  `create_time` timestamp not null default current_timestamp,
  -- 更新数据饿的时候 时间戳也可以自动更新
  `update_time` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  -- 建立索引 加快存储速度 同时unique key 代表是唯一的 不能重复
  unique key `idx_email` (`email`) using btree,
  unique key `idx_user_id` (`user_id`) using btree
) default charset = utf8mb4 collate = utf8mb4_general_ci;

Drop table if exists `community`;

create table `community` (
  `id` int(11) not null auto_increment,
  -- 板块id
  `community_id` int(10) unsigned not null,
  -- 板块名称
  `community_name` varchar(255) collate utf8mb4_general_ci not null,
  -- 板块介绍
  `introduction` varchar(256) collate utf8mb4_general_ci not null,
  `create_time` timestamp not null default current_timestamp,
  `update_time` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  unique key `idx_community_id` (`community_id`),
  unique key `idx_community_name` (`community_name`)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci;

insert into community (community_id, community_name, introduction) values ('58' ,'国内新闻', 'xinwen');
insert into community (community_id, community_name, introduction) values ('263' ,'生猪期货', 'qihuo');
insert into community (community_id, community_name, introduction) values ('147' ,'国际新闻', 'guojixinwen');
insert into community (community_id, community_name, introduction) values ('148' ,'行业点评', 'xingyedianping');
insert into community (community_id, community_name, introduction) values ('149' ,'原创分析', 'dujiafenxi');
insert into community (community_id, community_name, introduction) values ('70' ,'每日猪评', 'zhuping');
insert into community (community_id, community_name, introduction) values ('118' ,'展会报道', 'baodao');
insert into community (community_id, community_name, introduction) values ('170' ,'种猪资讯', 'zhongzhuzixun');
insert into community (community_id, community_name, introduction) values ('166' ,'种猪行业新闻', 'zhongzhuxingyexinwen');
insert into community (community_id, community_name, introduction) values ('143' ,'种猪企业', 'zhongzhuqiye');
insert into community (community_id, community_name, introduction) values ('173' ,'种猪企业访谈', 'zhongzhujishu');
insert into community (community_id, community_name, introduction) values ('221' ,'名企推荐', 'mingqituijianzz');
insert into community (community_id, community_name, introduction) values ('31' ,'猪场建设', 'zhuchangjs');
insert into community (community_id, community_name, introduction) values ('32' ,'繁育管理', 'shoujing');
insert into community (community_id, community_name, introduction) values ('91' ,'饲养管理', 'siliaoyy');
insert into community (community_id, community_name, introduction) values ('35' ,'猪场管理', 'kxyangzhu');
insert into community (community_id, community_name, introduction) values ('233' ,'批次化生产', 'pxhshenchan');
insert into community (community_id, community_name, introduction) values ('261' ,'养猪大会', 'liman');
insert into community (community_id, community_name, introduction) values ('81' ,'行情分析', 'hangqingfenxi');
insert into community (community_id, community_name, introduction) values ('68' ,'玉米价格', 'yumi');
insert into community (community_id, community_name, introduction) values ('67' ,'豆粕价格', 'doupo');
insert into community (community_id, community_name, introduction) values ('257' ,'猪粮比', 'zhuliangbi');
insert into community (community_id, community_name, introduction) values ('256' ,'饲料供需', 'siliaogongxu');
insert into community (community_id, community_name, introduction) values ('267' ,'饲料分析', 'siliaofenxi');
insert into community (community_id, community_name, introduction) values ('63'	,' 生猪价格', 'shengzhu');
insert into community (community_id, community_name, introduction) values ('64' ,'仔猪价格', 'zizhu');
insert into community (community_id, community_name, introduction) values ('65' ,'猪肉价格', 'zhurou');
insert into community (community_id, community_name, introduction) values ('115' ,'各省市猪价', 'shengshi');
insert into community (community_id, community_name, introduction) values ('90' ,'养猪新闻', 'xinxi');
insert into community (community_id, community_name, introduction) values ('88' ,'养猪技术', 'jishushipin');
insert into community (community_id, community_name, introduction) values ('260' ,'每日猪价', 'meirizhujia');
insert into community (community_id, community_name, introduction) values ('92' ,'专家讲座', 'jiangzuo');
insert into community (community_id, community_name, introduction) values ('93' ,'人物访谈', 'fangtanshipin');
insert into community (community_id, community_name, introduction) values ('94' ,'企业展示', 'qiye');
insert into community (community_id, community_name, introduction) values ('111' ,'养猪致富', 'qitashipin');

drop table if exists `post`;

create table `post` (
  `id` bigint(20) not null auto_increment,
  `post_id` bigint(20) not null comment '帖子id',
  `title` varchar(255) collate utf8mb4_general_ci not null comment '标题',
  `content` TEXT collate utf8mb4_general_ci not null comment '内容',
  `author_id` bigint(20) not null comment '作者id',
  `isnews` tinyint(4)  default '0' comment '是否为新闻',
  `news_url` varchar(255) collate utf8mb4_general_ci  comment '标题',
  `news_source` varchar(255) collate utf8mb4_general_ci  comment '新闻来源',
  `news_time` DATETIME  default current_timestamp comment '新闻发布时间',
  `image1` varchar(1000) collate utf8mb4_general_ci  comment '图片1',
  `image2` varchar(1000) collate utf8mb4_general_ci  comment '图片2',
  `image3` varchar(1000) collate utf8mb4_general_ci  comment '图片3',
  `isimage` tinyint(4)  default '0' comment '是否有图片',
  `isimage3` tinyint(4)  default '0' comment '是否有3张图片',
  `videoimage` varchar(1000) collate utf8mb4_general_ci  comment '视频图片1',
  `video` varchar(255) collate utf8mb4_general_ci  comment '视频',
  `isvideo` tinyint(4)  default '0' comment '是否为视频',
  `community_id` bigint(20) not null default '1' comment '板块id',
  `status` tinyint(4) not null default '1' comment '帖子状态',
  `create_time` timestamp not null default current_timestamp comment '创建时间',
  `update_time` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
  primary key (`id`),
  -- 唯一索引
  unique key `idx_post_id` (`post_id`),
  -- 普通索引
  key `idx_author_id` (`author_id`),
  key `idx_community_id` (`community_id`)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci;


create table `list` (
  `id` int(11) not null auto_increment,
  `list_id` bigint(20) unsigned not null,
  `content` TEXT collate utf8mb4_general_ci not null  ,
  `create_time` timestamp not null default current_timestamp,
  `update_time` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  unique key `idx_list_id` (`list_id`)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci;
