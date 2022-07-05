package redisMe

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"skeleton/app/global/variable"
	"time"
)

var ctx = context.Background()

// Client redis缓存客户端
type Client struct {
	rc *redis.Client
}

// GetKey 获取一个Key
func (c Client) GetKey(k string) (string, error) {
	err := c.Connect()
	if err != nil {
		return "", err
	}
	defer c.Close()
	return c.rc.Get(ctx, k).Result()
}

// SetKey 设置一个Key
func (c Client) SetKey(k, v string, ttl time.Duration) error {
	err := c.Connect()
	if err != nil {
		return err
	}
	defer c.Close()
	return c.rc.Set(ctx, k, v, ttl).Err()
}

// Connect Redis连接
func (c *Client) Connect() error {

	rClient := redis.NewClient(&redis.Options{
		Addr:       variable.ConfigYml.GetString("Redis.Host") + ":" + variable.ConfigYml.GetString("Redis.Port"),
		Password:   variable.ConfigYml.GetString("Redis.Password"),
		DB:         variable.ConfigYml.GetInt("Redis.Db"),
		MaxRetries: variable.ConfigYml.GetInt("Redis.MaxRetries"),
		//可扩展pool设置此处仅使用默认
	})

	timeoutCtx, cancelFunc := context.WithTimeout(ctx, time.Second*3)
	defer cancelFunc()

	_, err := rClient.Ping(timeoutCtx).Result()
	if err != nil {
		return errors.New("Redis连接失败！" + err.Error())
	} else {
		c.rc = rClient
	}
	return nil
}

// Close 关闭连接
func (c Client) Close() {
	_ = c.rc.Close()
}
