package config

import "github.com/spf13/viper"

const (
	RedisModeCluster = "cluster"
	RedisModeSingle  = "single"
)

type RedisCfg struct {
	Mode    string
	Single  RedisSingleCfg
	Cluster RedisClusterCfg
}

type RedisSingleCfg struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

type RedisClusterCfg struct {
	Addrs    []string
	Password string
	PoolSize int
}

func loadRedisCfg() error {
	r := RedisCfg{}
	r.Mode = viper.GetString("redis.mode")

	if r.Mode == "cluster" {
		r.Cluster.Addrs = viper.GetStringSlice("redis.cluster.addrs")
		r.Cluster.Password = viper.GetString("redis.cluster.password")
		r.Cluster.PoolSize = viper.GetInt("redis.cluster.pool_size")
	} else {
		// 默认 single
		r.Mode = "single"
		r.Single.Addr = viper.GetString("redis.single.addr")
		r.Single.Password = viper.GetString("redis.single.password")
		r.Single.DB = viper.GetInt("redis.single.db")
		r.Single.PoolSize = viper.GetInt("redis.single.pool_size")
	}

	cfg.redis = r
	return nil
}
