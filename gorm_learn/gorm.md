### GORM从零到入门

#### 1. GORM简介
- GORM是什么：Go语言中的ORM库
- GORM的主要特性与优势
- Go语言与数据库操作对比
- GORM支持的数据源（MySQL、PostgreSQL、SQLite等）

#### 2. 安装与环境配置
- 安装GORM库及对应的数据库驱动包
- 配置连接数据库的基本信息

#### 3. GORM基础操作
- 定义结构体模型与数据库表映射
- 基本CRUD操作（Create, Read, Update, Delete）
- 查询数据（Find, First, Last, Where, Preload等）
- 数据分页与排序

#### 4. 关联关系处理
- 一对一关联 (One-to-One)
- 一对多关联 (One-to-Many)
- 多对多关联 (Many-to-Many)
- BelongsTo、HasOne、HasMany、ToMany等关联方法详解

#### 5. 数据验证与事务
- 结构体字段标签及其验证功能
- 使用`Valid`函数进行模型校验
- 创建与管理事务

#### 6. 时间戳自动填充与软删除
- 自动创建时间（CreatedAt）和更新时间（UpdatedAt）
- 软删除（DeletedAt）实现逻辑删除

#### 7. 进阶操作
- 执行原生SQL查询
- 使用预编译SQL防止SQL注入
- 更新记录时排除未更改的字段
- 查询构造器的高级用法

#### 8. 高级特性
- 动态表名与表结构变更
- SQL钩子（Before/After Create/Update/Delete）
- 自定义日志与事件监听
- 分组查询与聚合函数

#### 9. 实战案例分析
- 建立简单的博客系统模型设计与数据库操作
- 用户权限系统的设计与实现
- 应用缓存集成（如Redis）

### GORM高级主题

#### 10. 性能优化与最佳实践
- 批量插入与批量更新
- 查询性能调优：避免N+1问题
- 数据库索引设计与查询优化
- 避免并发冲突的策略

#### 11. 微服务架构下的GORM应用
- 分布式事务处理
- 数据库读写分离与负载均衡
- ORM在微服务间的共享与隔离