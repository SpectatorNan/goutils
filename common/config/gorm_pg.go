package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type PgSql struct {
	Username     string
	Password     string
	Path         string
	Port         string `json:",default=5432"`
	SslMode      string `json:",default=pro,options=disable|enable"`
	TimeZone     string `json:",default=Asia/Shanghai"`
	Dbname       string
	MaxIdleConns int    `json:",default=10"` // 空闲中的最大连接数
	MaxOpenConns int    `json:",default=10"` // 打开到数据库的最大连接数
	LogMode      string `json:",default="`   // 是否开启Gorm全局日志
	LogZap       bool   // 是否通过zap写入日志文件
}

func (m *PgSql) Dsn() string {

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%s", m.Username, m.Password, m.Dbname, m.Path, m.Port, m.SslMode, m.TimeZone)
}

//func (m *PgSql) Dsn() string {
//	return "postgres://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "/" + m.Dbname + "?" + m.Config
//}

func (m *PgSql) GetLogMode() string {
	return m.LogMode
}

func ConnectPgSql(m PgSql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	newLogger := logger.New(
		log.New(os.Stderr, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	pgsqlCfg := postgres.Config{
		DSN:                  m.Dsn(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	db, err := gorm.Open(postgres.New(pgsqlCfg), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil
	} else {
		sqldb, _ := db.DB()
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
