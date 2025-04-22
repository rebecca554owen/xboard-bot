package mysql

import (
	"fmt"
	"log"
	"sync"
	"time"

	"xboard-bot/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	client Repo
	once   sync.Once
)

// Predicate 是where子句中的条件
type Predicate string

var (
	EqualPredicate              = Predicate("=")
	NotEqualPredicate           = Predicate("<>")
	GreaterThanPredicate        = Predicate(">")
	GreaterThanOrEqualPredicate = Predicate(">=")
	SmallerThanPredicate        = Predicate("<")
	SmallerThanOrEqualPredicate = Predicate("<=")
	LikePredicate               = Predicate("LIKE")
)

// Repo 定义了数据库仓库接口
// 包含获取数据库连接和关闭连接的方法
type Repo interface {
	// GetDb 获取GORM数据库连接实例
	GetDb() *gorm.DB
	// DbClose 关闭数据库连接
	DbClose() error
} 

// dbRepo 实现了Repo接口的数据库仓库结构体
type dbRepo struct {
	Db *gorm.DB // GORM数据库连接实例
}

// GetDbClient 获取数据库连接实例
func GetDbClient(cfg *config.Config) Repo {
	once.Do(func() {
		c, err := NewDb(cfg)
		if err != nil {
			panic("无法获取数据库连接")
		}
		client = c
	})
	return client
}

// NewDb 创建并返回一个新的数据库仓库实例
// 从配置中读取MySQL连接参数并建立连接
// 返回Repo接口实现和可能的错误
func NewDb(cfg *config.Config) (Repo, error) {
	db, err := dbConnect(cfg)
	if err != nil {
		return nil, err
	}

	log.Printf("数据库连接成功: %s@%s:%d/%s", 
		cfg.MySQL.User,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Database)

	return &dbRepo{
		Db: db,
	}, nil
}

// GetDb 获取GORM数据库连接实例
func (d *dbRepo) GetDb() *gorm.DB {
	return d.Db.Debug() // 启用了Debug模式，会打印SQL日志
	// return d.Db // 不启用Debug模式
}

// DbClose 关闭数据库连接
// 首先获取底层sql.DB实例，然后关闭连接
func (d *dbRepo) DbClose() error {
	sqlDB, err := d.Db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// dbConnect 建立到MySQL数据库的连接 
func dbConnect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("[数据库连接失败] 数据库名称: %s", cfg.MySQL.Database)
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 2)

	return db, nil
}
