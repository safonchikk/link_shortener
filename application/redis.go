package application

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"link_shortener/model"
)

type RedisRepo struct {
	Client *redis.Client
}

func (r *RedisRepo) Insert(ctx context.Context, link model.Link) error {
	data, err := json.Marshal(link)
	if err != nil {
		return fmt.Errorf("failed to make short link: %w", err)
	}
	key := link.Short
	res := r.Client.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to set: %w", err)
	}
	return nil
}

func (r *RedisRepo) FindByShort(ctx context.Context, short string) (model.Link, error) {
	res, err := r.Client.Get(ctx, short).Result()
	if err != nil {
		return model.Link{}, fmt.Errorf("failed to find: %w", err)
	}
	var data model.Link
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return model.Link{}, fmt.Errorf("error: %w", err)
	}
	return data, nil
}