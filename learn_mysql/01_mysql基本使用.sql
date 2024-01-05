# 库操作
-- create database learnsql default character set utf8 collate utf8_general_ci;

/* 
-- 查看库
show databases;
-- 进入库
use learnsql;
-- 查看库里面的表
show tables; 
-- 查看库的编码，有的编码写中文会报错，这个要注意
show create database learnsql;
-- 删除数据库
drop database learnsql;
-- 修改库的编码
alter database learnsql character set utf8;

 */

# 表操作 (前提必须是进入到某个库里面)

-- 建表

/* create table user(
  id         int(11) primary key auto_increment not null comment '序号',                         # 主键，自动增长
  username   varchar(50) unique                 not null comment '姓名',                         # 不能为空,不能重复
  sex        boolean   default true comment '性别 1 男 0 女',                                     # tinyint(1) 一位 默认值是1
  age        int(8)    default 18 comment '年龄，默认18',                                         # tinyint(1) 一位 默认值是1
  created_at timestamp default CURRENT_TIMESTAMP on update  CURRENT_TIMESTAMP comment '创建时间'  # 创建时间,时间自动创建
); */

/* 
字段约束
常见的字段约束

primary Key  # 记录的唯一标识，能够通过该标识确定唯一一条记录  一般用于记录id
default # 默认值
unique # 唯一约束
not null # 空值约束，不可以为空，不写就是可以为空
auto_increment # 自动增长

*/

/* 
表操作

-- 查看表结构
desc user;

*/


desc user;