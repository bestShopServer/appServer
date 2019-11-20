-- 用户类型信息表
drop table if exists user_type;
create table user_type
(
    type_id   tinyint auto_increment comment '用户类型ID',
    type_name char(128) default '' comment '用户类型名称',
    thumb     char(512) default '' comment '缩略图',
    info_time datetime  default current_timestamp comment '登记时间',
    primary key (type_id)
);


-- 用户基本信息表
drop table if exists user_base;
create table user_base
(
    user_id      integer auto_increment comment '用户ID',
    user_name    char(128)    default '' comment '用户名称',
    user_type_id tinyint      default 0 comment '用户类型',
    phone        char(20)     default '' comment '手机号',
    open_id      char(60)     default '' comment '小程序用户ID',
    union_id     char(60)     default '' comment '小程序唯一ID',
    nick_name    varchar(60)  default '' comment '昵称',
    head         varchar(128) default '' comment '头像',
    passwd       char(128)    default '' comment '登录密码',
    integral     int          default 0 comment '用户积分',
    last_time    datetime     default current_timestamp comment '最后时间',
    info_time    datetime     default current_timestamp comment '登记时间',
    primary key (user_id)
);
create unique index user_base_idx1 on user_base (phone);


-- 用户账户余额信息表
drop table if exists user_account;
create table user_account
(
    user_id   integer auto_increment comment '用户ID',
    bal       decimal(15, 4) default 0.00 comment '账户余额',
    info_time datetime       default current_timestamp comment '登记时间',
    primary key (user_id)
);


-- 用户账户余额明细信息表
drop table if exists user_account_list;
create table user_account_list
(
    user_id   integer auto_increment comment '用户ID',
    bal       decimal(15, 4) default 0.00 comment '账户余额',
    amt       decimal(15, 4) default 0.00 comment '消费金额负数是消费',
    info_time datetime       default current_timestamp comment '登记时间',
    primary key (user_id)
);

-- 收货地址表
drop table if exists user_address;
create table user_address
(
    addr_id   integer auto_increment comment '收货地址ID',
    user_id   integer   default 0 comment '用户ID',
    city_code char(10)  default '' comment '城市代码',
    address   char(128) default '' comment '收货地址',
    mobile    char(20)  default '' comment '收货手机号',
    info_time datetime  default current_timestamp comment '登记时间',
    primary key (user_id)
);

-- 商品分类信息表
drop table if exists goods_type;
create table goods_type
(
    goods_type_id   integer auto_increment comment '商品ID',
    goods_type_name char(128) default '' comment '商品名称',
    thumb           char(512) default '' comment '缩略图',
    info_time       datetime  default current_timestamp comment '登记时间',
    primary key (goods_type_id)
);


-- 商品基本信息表
drop table if exists goods_base;
create table goods_base
(
    goods_id      integer auto_increment comment '商品ID',
    goods_name    char(128)      default '' comment '商品名称',
    goods_type_id tinyint        default 0 comment '商品类型',
    user_type_id  tinyint        default 0 comment '用户类型-可查看',
    head          varchar(1024)  default '' comment '商品头像,展示首页',
    thumb         varchar(1024)  default '' comment '商品缩略图,多张用|分隔',
    price         decimal(15, 4) default 0.00 comment '定价',
    discount      decimal(10, 6) default 1.00 comment '折扣',
    sale_price    decimal(15, 4) default 0.00 comment '售价',
    details       text not null comment '商品详情',
    specs         text not null comment '商品规格',
    service        text not null comment '商品服务',
    remark        varchar(1024)  default '' comment '备注',
    info_time     datetime       default current_timestamp comment '登记时间',
    primary key (goods_id)
);


-- 商品订单表
drop table if exists goods_order;
create table goods_order
(
    goods_id     integer auto_increment comment '商品ID',
    out_trade_no char(128)      default '' comment '商户订单号',
    card_no      char(60)       default '' comment '消费卡号',
    amt          decimal(15, 4) default 0.00 comment '订单金额',
    num          tinyint        default 0 comment '商品数量',
    unit_amt     decimal(15, 4) default 0.00 comment '生成订单时商品单价',
    remark       varchar(1024)  default '' comment '备注',
    info_time    datetime       default current_timestamp comment '登记时间',
    primary key (goods_id)
);


-- 用户购物车表
drop table if exists user_address;
create table user_address
(
    cart_id   integer auto_increment comment '购物车ID',
    user_id   integer        default 0 comment '用户ID',
    goods_id  integer        default 0 comment '商品ID',
    addr_id   integer        default 0 comment '收货地址ID',
    amt       decimal(15, 4) default 0.00 comment '订单金额',
    num       tinyint        default 0 comment '商品数量',
    unit_amt  decimal(15, 4) default 0.00 comment '生成订单时商品单价',
    info_time datetime       default current_timestamp comment '登记时间',
    primary key (user_id)
);


-- 商品微信付款回调表
drop table if exists user_pay;
create table user_pay
(
    pay_id              int auto_increment,
    out_trade_no        char(32)       default '' comment '商户订单号',
    pay_amt             decimal(15, 2) default 0.00 comment '实际付款金额',
    prepay_id           char(64)       default '' comment '预支付交易会话标识',
    code_url            char(64)       default '' comment '二维码链接',
    total_fee           int            default 0 comment '订单总金额(分)',
    return_code         char(16)       default '' comment '状态码',
    return_msg          char(128)      default '' comment '返回信息',
    result_code         char(16)       default '' comment '状态码',
    err_code            char(32)       default '' comment '业务错误信息',
    err_code_des        char(128)      default '' comment '业务错误描述',
    content             text not null comment '发送内容',
    pay_content         text not null comment '小程序支付报文',
    notify_return_code  char(16)       default '' comment '回调状态码',
    notify_return_msg   char(128)      default '' comment '回调返回信息',
    notify_result_code  char(16)       default '' comment '回调业务状态码',
    notify_err_code     char(32)       default '' comment '回调业务错误信息',
    notify_err_code_des char(128)      default '' comment '回调业务错误描述',
    transaction_id      char(128)      default '' comment '微信支付订单号',
    time_end            char(14)       default '' comment '支付完成时间',
    status              int            default 0 comment '状态 0-正常 1-删除 2-已发货',
    remark              varchar(1024)  default '' comment '备注',
    infotime            datetime       default '2006-01-02 15:03:04' comment '登记时间',
    primary key (pay_id)
);
create unique index user_pay_idx1 on user_pay (out_trade_no);


-- 卡片信息表
drop table if exists card_base;
create table card_base
(
    card_id     integer auto_increment comment '卡片ID',
    card_no     char(32)       default '' comment '卡号',
    card_passwd char(128)      default '' comment '卡密',
    amt         decimal(15, 4) default 0.00 comment '金额',
    info_time   datetime       default current_timestamp comment '登记时间',
    primary key (card_id)
);

