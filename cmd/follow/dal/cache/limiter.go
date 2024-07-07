package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

func Limit(ctx context.Context, rate int, interval time.Duration) error {
	key := LimiterKey
	now := time.Now().UnixNano()
	RedisClient.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", now-int64(interval)))
	size, err := RedisClient.ZCard(ctx, key).Result()
	if err != nil {
		return err
	}
	if size >= int64(rate) {
		klog.Info("limit!")
		return errors.New("too many request")
	}
	RedisClient.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: float64(now)})
	return nil
}
