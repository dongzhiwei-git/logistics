CREATE TABLE `common_user`
(
    `id`        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`   bigint(20) unsigned NOT NULL COMMENT '用户ID',
    `user_name` varchar(100)        NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar`    varchar(200)        NOT NULL DEFAULT '' COMMENT '头像',
    `password`  varchar(60)         NOT NULL DEFAULT '' COMMENT '密码',
    `tel`       varchar(60)         NOT NULL DEFAULT '' COMMENT '电话',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='普通用户表';


CREATE TABLE `fans_user`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`      bigint(20) unsigned NOT NULL COMMENT '用户ID',
    `nickname`     varchar(100)        NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar`       varchar(200)        NOT NULL DEFAULT '' COMMENT '头像',
    `union_id`     varchar(60)         NOT NULL DEFAULT '' COMMENT '微信唯一ID',
    `xcx_openid`   varchar(60)         NOT NULL DEFAULT '' COMMENT '小程序openid',
    `gzh_openid`   varchar(60)         NOT NULL DEFAULT '' COMMENT '某一个公众号的openid',
    `has_sub`      tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '已订阅数',
    `sub_top`      int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '订阅上限',
    `expire_date`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
    `created_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

CREATE TABLE `student_set`
(
    `sno`   char(8)  NOT NULL COMMENT '学号',
    `sname` char(8)  NOT NULL COMMENT '姓名',
    `age`   smallint COMMENT '年龄',
    `sex`   char(2)  NOT NULL COMMENT '性别',
    `sdept` char(50) NOT NULL COMMENT '所在系',
    PRIMARY KEY (`sno`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4 COMMENT ='学生信息表';

CREATE TABLE `college_class`
(
    `cno`      char(4)  NOT NULL COMMENT '课程号',
    `cname`    char(20) NOT NULL COMMENT '课程名',
    `credit`   float COMMENT '学分',
    `pcno`     char(4) COMMENT '先修课',
    `describe` varchar(100) COMMENT '课程描述',
    PRIMARY KEY (`cno`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4 COMMENT ='学生课程表';

CREATE TABLE `select_class`
(
    `sno`   char(8) NOT NULL COMMENT '学号',
    `cno`   char(4) NOT NULL COMMENT '课程号',
    `grade` float COMMENT '成绩',
    PRIMARY KEY (`sno`),
    FOREIGN KEY (`sno`) REFERENCES student_set (`sno`),
    FOREIGN KEY (`cno`) REFERENCES college_class (`cno`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4 COMMENT ='学生成绩表';


CREATE TABLE `need`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `number`      varchar(30)         NOT NULL DEFAULT '' COMMENT '编号',
    `need_class`  char(4)             NOT NULL DEFAULT '' COMMENT '需求类别',
    `need_detail` varchar(100)        NOT NULL DEFAULT '' COMMENT '需求细节',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_number` (`number`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='需求表';

CREATE TABLE `store_to_center`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `number`       varchar(30)         NOT NULL DEFAULT '' COMMENT '编号',
    `user_name`    varchar(30)         NOT NULL DEFAULT '' COMMENT '客户名称',
    `product_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '产品名称',
    `need_sum`     bigint(20)          NOT NULL COMMENT '需求数量',
    `product_type` char(6)             NOT NULL DEFAULT '' COMMENT '需求类型',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_number` (`number`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='总仓到分仓表';

CREATE TABLE `center_to_customer`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `number`       varchar(30)         NOT NULL DEFAULT '' COMMENT '编号',
    `area`         varchar(30)         NOT NULL DEFAULT '' COMMENT '所属区域',
    `product_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '产品名称',
    `need_sum`     bigint(20)          NOT NULL COMMENT '需求数量',
    `product_type` char(6)             NOT NULL DEFAULT '' COMMENT '需求类型',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_number` (`number`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='分仓到客户';

# 入库表
CREATE TABLE `purchase`

(
    `id`             bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `area`           varchar(30)         NOT NULL DEFAULT '' COMMENT '所属区域',
    `product_name`   varchar(20)         NOT NULL DEFAULT '' COMMENT '备件名称',
    `order_sum`      int unsigned        NOT NULL COMMENT '订货数量',
    `left_sum`       int unsigned        NOT NULL COMMENT '剩余库存',
    `purchase_level` varchar(6)          NOT NULL DEFAULT '' COMMENT '采购预警',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='入库表';

# 出库表
CREATE TABLE `out`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `product_name`      varchar(30)         NOT NULL DEFAULT '' COMMENT '配送备件',
    `order_num`         varchar(30)         NOT NULL DEFAULT '' COMMENT '订货编号',
    `order_sum`         int unsigned        NOT NULL COMMENT '配送数量',
    `order_customer`    varchar(30)         NOT NULL DEFAULT '' COMMENT '订货客户',
    `order_customer_id` varchar(20)         NOT NULL DEFAULT '' COMMENT '客户ID',
    `to_area`           varchar(20)         NOT NULL DEFAULT '' COMMENT '配送地区',
    `arrive_time`       varchar(20)         NOT NULL DEFAULT '' COMMENT '抵达时间',
    `trade_status`      varchar(6)          NOT NULL DEFAULT '' COMMENT '交易状态',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_num` (`order_num`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='出库表';

#  车辆表
CREATE TABLE `vehicle`
(
    `id`                 bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `vehicle_num`        varchar(30)         NOT NULL DEFAULT '' COMMENT '车辆编号',
    `vi_sever_condition` varchar(30)         NOT NULL DEFAULT '' COMMENT '车辆使用情况',
    `vi_position`        varchar(30)         NOT NULL DEFAULT '' COMMENT '所处位置',
    `vi_load_weight`     varchar(30)         NOT NULL DEFAULT '' COMMENT '载货量',
    `vi_transit_value`   varchar(30)         NOT NULL DEFAULT '' COMMENT '运输成本',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_vehicle_num` (`vehicle_num`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='车辆表';


# 配送表
CREATE TABLE `distribution`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `position`     varchar(30)         NOT NULL DEFAULT '' COMMENT '所处位置',
    `product_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '备件名称',
    `product_sum`  int unsigned        NOT NULL COMMENT '库存数量',
    `product_type` varchar(6)          NOT NULL DEFAULT '' COMMENT '备件类型',
    `user_name`    varchar(30)         NOT NULL DEFAULT '' COMMENT '负责人人名称',
    `tel`          varchar(20)         NOT NULL DEFAULT '' COMMENT '联系电话',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='配送表';

#总仓信息表
CREATE TABLE `warehouse`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `product_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '备件名称',
    `stock_sum`    int unsigned        NOT NULL COMMENT '库存数量',
    `user_name`    varchar(30)         NOT NULL DEFAULT '' COMMENT '负责人人名称',
    `tel`          varchar(20)         NOT NULL DEFAULT '' COMMENT '联系电话',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='总仓信息表';

# 客户信息表
CREATE TABLE `customer`
(
    `id`               bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_name`        varchar(30)         NOT NULL DEFAULT '' COMMENT '客户名称',
    `order_num`        varchar(30)         NOT NULL DEFAULT '' COMMENT '订单编号',
    `product_name`     varchar(30)         NOT NULL DEFAULT '' COMMENT '备件名称',
    `need_sum`         int unsigned        NOT NULL COMMENT '所需数量',
    `product_type`     varchar(6)          NOT NULL DEFAULT '' COMMENT '备件类型',
    `position`         varchar(30)         NOT NULL DEFAULT '' COMMENT '所处位置',
    `responsible_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '负责人',
    `tel`              varchar(20)         NOT NULL DEFAULT '' COMMENT '联系电话',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='客户信息表';

# 采购端信息表
CREATE TABLE `procurement`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `position`     varchar(30)         NOT NULL DEFAULT '' COMMENT '所在区域',
    `product_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '备件名称',
    `order_sum`    int unsigned        NOT NULL COMMENT '订货量',
    `left_sum`     int unsigned        NOT NULL COMMENT '剩余库存',
    `product_type` varchar(6)          NOT NULL DEFAULT '' COMMENT '备件类型',
    `need_level`   varchar(8)          NOT NULL DEFAULT '' COMMENT '采购预警',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='采购端信息表';

# 库存信息表
CREATE TABLE `stock`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `product_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '备件名称',
    `left_sum`     int unsigned        NOT NULL COMMENT '剩余库存',
    `product_type` varchar(6)          NOT NULL DEFAULT '' COMMENT '备件类型',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='库存信息表';

# 配送车辆表
CREATE TABLE `vehicle`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `distribution_num`  varchar(30)         NOT NULL DEFAULT '' COMMENT '配送编号',
    `vehicle_num`       varchar(30)         NOT NULL DEFAULT '' COMMENT '车辆编号',
    `customer_name`     varchar(30)         NOT NULL DEFAULT '' COMMENT '订单客户名',
    `distribution_area` varchar(30)         NOT NULL DEFAULT '' COMMENT '配送地点',
    `product_name`      varchar(30)         NOT NULL DEFAULT '' COMMENT '采购名称',
    `product_sum`       int unsigned        NOT NULL COMMENT '备件数量',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='配送车辆表';

# 备件表
CREATE TABLE `product`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `product_num`  varchar(30)         NOT NULL DEFAULT '' COMMENT '备件编号',
    `product_name` varchar(30)         NOT NULL DEFAULT '' COMMENT '备件名称',
    `product_type` varchar(6)          NOT NULL DEFAULT '' COMMENT '备件类型',
    `product_sum`  int unsigned        NOT NULL COMMENT '备件数量',
    `in_date`      time                NOT NULL   COMMENT '入库日期',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='备件表';