package cache

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valkey-io/valkey-go"
)

type ICacheService interface {
	Get(ctx *fiber.Ctx, key string, value interface{}) error
	Set(ctx *fiber.Ctx, key string, value interface{}, expiration time.Duration) error
	Delete(ctx *fiber.Ctx, key string) error
}

type CacheService struct {
	rdb valkey.Client
}

func NewCacheService(rdb valkey.Client) *CacheService {
	return &CacheService{rdb: rdb}
}
