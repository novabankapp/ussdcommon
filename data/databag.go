package data

import (
	"context"
	"encoding/json"
	"github.com/novabankapp/ussdcommon/domain/models/adapters"
	"github.com/novabankapp/ussdcommon/stores"
)

// DataBag is a key-value store.
// Used to store data that will be available across USSD session's
// request.
type DataBag struct {
	name  string
	store stores.Store
}

func NewDataBag(store stores.Store, request *adapters.Request) *DataBag {
	name := request.PhoneNumber + "DataBag"
	return &DataBag{
		name:  name,
		store: store,
	}
}

// Set value in databag
func (d DataBag) Set(context context.Context, key, value string) error {
	return d.store.HashSetValue(context, d.name, key, value)
}

// Get value from databag
func (d DataBag) Get(context context.Context, key string) (string, error) {
	return d.store.HashGetValue(context, d.name, key)
}

// Exists verifies if value is in databag
func (d DataBag) Exists(context context.Context, key string) (bool, error) {
	return d.store.HashValueExists(context, d.name, key)
}

// Delete value from databag
func (d DataBag) Delete(context context.Context, key string) error {
	return d.store.HashDeleteValue(context, d.name, key)
}

// Clear databag
func (d DataBag) Clear(context context.Context) error {
	return d.store.HashDelete(context, d.name)
}

func (d DataBag) SetMarshaled(context context.Context, key string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return d.Set(context, key, string(b))
}

func (d DataBag) GetUnmarshaled(context context.Context, key string, v interface{}) error {
	str, err := d.Get(context, key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), v)
}
