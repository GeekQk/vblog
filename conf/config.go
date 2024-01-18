package conf

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 这里不采用直接暴露变量，而是通过函数来获取配置对象
var config *Config

// 保证只初始化一次
var once sync.Once

// C 函数返回Config类型的指针，如果config为nil，则初始化config为默认值。
func C() *Config {
	if config == nil {
		once.Do(func() {
			config = DefaultConfig()
		})
	}
	return config
}

// 程序配置对象 启动时会加载配置文件 为程序提供全局变量
// 把配置对象做成单例模式的全局变量，避免重复创建对象
type Config struct {
	MySql *MySql `json:"mysql" yaml:"mysql" toml:"mysql"`
}

// 返回默认配置对象 MySql是二级指针 应该有默认值 防止空指针
func DefaultConfig() *Config {
	return &Config{MySql: &MySql{
		Host:     "rm-bp13i71cqtb0f0b5nyo.mysql.rds.aliyuncs.com",
		Port:     3306,
		DB:       "vblog",
		Username: "root",
		Password: "qiKAI!!395166",
		Debug:    true,
	}}
}

// stringer实现
func (c *Config) String() string {
	sr, err := json.MarshalIndent(c, " ", "  ")
	if err != nil {
		return fmt.Sprintf("%s", err)
	}
	return string(sr)
}

type MySql struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`
	// 数据库锁
	l sync.Mutex
	// 判断是否实例化
	db *gorm.DB
}

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *MySql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

// 获取数据库连接实例
func (m *MySql) GetDB() *gorm.DB {
	m.l.Lock()
	defer m.l.Unlock()

	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db

		// 补充Debug配置
		if m.Debug {
			m.db = db.Debug()
		}
	}

	return m.db
}

// 返回真正连接实例
func (c *Config) DB() *gorm.DB {
	return c.MySql.GetDB()
}
