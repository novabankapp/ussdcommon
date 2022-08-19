package stores

type Store interface {
	Connect() error
	SetValue(key, value string) error
	GetValue(string) (string, error)
	ValueExists(string) (bool, error)
	DeleteValue(string) error
	HashSetValue(name, key, value string) error
	HashGetValue(name, key string) (string, error)
	HashValueExists(name, key string) (bool, error)
	HashDeleteValue(name, key string) error
	HashExists(string) (bool, error)
	HashDelete(string) error
	Close() error
}
