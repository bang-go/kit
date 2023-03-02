package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Config redis.Options
type Options struct {
	Ctx    context.Context
	Config Config
}

type Client struct {
	Ctx context.Context
	Rdb *redis.Client
}

var (
	Nil = redis.Nil
)

// New 新建实例
func New(opt *Options) *Client {
	cfg := redis.Options(opt.Config)
	rdb := redis.NewClient(&cfg)
	ctx := opt.Ctx
	if opt.Ctx == nil {
		ctx = context.Background()
	}
	return &Client{
		Ctx: ctx,
		Rdb: rdb,
	}
}
