package orm

import (
	"big_data_new/sErr"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var mysqlDefDb *MysqlDb
var mysqlDbs map[string]*MysqlDb

type MysqlDb struct {
	Host          string
	Port          string
	Database      string
	Charset       string
	UserName      string
	Password      string
	MaxIdleCounts int
	MaxLifetime   int
	DB            *gorm.DB
}

func init() {
	mysqlDbs = make(map[string]*MysqlDb)
}

func SetMysql(k string, db *MysqlDb, isDef bool) {
	mysqlDbs[k] = db
	if isDef {
		mysqlDefDb = db
	}
}

func GetMysql(k string) (*MysqlDb, bool) {
	ok, d := mysqlDbs[k]
	return ok, d
}

func GetMysqlDef() *MysqlDb {
	return mysqlDefDb
}

func NewDatabaseMysql(host, port, database, charset, username, password string,
	maxIdleCounts int,
	maxOpenCounts int,
	maxLifetime int,
	loggerLevel logger.LogLevel) (*MysqlDb, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  loggerLevel, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(username+":"+password+"@("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local"), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})

	if err != nil {
		return nil, sErr.NewByError(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if maxIdleCounts > 0 {
		sqlDB.SetMaxIdleConns(maxIdleCounts)
	}
	if maxOpenCounts > 0 {
		sqlDB.SetMaxOpenConns(maxOpenCounts)
	}
	if maxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return &MysqlDb{
		Host:          host,
		Port:          port,
		Database:      database,
		Charset:       charset,
		UserName:      username,
		Password:      password,
		MaxIdleCounts: maxIdleCounts,
		MaxLifetime:   maxLifetime,
		DB:            db,
	}, nil
}
