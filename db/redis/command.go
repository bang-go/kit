package redis

import "time"

func (s *Client) Get(key string) (value string, err error) {
	value, err = s.Rdb.Get(s.Ctx, key).Result()
	return
}

func (s *Client) Exists(key string) (value int64, err error) {
	value, err = s.Rdb.Exists(s.Ctx, key).Result()
	return
}

func (s *Client) Set(key string, value any, expiration time.Duration) (err error) {
	err = s.Rdb.Set(s.Ctx, key, value, expiration).Err()
	return
}

// SetNX 指定的key不存在时， 为key设置指定的值
func (s *Client) SetNX(key string, value any, expiration time.Duration) (err error) {
	err = s.Rdb.SetNX(s.Ctx, key, value, expiration).Err()
	return
}

// SetEx 为指定的 key 设置值及其过期时间。如果key已经存在， SETEX命令将会替换旧的值。
func (s *Client) SetEx(key string, value any, expiration time.Duration) (err error) {
	err = s.Rdb.SetEx(s.Ctx, key, value, expiration).Err()
	return
}

func (s *Client) Incr(key string) (value int64, err error) {
	value, err = s.Rdb.Incr(s.Ctx, key).Result()
	return
}

func (s *Client) Decr(key string) (value int64, err error) {
	value, err = s.Rdb.Decr(s.Ctx, key).Result()
	return
}

// Do 自定义命令
func (s *Client) Do(args ...any) (value any, err error) {
	value, err = s.Rdb.Do(s.Ctx, args).Result()
	return
}
