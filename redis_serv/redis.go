package redis_serv

import (
	"github.com/redis/go-redis/v9"
	"iris-init/config"
	"iris-init/global"
)

var (
	rdb redis.UniversalClient
)

func GetRdb() redis.UniversalClient {
	return rdb
}

func InitRedis() error {
	cfg := config.GetRedisCfg()
	if cfg.Mode == config.RedisModeCluster {
		rdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    cfg.Cluster.Addrs,
			Password: cfg.Cluster.Password,
			PoolSize: cfg.Cluster.PoolSize,
		})
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.Single.Addr,
			Password: cfg.Single.Password,
			DB:       cfg.Single.DB,
			PoolSize: cfg.Single.PoolSize,
		})
	}

	return rdb.Ping(global.GetGlobalCtx()).Err()
}
