package stores

import "context"

type MockStore struct {
}

func (m MockStore) Connect(context context.Context) error {
	return nil
}

func (m MockStore) SetValue(context context.Context, key, value string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) GetValue(context context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) ValueExists(context context.Context, key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) DeleteValue(context context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) HashSetValue(context context.Context, name, key, value string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) HashGetValue(context context.Context, name, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) HashValueExists(context context.Context, name, key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) HashDeleteValue(context context.Context, name, key string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) HashExists(context context.Context, name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) HashDelete(context context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockStore) Close(context context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewMockStore() Store {
	return &MockStore{}
}
