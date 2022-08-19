package stores

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	client redis.UniversalClient
}

func (r redisStore) Connect(context context.Context) error {
	return nil
}

func (r redisStore) SetValue(context context.Context, key, value string) error {
	return r.client.Set(context, key, value, 300).Err()

}

func (r redisStore) GetValue(context context.Context, key string) (string, error) {
	return r.client.Get(context, key).Result()
}

func (r redisStore) ValueExists(context context.Context, key string) (bool, error) {
	err := r.client.Exists(context, key).Err()
	if err != nil {
		return false, err
	}
	return true, nil

}

func (r redisStore) DeleteValue(context context.Context, key string) error {
	return r.client.Del(context, key).Err()
}

func (r redisStore) HashSetValue(context context.Context, name, key, value string) error {
	return r.client.HSet(context, key, name, value).Err()
}

func (r redisStore) HashGetValue(context context.Context, name, key string) (string, error) {
	return r.client.HGet(context, key, name).Result()
}

func (r redisStore) HashValueExists(context context.Context, name, key string) (bool, error) {
	return r.client.HExists(context, key, name).Result()
}

func (r redisStore) HashDeleteValue(context context.Context, name, key string) error {
	return r.client.HDel(context, key, name).Err()
}

func (r redisStore) HashExists(context context.Context, name string) (bool, error) {
	return r.ValueExists(context, name)
}

func (r redisStore) HashDelete(context context.Context, name string) error {
	return r.DeleteValue(context, name)
}

func (r redisStore) Close(context context.Context) error {
	return r.client.Close()
}

func NewRedisStore(client redis.UniversalClient) Store {
	return &redisStore{
		client: client,
	}
}
