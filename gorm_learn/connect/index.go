package connect

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type StudentInfo struct {
  Name      string
  Age       int
  MyStudent string
}

func ConnectFunc() (db *gorm.DB) {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "learn_gorm"
	timeout := "10s"

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	// 设置日志
	var mysqlLogger logger.Interface
	// 要显示的日志等级 gorm的默认日志是只打印错误和慢SQL
	mysqlLogger = logger.Default.LogMode(logger.Info)

	// 自定义日志显示
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 使用彩色打印
		},
	)

	// 连接MYSQL，获得DB类型实例，用于后面的数据库读写操作
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 跳过默认事务
		// 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这样可以获得60%的性能提升
		SkipDefaultTransaction: true,
		// 默认日志设置
		Logger: mysqlLogger,
		// Logger: newLogger, // 自定义日志设置
		// 更改命名约定
		NamingStrategy: schema.NamingStrategy{ 
			TablePrefix:   "t_",  // 表名前缀
			SingularTable: true, // 单数表名 eg:user  false: users  true: user 
			NoLowerCase:   false, // 驼峰 eg:MyStudent false:  my_student  true: mystudent
		},
	})
	if err != nil{
		panic("连接数据库失败，error: "+err.Error())
	}

	// 日志打印
	var model StudentInfo
	session := db.Session(&gorm.Session{Logger: newLogger})
	session.First(&model)  // SELECT * FROM `t_students` ORDER BY `t_students`.`name` LIMIT 1
	
	db.Debug().First(&model)

	return
}
