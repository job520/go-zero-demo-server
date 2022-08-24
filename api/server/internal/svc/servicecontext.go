package svc

import (
	"demo/api/server/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.CacheRedis[0].Host,
		Password: c.CacheRedis[0].Pass, // no password set
		DB:       0,                    // use default DB
	})
	db, _ := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	return &ServiceContext{
		Config: c,
		Db:     db,
		Redis:  rdb,
	}
}
