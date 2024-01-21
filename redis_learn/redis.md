### Redis 从零到入门

#### 1. Redis 简介

- Redis 是什么
- Redis 的主要特性：内存存储、持久化、网络通信、数据类型丰富
- Redis 与其他数据库的对比
- Redis 的应用场景与案例

#### 2. 安装和配置 Redis

- 在 Linux 环境下的安装步骤
- Redis 配置文件详解（redis.conf）
- 启动、停止及重启 Redis 服务
- Redis 客户端工具使用

#### 3. Redis 基本操作

- 连接 Redis 服务器
- Redis 命令行基础
- String 类型的基本操作：SET, GET, INCR, MGET 等
- List、Set、Sorted Set、Hash 等复杂数据类型的介绍与操作
- Key 过期时间设置与管理
- 数据持久化策略：RDB (Redis Database) 快照方式和 AOF (Append Only File) 日志方式

#### 4. Redis 进阶内容

- Key 管理：DEL/TTL/EXPIRE/RENAMENX 等
- 数据结构操作：UNION、INTER、DIFF 等
- 模糊查询：KEYS/PATTERN MATCHING
- 消息订阅与发布（Pub/Sub）
- 事务处理（MULTI/EXEC/DISCARD/WATCH）
- EVAL/EVALSHA/LUA 脚本在 Redis 中的应用
- Redis 集群搭建与原理
- 分布式锁实现
- Redis Sentinel 高可用解决方案
- Redis 缓存策略设计

#### 5.Redis 高可用与集群方案

- 主从复制（Replication）原理与配置
- Sentinel 哨兵模式实现自动故障转移
- Redis Cluster 集群搭建与使用
- 分片（Sharding）策略与实践

#### 6. Redis 性能优化与运维实践

- 内存优化与监控
- 大 key 问题分析与处理
- Redis 主从复制机制及故障恢复
- 使用 Redis 缓存穿透、缓存雪崩等问题及其解决方案
- 性能调优与压力测试

### Redis 高级主题

- Redis 模块开发
- Redis Streams 的使用与实战
- Redis6 的新特性（如 ACL 权限控制、多线程 IO 模型等）
- Redis 在微服务架构中的应用模式
- 结合实际业务场景进行案例分析与设计
