package dbutil

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-websocket/config"
	"go-websocket/tools"
	"go-websocket/tools/fileutil"
	"go-websocket/tools/timer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"sync"
	"time"
)

var (
	once     sync.Once
	Instance *gorm.DB
)

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func InitDbLine() {
	once.Do(func() {
		Instance = connect()
	})
}

// 获取MySQL实例化
func GetMysqlClient() *gorm.DB {
	if Instance == nil {
		tools.EchoErrorExit("mysql is not initialized")
	}
	return Instance
}

func connect() *gorm.DB {
	dbLines := config.GetConf()

	var db *gorm.DB
	var err error

	db, err = gorm.Open(getDialector(dbLines.Mysql.Addr), getGormConfig())

	if err != nil {
		tools.EchoErrorExit(fmt.Sprintf("Failed to connect to mysql database: %v", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		tools.EchoErrorExit(fmt.Sprintf("Failed to initialize mysql database: %v", err))
	}

	//最大空闲连接数（MaxIdleConns）：这是连接池中允许保持的最大空闲连接数。空闲连接是指当前没有被应用程序使用的连接。当应用程序需要新的数据库连接时，它会首先尝试从空闲连接中获取。如果空闲连接数已达到最大限制，则新的连接将被创建。
	sqlDB.SetMaxIdleConns(dbLines.Mysql.SetMaxIdleConn) // SetMaxIdleConn 设置空闲连接池中连接的最大数量
	//最大打开连接数（MaxOpenConns）：这是连接池允许打开的最大连接数，包括活动连接和空闲连接。当应用程序需要新的数据库连接时，如果当前活动和空闲连接数之和已达到最大限制，则新的连接请求将被阻塞，直到有连接可用。
	sqlDB.SetMaxOpenConns(dbLines.Mysql.SetMaxOpenConn) // SetMaxOpenConn 设置打开数据库连接的最大数量。
	//连接的最大生存时间（ConnMaxLifetime）：这是连接在连接池中允许存在的最长时间。当连接的使用时间超过这个阈值时，连接将被关闭并从连接池中移除。这个设置可以防止连接在长时间不使用后变得不稳定或过时。
	//fmt.Println(dbLines.Mysql.SetMaxIdleConn, dbLines.Mysql.SetMaxOpenConn, "dbLines.Mysql.SetConnMaxLifetime * time.Minute:", dbLines.Mysql.SetConnMaxLifetime*time.Minute)
	sqlDB.SetConnMaxLifetime(dbLines.Mysql.SetConnMaxLifetime * time.Minute) // SetConnMaxLifetime 设置了连接可复用的最大时间。
	return db
}

func getDialector(dsn string) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  false, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    false, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   false, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})
}

func getGormConfig() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   dbLines.Mysql.TablePrefix,         // 表名前缀，将在所有表名前添加该前缀。表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user` 是否使用单数表名，如果设置为 true，则表名将使用单数形式，例如 User 对应表名 user。
			//NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。 自定义的名称替换器，可以通过该字段设置替换规则，将结构体或字段名中的某些字符串替换为指定的字符串。
			NoLowerCase: false, //是否禁用表名和字段名的小写转换，如果设置为 true，则表名和字段名将保持原样，不进行小写转换。举个例子，如果数据库中的表名是 "UserInfo"，默认情况下 GORM 会将其转换为小写形式 "userinfo"。但是如果将 NoLowerCase 设置为 true，那么 GORM 将保持表名的原始大小写，即 "UserInfo"。同样的规则也适用于字段名。
		},
		Logger:                                   getLogger(),
		DryRun:                                   false, //生成 SQL 但不执行，可以用于准备或测试生成的 SQL
		PrepareStmt:                              true,  //PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		DisableAutomaticPing:                     true,  //在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: true,  //在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true，参考 迁移 获取详情
		AllowGlobalUpdate:                        true,  //GORM 默认不允许进行全局 update/delete，该操作会返回 ErrMissingWhereClause 错误。 您可以通过将一个选项设置为 true 来启用它，例如
	}
}

func getLogger() logger.Interface {
	//Silent：静默模式，不输出任何日志。 Error：错误级别，只输出错误日志。 Warn：警告级别，输出错误和警告日志。 Info：信息级别，输出错误、警告和信息日志。
	return logger.New(
		initLog(),
		logger.Config{
			SlowThreshold:             time.Duration(config.GetConf().Mysql.SlowThreshold) * time.Millisecond, // 设置慢查询的阈值，超过该阈值的查询将被认为是慢查询，默认单位为纳秒
			LogLevel:                  getLogLevel(),                                                          // GORM 定义了这些日志级别：Silent、Error、Warn、Info
			IgnoreRecordNotFoundError: true,                                                                   // 设置是否忽略 gorm.ErrRecordNotFound 错误，如果设置为 true，则不会将该错误记录到日志中。
			ParameterizedQueries:      false,                                                                  // 设置是否在日志中包含 SQL 查询的参数信息 如果将其设置为 true，则日志中将不会包含参数信息，只会显示占位符。
			Colorful:                  false,                                                                  // 设置是否在终端中输出带有颜色的日志。
		},
	)
}

// 初始化日志
func initLog() *log.Logger {
	mysqlConfig := config.GetConf()
	if mysqlConfig.Mysql.Cmd {

		return log.New(os.Stdout, "\r\n", log.LstdFlags) // 打印到控制台

	} else {
		fileLogName := fmt.Sprintf("%s/mysql_%s.log", fileutil.GetAbsDirPath(mysqlConfig.Mysql.LogFolder), timer.GetDateId())
		// 创建日志文件
		file, err := os.OpenFile(fileLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			// 错误处理
			tools.EchoError(fmt.Sprintf("Open MySQL file err: %s", err))
		}

		return log.New(file, "\r\n", log.LstdFlags) // 保存到文件
	}
}

// 日志级别
func getLogLevel() logger.LogLevel {
	//日志级别 0 Silent：静默模式，不输出任何日志。 1 Error：错误级别，只输出错误日志。 2 Warn：警告级别，输出错误和警告日志。 3 Info：信息级别，输出错误、警告和信息日志。
	switch config.GetConf().Mysql.LogLevel {
	case 0:
		return logger.Silent
	case 1:
		return logger.Error
	case 2:
		return logger.Warn
	case 3:
		return logger.Info
	}
	return logger.Silent
}

//func t() *gorm.Config {
//	return &gorm.Config{
//		SkipDefaultTransaction :false,//默认情况下，GORM 在执行单个的创建、更新、删除操作时会自动开启事务来确保数据库数据的完整性。你可以将该字段设置为 true 来禁用默认的事务功能。
//		NamingStrategy : schema.NamingStrategy{},//表和字段命名策略，用于指定自定义的命名策略，例如表名前缀、驼峰命名转换等。
//		//logger logger.Interface//: 自定义的日志记录器，用于记录 GORM 的日志信息。
//		FullSaveAssociations :true,//是否完整保存关联数据，默认为 false。设置为 true 时，GORM 在保存关联数据时会完整保存所有相关的关联数据。
//		NowFunc func() time.Time // 创建新时间戳时使用的函数。可以自定义一个函数，返回当前时间，用于生成时间戳字段的默认值。
//		DryRun bool //是否生成 SQL 语句但不执行，默认为 false。设置为 true 时，GORM 会生成 SQL 语句但不实际执行，用于调试和测试目的。
//		PrepareStmt bool //是否使用缓存的语句执行查询，默认为 false。设置为 true 时，GORM 会将查询语句缓存起来，提高后续查询的性能。
//		DisableAutomaticPing bool //是否禁用自动连接检查，默认为 false。设置为 true 时，GORM 不会自动进行数据库连接检查。
//		DisableForeignKeyConstraintWhenMigrating bool //在迁移时是否禁用外键约束，默认为 false。设置为 true 时，GORM 在执行数据库迁移时会禁用外键约束。
//		IgnoreRelationshipsWhenMigrating bool // 在迁移时是否忽略关联关系，默认为 false。设置为 true 时，GORM 在执行数据库迁移时会忽略模型之间的关联关系
//		DisableNestedTransaction bool //是否禁用嵌套事务，默认为 false。设置为 true 时，GORM 不支持嵌套事务。
//		AllowGlobalUpdate bool //是否允许全局更新，默认为 false。设置为 true 时，GORM 允许在更新操作时更新所有字段，而不仅仅是变化的字段。
//		QueryFields bool //: 是否执行包含所有表字段的 SQL 查询，默认为 false。设置为 true 时，GORM 执行的 SQL 查询会包含所有表字段。
//		CreateBatchSize int // 默认的批量创建大小，默认为 0，表示未设置批量创建大小。设置为大于 0 的数值时，GORM 在进行批量创建时会按照指定大小进行分批创建。
//		TranslateError bool //是否启用错误翻译，默认为 false。设置为 true 时，GORM 会将底层数据库错误翻译为 GORM 定义的错误类型。
//		ClauseBuilders map[string]clause.ClauseBuilder //自定义的查询构建器，用于扩展 GORM 的查询功能。
//		ConnPool ConnPool //数据库连接池，用于管理数据库连接。
//		Dialector //数据库方言，定义了与具体数据库的交互方式。
//		Plugins map[string]Plugin // 注册的插件，用于扩展 GORM 的功能。
//		callbacks  *callbacks // GORM 的回调函数，用于处理各种生命周期事件。
//		cacheStore *sync.Map //缓存存储，用于缓存一些 GORM 的内部数据。
//	}
//}
