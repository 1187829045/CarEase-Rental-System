package Init

import (
	"car.rental/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func MysqlInit() {
	//charset=utf8mb4：指定字符集为 utf8mb4，这是 UTF-8 的一个超集，支持存储 Emoji 字符等 4 字节字符。确保数据在存储和读取时不会出现字符编码问题。
	//parseTime=True：将 MySQL 中的 DATETIME、DATE、TIMESTAMP 字段自动解析为 Go 语言的 time.Time 类型。如果不设置，时间类型可能会被解析为字符串。
	//loc=Local：设置时区为本地时间。MySQL 中的时间通常以 UTC 存储，使用 loc=Local 可以让 Go 语言中的 time.Time 类型自动转换为本地时区。

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)
	var err error

	global.DB, err = gorm.Open(mysql.Open("root:123456@tcp(116.62.145.236:3307)/car.rental?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

}
