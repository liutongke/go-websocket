package DbLine

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-websocket/config"
	"go-websocket/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var (
	once     sync.Once
	Instance *gorm.DB
)

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func init() {
	config.Init() //初始化配置文件,test用例时候使用的平时可以删除
	once.Do(func() {
		Instance = connect()
	})
}

//获取MySQL实例化
func GetMysqlClient() *gorm.DB {
	return Instance
}

func connect() *gorm.DB {
	dbLines := config.GetConfClient()
	dsn := dbLines.Mysql.Addr + "?charset=utf8&parseTime=True&loc=Local"
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second,   // 慢 SQL 阈值
	//		LogLevel:                  logger.Silent, // Log level
	//		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	//		Colorful:                  false,         // 禁用彩色打印
	//	},
	//)
	var db *gorm.DB
	var err error
	if utils.IsDebug() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   dbLines.Mysql.TablePrefix, // 表名前缀，`User`表为`t_users`
				SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
				//NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   dbLines.Mysql.TablePrefix, // 表名前缀，`User`表为`t_users`
				SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
				//NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
			},
		})
	}
	if err != nil {
		panic("gorm连接错误：" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("gorm数据库连接池错误：" + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。
	return db
}
