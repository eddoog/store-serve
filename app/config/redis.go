package config

import "os"

type Redis struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func LoadRedis() Redis {
	return Redis{
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       0,
	}
}
