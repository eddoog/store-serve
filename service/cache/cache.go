package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valkey-io/valkey-go"
)

func (cs *CacheService) Get(ctx *fiber.Ctx, key string, value interface{}) error {
	res, err := cs.rdb.Do(ctx.Context(), cs.rdb.B().Get().Key(key).Build()).ToString()

	if err != nil {
		if err == valkey.Nil {
			return fmt.Errorf("cache: key %s not found", key)
		}
		return err
	}

	return json.Unmarshal([]byte(res), value)
}

func (cs *CacheService) Set(ctx *fiber.Ctx, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return cs.rdb.Do(ctx.Context(), cs.rdb.B().Set().Key(key).Value(string(data)).Ex(expiration).Build()).Error()
}

func (cs *CacheService) Delete(ctx *fiber.Ctx, key string) error {
	return cs.rdb.Do(ctx.Context(), cs.rdb.B().Del().Key(key).Build()).Error()
}
