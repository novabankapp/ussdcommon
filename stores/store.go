package stores

import "context"

type Store interface {
	Connect(context context.Context) error
	SetValue(context context.Context, key, value string) error
	GetValue(context context.Context, key string) (string, error)
	ValueExists(context context.Context, key string) (bool, error)
	DeleteValue(context context.Context, key string) error
	HashSetValue(context context.Context, name, key, value string) error
	HashGetValue(context context.Context, name, key string) (string, error)
	HashValueExists(context context.Context, name, key string) (bool, error)
	HashDeleteValue(context context.Context, name, key string) error
	HashExists(context context.Context, name string) (bool, error)
	HashDelete(context context.Context, name string) error
	Close(context context.Context) error
}
