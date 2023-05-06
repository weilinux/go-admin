package myrds

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/weilinux/go-gin-skeleton-auth/app"
	"go.uber.org/zap"
)

type redisConfig struct {
	Server  string
	Auth    string
	Db      int
	Disable bool
}

var (
	cfg  redisConfig
	pool *redis.Pool
)

// redis Prefix
const redisPrefix = "rds:"

// GenRedisKey func
func GenRedisKey(tpl string, keys ...interface{}) string {
	if len(keys) == 0 {
		return redisPrefix + tpl
	}

	return redisPrefix + fmt.Sprintf(tpl, keys...)
}

// InitRedis
// init redis connection pool
// redigo doc https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
func InitRedis() (err error) {
	// 从配置文件获取redis的ip以及db
	err = app.Config.MapStruct("redis", &cfg)
	if err != nil {
		return
	}
	if cfg.Disable {
		return
	}

	fmt.Printf("redis - server=%s db=%d auth=%s\n", cfg.Server, cfg.Db, cfg.Auth)

	// 建立连接池
	pool = app.NewRedisPool(cfg.Server, cfg.Auth, cfg.Db)

	// closePool() 在main函数中检测关闭程序则关闭连接池
	return
}

func ClosePool() error {
	if cfg.Disable {
		return nil
	}
	return pool.Close()
}

func Connection() redis.Conn {
	app.Logger.Info("get new redis connection from pool",
		zap.Namespace("context"),
		zap.Int("IdleCount", pool.IdleCount()),
		zap.Int("ActiveCount", pool.ActiveCount()),
	)

	// 记录操作的日志
	if app.Debug {
		return redis.NewLoggingConn(pool.Get(), zap.NewStdLog(app.Logger), "rds")
	}

	return pool.Get()
}

func WithConnection(fn func(c redis.Conn) (interface{}, error)) (interface{}, error) {
	conn := Connection()
	defer conn.Close()

	return fn(conn)
}

func HasZSet(key string) bool {
	count, _ := redis.Int(WithConnection(func(c redis.Conn) (interface{}, error) {
		return c.Do("zCard", key)
	}))
	return count > 0
}
