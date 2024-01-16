### Redis从零到入门

#### 1. Redis简介
- Redis是什么
- Redis的主要特性：内存存储、持久化、网络通信、数据类型丰富
- Redis与其他数据库的对比
- Redis的应用场景与案例

#### 2. 安装和配置Redis
- 在Linux环境下的安装步骤
- Redis配置文件详解（redis.conf）
- 启动、停止及重启Redis服务
- Redis客户端工具使用

#### 3. Redis基本操作
- 连接Redis服务器
- Redis命令行基础
- String类型的基本操作：SET, GET, INCR, MGET等
- List、Set、Sorted Set、Hash等复杂数据类型的介绍与操作
- Key过期时间设置与管理
- 数据持久化策略：RDB (Redis Database) 快照方式和AOF (Append Only File) 日志方式

#### 4. Redis进阶内容
- Key管理：DEL/TTL/EXPIRE/RENAMENX等
- 数据结构操作：UNION、INTER、DIFF等
- 模糊查询：KEYS/PATTERN MATCHING
- 消息订阅与发布（Pub/Sub）
- 事务处理（MULTI/EXEC/DISCARD/WATCH）
- EVAL/EVALSHA/LUA脚本在Redis中的应用
- Redis集群搭建与原理
- 分布式锁实现
- Redis Sentinel高可用解决方案
- Redis缓存策略设计

#### 5.Redis高可用与集群方案
- 主从复制（Replication）原理与配置
- Sentinel哨兵模式实现自动故障转移
- Redis Cluster集群搭建与使用
- 分片（Sharding）策略与实践

#### 6. Redis性能优化与运维实践
- 内存优化与监控
- 大key问题分析与处理
- Redis主从复制机制及故障恢复
- 使用Redis缓存穿透、缓存雪崩等问题及其解决方案
- 性能调优与压力测试

### Redis高级主题
- Redis模块开发
- Redis Streams的使用与实战
- Redis6的新特性（如ACL权限控制、多线程IO模型等）
- Redis在微服务架构中的应用模式
- 结合实际业务场景进行案例分析与设计